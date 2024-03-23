package structure

import (
	"errors"
	"fmt"
	"reflect"
)

func Must[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}

func ConvertTo[T any](from any) (T, error) {
	return ConvertToWithOption[T](from, defaultMapOption)
}

func MustConvertTo[T any](from any) T {
	return Must(ConvertTo[T](from))
}

func ConvertToWithOption[T any](from any, option *MapOption) (t T, err error) {
	result, err := ConvertToTypeWithOption(from, reflect.TypeOf((*T)(nil)).Elem(), option)
	if err != nil {
		return t, err
	} else {
		return result.(T), nil
	}
}

func ConvertToType(from any, toType reflect.Type) (any, error) {
	return ConvertToTypeWithOption(from, toType, defaultMapOption)
}

func MustConvertToType(from any, toType reflect.Type) any {
	return Must(ConvertToType(from, toType))
}

func ConvertToTypeWithOption(from any, toType reflect.Type, option *MapOption) (any, error) {
	if option == nil {
		option = defaultMapOption
	}

	result := reflect.New(toType).Elem()
	err := MapToValueWithOption(from, result, option)
	if err != nil {
		return nil, err
	}
	return result.Interface(), nil
}

func ConvertToKind(from any, toKind reflect.Kind) (any, error) {
	return ConvertToKindWithOption(from, toKind, defaultMapOption)
}

func MustConvertToKind(from any, toKind reflect.Kind) any {
	return Must(ConvertToKind(from, toKind))
}

func ConvertToKindWithOption(from any, toKind reflect.Kind, option *MapOption) (any, error) {
	if option == nil {
		option = defaultMapOption
	}

	switch toKind {
	case reflect.Bool:
		return ConvertTo[bool](from)
	case reflect.Int:
		return ConvertTo[int](from)
	case reflect.Int8:
		return ConvertTo[int8](from)
	case reflect.Int16:
		return ConvertTo[int16](from)
	case reflect.Int32:
		return ConvertTo[int32](from)
	case reflect.Int64:
		return ConvertTo[int64](from)
	case reflect.Uint:
		return ConvertTo[uint](from)
	case reflect.Uint8:
		return ConvertTo[uint8](from)
	case reflect.Uint16:
		return ConvertTo[uint16](from)
	case reflect.Uint32:
		return ConvertTo[uint32](from)
	case reflect.Uint64:
		return ConvertTo[uint64](from)
	case reflect.Float32:
		return ConvertTo[float32](from)
	case reflect.Float64:
		return ConvertTo[float64](from)
	case reflect.String:
		return ConvertTo[string](from)
	default:
		return nil, errors.New(fmt.Sprintf("unsupported kind %s", toKind))
	}
}
