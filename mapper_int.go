package structure

import (
	"reflect"
	"strconv"
)

func init() {
	registerMapper(intType, boolType, int2boolMapper)
	registerMapper(intType, intType, int2intMapper)
	registerMapper(intType, uintType, int2uintMapper)
	registerMapper(intType, float64Type, int2floatMapper)
	registerMapper(intType, stringType, int2stringMapper)
}

func int2boolMapper(from reflect.Value, to reflect.Value, _ *Option) error {
	to.SetBool(from.Int() != 0)
	return nil
}

func int2intMapper(from reflect.Value, to reflect.Value, _ *Option) error {
	to.SetInt(from.Int())
	return nil
}

func int2uintMapper(from reflect.Value, to reflect.Value, _ *Option) error {
	to.SetUint(uint64(from.Int()))
	return nil
}

func int2floatMapper(from reflect.Value, to reflect.Value, _ *Option) error {
	to.SetFloat(float64(from.Int()))
	return nil
}

func int2stringMapper(from reflect.Value, to reflect.Value, _ *Option) error {
	to.SetString(strconv.Itoa(int(from.Int())))
	return nil
}
