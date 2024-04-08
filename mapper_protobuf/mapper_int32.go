package mapper_protobuf

import (
	"github.com/expgo/structure"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"reflect"
)

func init() {
	structure.RegisterMapper[wrapperspb.Int32Value, int32](int32value2intMapper)
	structure.RegisterMapper[int32, wrapperspb.Int32Value](int2int32valueMapper)
	structure.RegisterMapper[wrapperspb.Int32Value, int16](int32value2intMapper)
	structure.RegisterMapper[int16, wrapperspb.Int32Value](int2int32valueMapper)
	structure.RegisterMapper[wrapperspb.Int32Value, int8](int32value2intMapper)
	structure.RegisterMapper[int8, wrapperspb.Int32Value](int2int32valueMapper)
}

func int32value2intMapper(from reflect.Value, to reflect.Value, _ *structure.Option) error {
	to.SetInt(int64(from.Interface().(wrapperspb.Int32Value).Value))
	return nil
}

func int2int32valueMapper(from reflect.Value, to reflect.Value, _ *structure.Option) error {
	to.FieldByName("Value").SetInt(from.Int())
	return nil
}
