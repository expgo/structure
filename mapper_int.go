package structure

import (
	"reflect"
	"strconv"
)

func init() {
	RegisterMapper[int, bool](int2boolMapper)
	RegisterMapper[int, int](int2intMapper)
	RegisterMapper[int, uint](int2uintMapper)
	RegisterMapper[int, float64](int2floatMapper)
	RegisterMapper[int, string](int2stringMapper)
}

func int2boolMapper(from reflect.Value, to reflect.Value) error {
	to.SetBool(from.Int() != 0)
	return nil
}

func int2intMapper(from reflect.Value, to reflect.Value) error {
	to.SetInt(from.Int())
	return nil
}

func int2uintMapper(from reflect.Value, to reflect.Value) error {
	to.SetUint(uint64(from.Int()))
	return nil
}

func int2floatMapper(from reflect.Value, to reflect.Value) error {
	to.SetFloat(float64(from.Int()))
	return nil
}

func int2stringMapper(from reflect.Value, to reflect.Value) error {
	to.SetString(strconv.Itoa(int(from.Int())))
	return nil
}
