package structure

import (
	"fmt"
	"reflect"
	"sync"
)

var typeAliasMap = make(map[reflect.Type]reflect.Type)
var typeAliasMapLock = &sync.RWMutex{}

func init() {
	AddTypeAliasMap[int8, int]()
	AddTypeAliasMap[int16, int]()
	AddTypeAliasMap[int32, int]()
	AddTypeAliasMap[int64, int]()

	AddTypeAliasMap[uint8, uint]()
	AddTypeAliasMap[uint16, uint]()
	AddTypeAliasMap[uint32, uint]()
	AddTypeAliasMap[uint64, uint]()

	AddTypeAliasMap[float32, float64]()
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
