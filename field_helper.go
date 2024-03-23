package structure

import (
	"errors"
	"reflect"
)

type fieldHelper struct {
	rv  reflect.Value
	Err error
}

func FieldHelper(s any) *fieldHelper {
	ret := &fieldHelper{}

	ret.rv = reflect.ValueOf(s)
	// 如果传入的是指针，需要使用Elem()来获取实际的值
	if ret.rv.Kind() == reflect.Ptr {
		ret.rv = ret.rv.Elem()
	}

	// 确保这是一个结构体
	if ret.rv.Kind() != reflect.Struct {
		ret.Err = errors.New("input is not struct")
	}

	return ret
}

func (fh *fieldHelper) GetValue2Do(fieldName string, toDo func(fieldName string, value any) bool) bool {
	if fh.Err != nil {
		return false
	}

	if toDo == nil {
		panic("toDo function cannot be empty")
	}

	fv := fh.rv.FieldByName(fieldName)

	if !fv.IsValid() {
		return false
	}

	return toDo(fieldName, fv.Interface())
}
