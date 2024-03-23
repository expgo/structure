package mapper_protobuf

import (
	"google.golang.org/protobuf/types/known/wrapperspb"
	"reflect"
	"testing"
)

func TestInt64value2intMapper(t *testing.T) {
	tests := []struct {
		name  string
		input wrapperspb.Int64Value
		want  any
	}{
		{"PositiveMapping", wrapperspb.Int64Value{Value: 5}, 5},
		{"NegativeMapping", wrapperspb.Int64Value{Value: -5}, -5},
		{"MappingWithZero", wrapperspb.Int64Value{Value: 0}, 0},
		{"MappingWithMaxInt64", wrapperspb.Int64Value{Value: 9223372036854775807}, 9223372036854775807},
		{"MappingWithMinInt64", wrapperspb.Int64Value{Value: -9223372036854775808}, -9223372036854775808},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			from := reflect.ValueOf(tt.input)
			var to int
			toV := reflect.ValueOf(&to).Elem()
			err := int64value2intMapper(from, toV)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}
			if to != tt.want {
				t.Errorf("Unexpected value: got %v, want %v", to, tt.want)
			}
		})
	}
}

func TestInt2int64valueMapper(t *testing.T) {
	tests := []struct {
		name  string
		want  wrapperspb.Int64Value
		input int
	}{
		{"PositiveMapping", wrapperspb.Int64Value{Value: 5}, 5},
		{"NegativeMapping", wrapperspb.Int64Value{Value: -5}, -5},
		{"MappingWithZero", wrapperspb.Int64Value{Value: 0}, 0},
		{"MappingWithMaxInt64", wrapperspb.Int64Value{Value: 9223372036854775807}, 9223372036854775807},
		{"MappingWithMinInt64", wrapperspb.Int64Value{Value: -9223372036854775808}, -9223372036854775808},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			from := reflect.ValueOf(tt.input)
			var to wrapperspb.Int64Value
			toV := reflect.ValueOf(&to).Elem()
			err := int2int64valueMapper(from, toV)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}
			if to.Value != tt.want.Value {
				t.Errorf("Unexpected value: got %v, want %v", to.Value, tt.want.Value)
			}
		})
	}
}
