package mapper_protobuf

import (
	"github.com/expgo/structure"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"reflect"
)

func init() {
	structure.RegisterMapper[wrapperspb.Int64Value, int](int64value2intMapper)
	structure.RegisterMapper[int, wrapperspb.Int64Value](int2int64valueMapper)
	structure.RegisterMapper[wrapperspb.Int64Value, int64](int64value2intMapper)
	structure.RegisterMapper[int64, wrapperspb.Int64Value](int2int64valueMapper)
}

func int64value2intMapper(from reflect.Value, to reflect.Value) error {
	to.SetInt(from.Interface().(wrapperspb.Int64Value).Value)
	return nil
}

func int2int64valueMapper(from reflect.Value, to reflect.Value) error {
	to.FieldByName("Value").SetInt(from.Int())
	return nil
}
