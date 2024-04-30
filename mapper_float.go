package structure

import (
	"reflect"
	"strconv"
)

func init() {
	registerMapper(float64Type, boolType, float2boolMapper)
	registerMapper(float64Type, intType, float2intMapper)
	registerMapper(float64Type, uintType, float2uintMapper)
	registerMapper(float64Type, float64Type, float2floatMapper)
	registerMapper(float64Type, stringType, float2stringMapper)
}

func float2boolMapper(from reflect.Value, to reflect.Value, _ *Option) error {
	to.SetBool(from.Float() != 0)
	return nil
}

func float2intMapper(from reflect.Value, to reflect.Value, _ *Option) error {
	to.SetInt(int64(from.Float()))
	return nil
}

func float2uintMapper(from reflect.Value, to reflect.Value, _ *Option) error {
	to.SetUint(uint64(from.Float()))
	return nil
}

func float2floatMapper(from reflect.Value, to reflect.Value, _ *Option) error {
	to.SetFloat(from.Float())
	return nil
}

func float2stringMapper(from reflect.Value, to reflect.Value, _ *Option) error {
	to.SetString(strconv.FormatFloat(from.Float(), 'g', -1, 64))
	return nil
}
