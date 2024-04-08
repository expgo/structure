package structure

import (
	"fmt"
	"reflect"
	"strings"
)

func init() {
	RegisterKindMapper(reflect.Struct, reflect.Map, struct2mapMapper)
	RegisterKindMapper(reflect.Struct, reflect.Struct, struct2structMapper)
}

func getKind(val reflect.Value) reflect.Kind {
	kind := val.Kind()

	switch {
	case kind >= reflect.Int && kind <= reflect.Int64:
		return reflect.Int
	case kind >= reflect.Uint && kind <= reflect.Uint64:
		return reflect.Uint
	case kind >= reflect.Float32 && kind <= reflect.Float64:
		return reflect.Float32
	default:
		return kind
	}
}

func isEmptyValue(v reflect.Value) bool {
	switch getKind(v) {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return false
}

func isStructTypeConvertibleToMap(typ reflect.Type, checkMapstructureTags bool, tagName string) bool {
	for i := 0; i < typ.NumField(); i++ {
		f := typ.Field(i)
		if f.PkgPath == "" && !checkMapstructureTags { // check for unexported fields
			return true
		}
		if checkMapstructureTags && f.Tag.Get(tagName) != "" { // check for mapstructure tags inside
			return true
		}
	}
	return false
}

func dereferencePtrToStructIfNeeded(v reflect.Value, tagName string) reflect.Value {
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return v
	}
	deref := v.Elem()
	derefT := deref.Type()
	if isStructTypeConvertibleToMap(derefT, true, tagName) {
		return deref
	}
	return v
}

func struct2mapMapper(from reflect.Value, to reflect.Value, option *Option) error {
	mapType := to.Type()
	mapElemType := mapType.Elem()

	if to.IsNil() {
		to.Set(reflect.MakeMap(mapType))
	}

	typ := from.Type()
	for i := 0; i < typ.NumField(); i++ {
		// Get the StructField first since this is a cheap operation. If the
		// field is unexported, then ignore it.
		f := typ.Field(i)
		if f.PkgPath != "" {
			continue
		}

		// Next get the actual value of this field and verify it is assignable
		// to the map value.
		v := from.Field(i)
		if !v.Type().AssignableTo(mapElemType) {
			cv, err := ConvertToTypeWithOption(v.Interface(), mapElemType, option)
			if err != nil {
				return err
			}
			v = reflect.ValueOf(cv)
		}

		tagValue := f.Tag.Get(option.TagName)
		keyName := f.Name

		if tagValue == "" && option.IgnoreUntaggedFields {
			continue
		}

		// If Squash is set in the config, we squash the field down.
		squash := option.Squash && v.Kind() == reflect.Struct && f.Anonymous

		v = dereferencePtrToStructIfNeeded(v, option.TagName)

		// Determine the name of the key in the map
		if index := strings.Index(tagValue, ","); index != -1 {
			if tagValue[:index] == "-" {
				continue
			}
			// If "omitempty" is specified in the tag, it ignores empty values.
			if strings.Index(tagValue[index+1:], "omitempty") != -1 && isEmptyValue(v) {
				continue
			}

			// If "squash" is specified in the tag, we squash the field down.
			squash = squash || strings.Index(tagValue[index+1:], "squash") != -1
			if squash {
				// When squashing, the embedded type can be a pointer to a struct.
				if v.Kind() == reflect.Ptr && v.Elem().Kind() == reflect.Struct {
					v = v.Elem()
				}

				// The final type must be a struct
				if v.Kind() != reflect.Struct {
					return fmt.Errorf("cannot squash non-struct type '%s'", v.Type())
				}
			}
			if keyNameTagValue := tagValue[:index]; keyNameTagValue != "" {
				keyName = keyNameTagValue
			}
		} else if len(tagValue) > 0 {
			if tagValue == "-" {
				continue
			}
			keyName = tagValue
		}

		switch v.Kind() {
		// this is an embedded struct, so handle it differently
		case reflect.Struct:
			x := reflect.New(v.Type())
			x.Elem().Set(v)

			vType := to.Type()
			vKeyType := vType.Key()
			vElemType := vType.Elem()
			mType := reflect.MapOf(vKeyType, vElemType)
			vMap := reflect.MakeMap(mType)

			// Creating a pointer to a map so that other methods can completely
			// overwrite the map if need be (looking at you decodeMapFromMap). The
			// indirection allows the underlying map to be settable (CanSet() == true)
			// where as reflect.MakeMap returns an unsettable map.
			addrVal := reflect.New(vMap.Type())
			reflect.Indirect(addrVal).Set(vMap)

			err := MapToValueWithOption(x.Interface(), reflect.Indirect(addrVal), option)
			if err != nil {
				return err
			}

			// the underlying map may have been completely overwritten so pull
			// it indirectly out of the enclosing value.
			vMap = reflect.Indirect(addrVal)

			if squash {
				for _, k := range vMap.MapKeys() {
					to.SetMapIndex(k, vMap.MapIndex(k))
				}
			} else {
				to.SetMapIndex(reflect.ValueOf(keyName), vMap)
			}

		default:
			to.SetMapIndex(reflect.ValueOf(keyName), v)
		}
	}

	return nil
}

func struct2structMapper(from reflect.Value, to reflect.Value, option *Option) error {
	mapType := reflect.TypeOf((map[string]any)(nil))
	mapValue := reflect.MakeMap(mapType)

	if err := struct2mapMapper(from, mapValue, option); err != nil {
		return err
	}

	return map2structMapper(mapValue, to, option)
}
