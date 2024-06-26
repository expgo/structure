package structure

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"sync"
)

type Mapper func(from reflect.Value, to reflect.Value, option *Option) error

type typeKey struct {
	from reflect.Type
	to   reflect.Type
}

var typeMappers = make(map[typeKey]Mapper)
var typeMappersLock = &sync.RWMutex{}

type kindKey struct {
	from reflect.Kind
	to   reflect.Kind
}

var kindMappers = make(map[kindKey]Mapper)
var kindMappersLock = &sync.RWMutex{}

func RegisterMapper[From any, To any](mapper Mapper) {
	registerMapper(reflect.TypeOf((*From)(nil)).Elem(), reflect.TypeOf((*To)(nil)).Elem(), mapper)
}

func registerMapper(fromType reflect.Type, toType reflect.Type, mapper Mapper) {
	key := typeKey{from: fromType, to: toType}

	typeMappersLock.Lock()
	defer typeMappersLock.Unlock()

	if _, got := typeMappers[key]; got {
		panic(fmt.Sprintf("type %+v already registed", key))
	} else {
		typeMappers[key] = mapper
	}
}

func RegisterKindMapper(from reflect.Kind, to reflect.Kind, mapper Mapper) {
	key := kindKey{from: from, to: to}
	kindMappersLock.Lock()
	defer kindMappersLock.Unlock()

	if _, got := kindMappers[key]; got {
		panic(fmt.Sprintf("type %+v already registed", key))
	} else {
		kindMappers[key] = mapper
	}
}

func ReplaceMapper[From any, To any](mapper Mapper) {
	fromType := reflect.TypeOf((*From)(nil)).Elem()
	toType := reflect.TypeOf((*To)(nil)).Elem()

	typeMappersLock.Lock()
	defer typeMappersLock.Unlock()

	typeMappers[typeKey{from: fromType, to: toType}] = mapper
}

func NewOption() *Option {
	return &Option{
		ZeroFields:           true,
		WeaklyTypedInput:     true,
		StringSplitSeparator: ",",
		TagName:              "mapping",
		MatchName:            strings.EqualFold,
	}
}

var defaultOption = NewOption()

func Map(from, to any) error {
	return MapWithOption(from, to, defaultOption)
}

func MapWithOption(from, to any, option *Option) error {
	return MapToValueWithOption(from, reflect.Indirect(reflect.ValueOf(to)), option)
}

func MapToValue(from any, to reflect.Value) error {
	return MapToValueWithOption(from, to, defaultOption)
}

func MapToValueWithOption(from any, to reflect.Value, option *Option) error {
	if option == nil {
		option = defaultOption
	}

	if to.Kind() != reflect.Map {
		if !to.CanSet() {
			return errors.New("to value can't be set")
		}
	}

	var fromVal reflect.Value
	if from != nil {
		fromVal = reflect.ValueOf(from)

		// We need to check here if input is a typed nil. Typed nils won't
		// match the "input == nil" below so we check that here.
		if fromVal.Kind() == reflect.Ptr && fromVal.IsNil() {
			from = nil
		}
	}

	if from == nil {
		// If the data is nil, then we don't set anything, unless ZeroFields is set
		// to true.
		if option.ZeroFields {
			to.Set(reflect.Zero(to.Type()))
		}
		return nil
	}

	return Value2ValueWithOption(fromVal, to, option)
}

func value2valuePtrWithOption(from reflect.Value, to reflect.Value, option *Option) error {
	toElemType := to.Type()
	if toElemType.Kind() == reflect.Ptr {
		toElemType = toElemType.Elem()
	}

	if to.CanSet() {
		realTo := to
		if to.Type().Kind() == reflect.Ptr && realTo.IsNil() {
			realTo = reflect.New(toElemType)
		}

		if err := Value2ValueWithOption(reflect.Indirect(from), reflect.Indirect(realTo), option); err != nil {
			return err
		}

		to.Set(realTo)
	} else {
		if err := Value2ValueWithOption(reflect.Indirect(from), reflect.Indirect(to), option); err != nil {
			return err
		}
	}
	return nil
}

