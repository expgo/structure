package structure

import (
	"reflect"
	"strconv"
)

func init() {
	RegisterMapper[float64, bool](float2boolMapper)
	RegisterMapper[float64, int](float2intMapper)
	RegisterMapper[float64, uint](float2uintMapper)
	RegisterMapper[float64, float64](float2floatMapper)
	RegisterMapper[float64, string](float2stringMapper)
}

func float2boolMapper(from reflect.Value, to reflect.Value) error {
	to.SetBool(from.Float() != 0)
	return nil
}

func float2intMapper(from reflect.Value, to reflect.Value) error {
	to.SetInt(int64(from.Float()))
	return nil
}

func float2uintMapper(from reflect.Value, to reflect.Value) error {
	to.SetUint(uint64(from.Float()))
	return nil
}

func float2floatMapper(from reflect.Value, to reflect.Value) error {
	to.SetFloat(from.Float())
	return nil
}

func float2stringMapper(from reflect.Value, to reflect.Value) error {
	to.SetString(strconv.FormatFloat(from.Float(), 'g', -1, 64))
	return nil
}
