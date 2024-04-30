package structure

import "reflect"

var (
	boolType    = reflect.TypeOf((*bool)(nil)).Elem()
	intType     = reflect.TypeOf((*int)(nil)).Elem()
	uintType    = reflect.TypeOf((*uint)(nil)).Elem()
	float64Type = reflect.TypeOf((*float64)(nil)).Elem()
	stringType  = reflect.TypeOf((*string)(nil)).Elem()
)
