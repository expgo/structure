package structure

import (
	"fmt"
	"reflect"
	"sync"
)

var typeAliasMap = make(map[reflect.Type]reflect.Type)
var typeAliasMapLock = &sync.RWMutex{}

func init() {
	addTypeAliasMap(int8Type, intType)
	addTypeAliasMap(int16Type, intType)
	addTypeAliasMap(int32Type, intType)
	addTypeAliasMap(int64Type, intType)

	addTypeAliasMap(uint8Type, uintType)
	addTypeAliasMap(uint16Type, uintType)
	addTypeAliasMap(uint32Type, uintType)
	addTypeAliasMap(uint64Type, uintType)

	addTypeAliasMap(float32Type, float64Type)
}

func AddTypeAliasMap[Name any, Alias any]() {
	addTypeAliasMap(reflect.TypeOf((*Name)(nil)).Elem(), reflect.TypeOf((*Alias)(nil)).Elem())
}

func addTypeAliasMap(nameType reflect.Type, aliasType reflect.Type) {
	typeAliasMapLock.Lock()
	defer typeAliasMapLock.Unlock()

	if v, got := typeAliasMap[nameType]; got {
		panic(fmt.Sprintf("type '%v' already set alias to '%v'", nameType, v))
	} else {
		typeAliasMap[nameType] = aliasType
	}
}
