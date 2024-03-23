package structure

import "reflect"

func init() {
	RegisterMapper[bool, bool](bool2boolMapper)
	RegisterMapper[bool, int](bool2intMapper)
	RegisterMapper[bool, uint](bool2uintMapper)
	RegisterMapper[bool, float64](bool2floatMapper)
	RegisterMapper[bool, string](bool2stringMapper)
}

func bool2boolMapper(from reflect.Value, to reflect.Value) error {
	to.SetBool(from.Bool())
	return nil
}

func bool2intMapper(from reflect.Value, to reflect.Value) error {
	if from.Bool() {
		to.SetInt(1)
	} else {
		to.SetInt(0)
	}
	return nil
}

func bool2uintMapper(from reflect.Value, to reflect.Value) error {
	if from.Bool() {
		to.SetUint(1)
	} else {
		to.SetUint(0)
	}
	return nil
}

func bool2floatMapper(from reflect.Value, to reflect.Value) error {
	if from.Bool() {
		to.SetFloat(1)
	} else {
		to.SetFloat(0)
	}
	return nil
}

func bool2stringMapper(from reflect.Value, to reflect.Value) error {
	if from.Bool() {
		to.SetString("true")
	} else {
		to.SetString("false")
	}
	return nil
}
