package structure

import (
	"reflect"
	"strconv"
)

func init() {
	registerMapper(uintType, boolType, uint2boolMapper)
	registerMapper(uintType, intType, uint2intMapper)
	registerMapper(uintType, uintType, uint2uintMapper)
	registerMapper(uintType, float64Type, uint2floatMapper)
	registerMapper(uintType, stringType, uint2stringMapper)
}

func uint2boolMapper(from reflect.Value, to reflect.Value, _ *Option) error {
	to.SetBool(from.Uint() != 0)
	return nil
}

func uint2intMapper(from reflect.Value, to reflect.Value, _ *Option) error {
	to.SetInt(int64(from.Uint()))
	return nil
}

func uint2uintMapper(from reflect.Value, to reflect.Value, _ *Option) error {
	to.SetUint(from.Uint())
	return nil
}

func uint2floatMapper(from reflect.Value, to reflect.Value, _ *Option) error {
	to.SetFloat(float64(from.Uint()))
	return nil
}

func uint2stringMapper(from reflect.Value, to reflect.Value, _ *Option) error {
	to.SetString(strconv.FormatUint(from.Uint(), 10))
	return nil
}