func value2valueSliceWithOption(from reflect.Value, to reflect.Value, option *Option) error {
	dataVal := from
	fromKind := dataVal.Kind()
	valType := to.Type()
	valElemType := valType.Elem()
	sliceType := reflect.SliceOf(valElemType)

	// If we have a non array/slice type then we first attempt to convert.
	if fromKind != reflect.Array && fromKind != reflect.Slice {
		if option.WeaklyTypedInput {
			switch {
			// Slice and array we use the normal logic
			case fromKind == reflect.Slice, fromKind == reflect.Array:
				break

			// Empty maps turn into empty slices
			case fromKind == reflect.Map:
			//if dataVal.Len() == 0 {
			//	to.Set(reflect.MakeSlice(sliceType, 0, 0))
			//	return nil
			//}
			//// Create slice of maps of other sizes
			//return d.decodeSlice(name, []interface{}{data}, val)

			//case fromKind == reflect.String && valElemType.Kind() == reflect.Uint8:
			//return d.decodeSlice(name, []byte(dataVal.String()), val)
			case fromKind == reflect.String:
				return Value2ValueWithOption(reflect.ValueOf(strings.Split(from.String(), option.StringSplitSeparator)), to, option)
			// All other types we try to convert to the slice type
			// and "lift" it into it. i.e. a string becomes a string slice.
			default:
				// Just re-try this function with data as a slice.
				//return d.decodeSlice(name, []interface{}{data}, val)
			}
		}

		return fmt.Errorf("source data must be an array or slice, got %s", fromKind)
	}

	// If the input value is nil, then don't allocate since empty != nil
	if fromKind != reflect.Array && dataVal.IsNil() {
		return nil
	}

	valSlice := to
	if valSlice.IsNil() || option.ZeroFields {
		// Make a new slice to hold our result, same size as the original data.
		valSlice = reflect.MakeSlice(sliceType, dataVal.Len(), dataVal.Len())
	} else if valSlice.Len() > dataVal.Len() {
		valSlice = valSlice.Slice(0, dataVal.Len())
	}

	// Accumulate any errors
	errors := make([]error, 0)

	for i := 0; i < dataVal.Len(); i++ {
		currentData := dataVal.Index(i)
		for valSlice.Len() <= i {
			valSlice = reflect.Append(valSlice, reflect.Zero(valElemType))
		}
		currentField := valSlice.Index(i)

		if err := Value2ValueWithOption(currentData, currentField, option); err != nil {
			errors = appendErrors(errors, err)
		}
	}

	// Finally, set the value to the slice we built up
	to.Set(valSlice)

	// If there were errors, we return those
	if len(errors) > 0 {
		return &Error{errors}
	}

	return nil
}

func Value2ValueWithOption(from reflect.Value, to reflect.Value, option *Option) error {
	if !from.IsValid() {
		// If the input value is invalid, then we just set the value
		// to be the zero value.
		to.Set(reflect.Zero(to.Type()))
		return nil
	}

	if from.CanInterface() {
		if wrapper, ok := from.Interface().(ValueWrapper); ok {
			return MapToValueWithOption(wrapper.Value(), to, option)
		}
	}

	if to.Kind() == reflect.Ptr || from.Kind() == reflect.Ptr {
		return value2valuePtrWithOption(from, to, option)
	}

	if to.Kind() == reflect.Slice {
		return value2valueSliceWithOption(from, to, option)
	}

	fromType := from.Type()
	toType := to.Type()

	kindMappersLock.RLock()
	mapper, ok := kindMappers[kindKey{from: fromType.Kind(), to: toType.Kind()}]
	kindMappersLock.RUnlock()

	if ok {
		if err := mapper(from, to, option); err != nil {
			return err
		}
		return nil
	}

	// if to kind is a interface, and from type can convert to , convert and return
	if toType.Kind() == reflect.Interface && fromType.ConvertibleTo(toType) {
		to.Set(from.Convert(toType))
		return nil
	}

	// use type and to, get mapper
	typeMappersLock.RLock()
	mapper, ok = typeMappers[typeKey{from: fromType, to: toType}]
	typeMappersLock.RUnlock()

	if !ok {
		// do type alias
		typeAliasMapLock.RLock()
		if fromTypeInMap, ok := typeAliasMap[fromType]; ok {
			fromType = fromTypeInMap
		}
		if toTypeInMap, ok := typeAliasMap[toType]; ok {
			toType = toTypeInMap
		}
		typeAliasMapLock.RUnlock()

		typeMappersLock.RLock()
		mapper, ok = typeMappers[typeKey{from: fromType, to: toType}]
		typeMappersLock.RUnlock()
	}

	if ok {
		if err := mapper(from, to, option); err != nil {
			return err
		}
		return nil
	}

	// get all implements interface, no err, return direct
	for k, v := range typeMappers {
		if k.to.Kind() == reflect.Interface &&
			(toType.Implements(k.to) || reflect.PtrTo(toType).Implements(k.to)) {
			if err := v(from, to, option); err == nil {
				return nil
			}
		}
	}

	// if the from and to type is same, set and return direct
	if fromType == toType {
		to.Set(from)
		return nil
	}

	return errors.New(fmt.Sprintf("no mapper found for type %+v to %+v", fromType, toType))
}
