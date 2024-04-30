package structure

import "reflect"

func init() {
	registerMapper(boolType, boolType, bool2boolMapper)
	registerMapper(boolType, intType, bool2intMapper)
	registerMapper(boolType, uintType, bool2uintMapper)
	registerMapper(boolType, float64Type, bool2floatMapper)
	registerMapper(boolType, stringType, bool2stringMapper)
}

func bool2boolMapper(from reflect.Value, to reflect.Value, _ *Option) error {
	to.SetBool(from.Bool())
	return nil
}

func bool2intMapper(from reflect.Value, to reflect.Value, _ *Option) error {
	if from.Bool() {
		to.SetInt(1)
	} else {
		to.SetInt(0)
	}
	return nil
}

func bool2uintMapper(from reflect.Value, to reflect.Value, _ *Option) error {
	if from.Bool() {
		to.SetUint(1)
	} else {
		to.SetUint(0)
	}
	return nil
}

func bool2floatMapper(from reflect.Value, to reflect.Value, _ *Option) error {
	if from.Bool() {
		to.SetFloat(1)
	} else {
		to.SetFloat(0)
	}
	return nil
}

func bool2stringMapper(from reflect.Value, to reflect.Value, _ *Option) error {
	if from.Bool() {
		to.SetString("true")
	} else {
		to.SetString("false")
	}
	return nil
}
