package mapper_protobuf

import (
	"github.com/expgo/structure"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"reflect"
	"testing"
)

func TestMustConvertTo(t *testing.T) {
	//bTrue := true
	//bFalse := false
	int1 := 1
	int_32_5 := int32(5)
	//int0 := 0
	//uint1 := uint(1)
	//uint0 := uint(0)
	//float32_1 := float32(1.0)
	//float64_0 := 0.0
	//strTrue := "true"
	//strFalse := "false"
	//int_123 := -123
	//uint456 := uint(456)
	//float456 := 456.0
	//str123 := "123"
	//urlAddress, _ := url.Parse("https://www.example.com")
	//durationPeriod, _ := time.ParseDuration("1h30m")
	//time0128, _ := time.Parse("2006-01-02T15:04:05Z07:00", "2024-01-18T09:12:34+08:00")

	tests := []struct {
		name  string
		fn    func() any
		want  any
		panic bool
	}{
		// int to int64value and int64value to int
		{"int to int64value", func() any { return structure.MustConvertTo[wrapperspb.Int64Value](1) }, *wrapperspb.Int64(1), false},
		{"int to *int64value", func() any { return structure.MustConvertTo[*wrapperspb.Int64Value](1) }, wrapperspb.Int64(1), false},
		{"*int to int64value", func() any { return structure.MustConvertTo[wrapperspb.Int64Value](&int1) }, *wrapperspb.Int64(1), false},
		{"*int to *int64value", func() any { return structure.MustConvertTo[*wrapperspb.Int64Value](&int1) }, wrapperspb.Int64(1), false},
		{"*int to *int64value nil", func() any { return structure.MustConvertTo[*wrapperspb.Int64Value](nil) }, (*wrapperspb.Int64Value)(nil), false},
		{"int64value to int", func() any { return structure.MustConvertTo[int](*wrapperspb.Int64(1)) }, 1, false},
		{"int64value to *int", func() any { return structure.MustConvertTo[*int](*wrapperspb.Int64(1)) }, &int1, false},
		{"*int64value to int", func() any { return structure.MustConvertTo[int](wrapperspb.Int64(1)) }, 1, false},
		{"*int64value to *int", func() any { return structure.MustConvertTo[*int](wrapperspb.Int64(1)) }, &int1, false},
		{"*int64value to *int nil", func() any { return structure.MustConvertTo[*int](nil) }, (*int)(nil), false},

		// int32 to int32value and int32value to int32
		{"int32 to int32value", func() any { return structure.MustConvertTo[wrapperspb.Int32Value](int_32_5) }, *wrapperspb.Int32(5), false},
		{"int to *int64value", func() any { return structure.MustConvertTo[*wrapperspb.Int32Value](int_32_5) }, wrapperspb.Int32(5), false},
		{"*int to int64value", func() any { return structure.MustConvertTo[wrapperspb.Int32Value](&int_32_5) }, *wrapperspb.Int32(5), false},
		{"*int to *int64value", func() any { return structure.MustConvertTo[*wrapperspb.Int32Value](&int_32_5) }, wrapperspb.Int32(5), false},
		{"*int to *int64value nil", func() any { return structure.MustConvertTo[*wrapperspb.Int32Value](nil) }, (*wrapperspb.Int32Value)(nil), false},
		{"int64value to int", func() any { return structure.MustConvertTo[int32](*wrapperspb.Int32(5)) }, int_32_5, false},
		{"int64value to *int", func() any { return structure.MustConvertTo[*int32](*wrapperspb.Int32(5)) }, &int_32_5, false},
		{"*int64value to int", func() any { return structure.MustConvertTo[int32](wrapperspb.Int32(5)) }, int_32_5, false},
		{"*int64value to *int", func() any { return structure.MustConvertTo[*int32](wrapperspb.Int32(5)) }, &int_32_5, false},
		{"*int64value to *int nil", func() any { return structure.MustConvertTo[*int32](nil) }, (*int32)(nil), false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if !tc.panic {
						t.Errorf("The code panicked: %+v", r)
					}
				} else {
					if tc.panic {
						t.Errorf("The code did not panic")
					}
				}
			}()

			// Convert 'fn' to the same type as 'to'.
			// This may panic, depending on the test case.
			result := tc.fn()

			if !tc.panic {
				// Check the result.
				if !reflect.DeepEqual(result, tc.want) {
					t.Errorf("Want %v, but got %v", tc.want, result)
				}
			}
		})
	}
}
