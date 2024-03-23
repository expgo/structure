package structure

import (
	"reflect"
	"strconv"
)

func init() {
	RegisterMapper[uint, bool](uint2boolMapper)
	RegisterMapper[uint, int](uint2intMapper)
	RegisterMapper[uint, uint](uint2uintMapper)
	RegisterMapper[uint, float64](uint2floatMapper)
	RegisterMapper[uint, string](uint2stringMapper)
}

func uint2boolMapper(from reflect.Value, to reflect.Value) error {
	to.SetBool(from.Uint() != 0)
	return nil
}

func uint2intMapper(from reflect.Value, to reflect.Value) error {
	to.SetInt(int64(from.Uint()))
	return nil
}

func uint2uintMapper(from reflect.Value, to reflect.Value) error {
	to.SetUint(from.Uint())
	return nil
}

func uint2floatMapper(from reflect.Value, to reflect.Value) error {
	to.SetFloat(float64(from.Uint()))
	return nil
}

func uint2stringMapper(from reflect.Value, to reflect.Value) error {
	to.SetString(strconv.FormatUint(from.Uint(), 10))
	return nil
}
