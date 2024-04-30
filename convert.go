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
	return ConvertToWithOption[T](from, defaultOption)
}

func MustConvertTo[T any](from any) T {
	return Must(ConvertTo[T](from))
}

func ConvertToWithOption[T any](from any, option *Option) (t T, err error) {
	result, err := ConvertToTypeWithOption(from, reflect.TypeOf((*T)(nil)).Elem(), option)
	if err != nil {
		return t, err
	} else {
		return result.(T), nil
	}
}

func ConvertToType(from any, toType reflect.Type) (any, error) {
	return ConvertToTypeWithOption(from, toType, defaultOption)
}

func MustConvertToType(from any, toType reflect.Type) any {
	return Must(ConvertToType(from, toType))
}

func ConvertToTypeWithOption(from any, toType reflect.Type, option *Option) (any, error) {
	if option == nil {
		option = defaultOption
	}

	result := reflect.New(toType).Elem()
	err := MapToValueWithOption(from, result, option)
	if err != nil {
		return nil, err
	}
	return result.Interface(), nil
}

func ConvertToKind(from any, toKind reflect.Kind) (any, error) {
	return ConvertToKindWithOption(from, toKind, defaultOption)
}

func MustConvertToKind(from any, toKind reflect.Kind) any {
	return Must(ConvertToKind(from, toKind))
}

func ConvertToKindWithOption(from any, toKind reflect.Kind, option *Option) (any, error) {
	if option == nil {
		option = defaultOption
	}

	switch toKind {
	case reflect.Bool:
		return ConvertToTypeWithOption(from, boolType, option)
	case reflect.Int:
		return ConvertToTypeWithOption(from, intType, option)
	case reflect.Int8:
		return ConvertToTypeWithOption(from, int8Type, option)
	case reflect.Int16:
		return ConvertToTypeWithOption(from, int16Type, option)
	case reflect.Int32:
		return ConvertToTypeWithOption(from, int32Type, option)
	case reflect.Int64:
		return ConvertToTypeWithOption(from, int64Type, option)
	case reflect.Uint:
		return ConvertToTypeWithOption(from, uintType, option)
	case reflect.Uint8:
		return ConvertToTypeWithOption(from, uint8Type, option)
	case reflect.Uint16:
		return ConvertToTypeWithOption(from, uint16Type, option)
	case reflect.Uint32:
		return ConvertToTypeWithOption(from, uint32Type, option)
	case reflect.Uint64:
		return ConvertToTypeWithOption(from, uint64Type, option)
	case reflect.Float32:
		return ConvertToTypeWithOption(from, float32Type, option)
	case reflect.Float64:
		return ConvertToTypeWithOption(from, float64Type, option)
	case reflect.String:
		return ConvertToTypeWithOption(from, stringType, option)
	default:
		return nil, errors.New(fmt.Sprintf("unsupported kind %s", toKind))
	}
}
