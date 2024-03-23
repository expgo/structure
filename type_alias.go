package structure

import (
	"fmt"
	"github.com/expgo/generic"
	"reflect"
)

var typeAliasMap = generic.Map[reflect.Type, reflect.Type]{}

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
	nameType := reflect.TypeOf((*Name)(nil)).Elem()
	aliasType := reflect.TypeOf((*Alias)(nil)).Elem()

	if v, got := typeAliasMap.LoadOrStore(nameType, aliasType); got {
		panic(fmt.Sprintf("type '%v' already set alias to '%v'", nameType, v))
	}
}
