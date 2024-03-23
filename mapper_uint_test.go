package structure

import (
	"reflect"
	"strconv"
	"testing"
)

func TestUint2boolMapper(t *testing.T) {
	tests := []struct {
		name   string
		input  uint64
		output bool
	}{
		{"Zero", 0, false},
		{"Positive", 1, true},
		{"LargePositive", 999999, true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			from := reflect.ValueOf(test.input)
			to := reflect.New(reflect.TypeOf(test.output)).Elem()

			err := uint2boolMapper(from, to)
			if err != nil {
				t.Errorf("got unexpected error: %v", err)
			}

			if to.Bool() != test.output {
				t.Errorf("expected %v, got %v", test.output, to.Bool())
			}
		})
	}
}

func TestUint2intMapper(t *testing.T) {
	// table driven tests
	var tests = []struct {
		uname string // test case name
		uin   uint64 // unsigned integer input for the function
		exp   int64  // expected result after conversion
	}{
		{"zero", 0, 0},
		{"one", 1, 1},
		{"maxUint64", ^uint64(0), -1}, // max value possible for a uint64
	}

	for _, tt := range tests {
		t.Run(tt.uname, func(t *testing.T) {
			inp := reflect.ValueOf(tt.uin)
			ex := reflect.New(reflect.TypeOf(tt.exp)).Elem()
			err := uint2intMapper(inp, ex)

			if err != nil {
				t.Fatalf("Expected no error, but got %s", err.Error())
			}

			if ex.Int() != tt.exp {
				t.Fatalf("Expected int64 value %d, but got %d", tt.exp, ex.Int())
			}
		})
	}
}

func TestUint2uintMapper(t *testing.T) {
	// Define the test cases
	testCases := []struct {
		name    string
		from    uint64
		to      uint64
		want    uint64
		wantErr bool
	}{
		{
			name:    "simple",
			from:    5,
			to:      0,
			want:    5,
			wantErr: false,
		},
		{
			name:    "maxUint64",
			from:    ^uint64(0), // Max value
			to:      0,
			want:    ^uint64(0),
			wantErr: false,
		},
		{
			name:    "minUint64",
			from:    0,          // Min value
			to:      ^uint64(0), // Max value
			want:    0,
			wantErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create reflect.value items
			fromVal := reflect.ValueOf(tc.from)
			toVal := reflect.ValueOf(&tc.to).Elem()

			if err := uint2uintMapper(fromVal, toVal); (err != nil) != tc.wantErr {
				t.Errorf("uint2uintMapper() error = %v, wantErr %v", err, tc.wantErr)
			}
			if got := toVal.Uint(); got != tc.want {
				t.Errorf("uint2uintMapper() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestUint2FloatMapper(t *testing.T) {
	type args struct {
		from reflect.Value
		to   reflect.Value
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "For zero uint",
			args: args{
				from: reflect.ValueOf(uint(0)),
				to:   reflect.ValueOf(new(float64)).Elem(),
			},
			wantErr: false,
		},
		{
			name: "For positive uint",
			args: args{
				from: reflect.ValueOf(uint(42)),
				to:   reflect.ValueOf(new(float64)).Elem(),
			},
			wantErr: false,
		},
		{
			name: "For max uint",
			args: args{
				from: reflect.ValueOf(uint(^uint(0))),
				to:   reflect.ValueOf(new(float64)).Elem(),
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt // capture range variable
		t.Run(tt.name, func(t *testing.T) {
			if err := uint2floatMapper(tt.args.from, tt.args.to); (err != nil) != tt.wantErr {
				t.Errorf("uint2floatMapper() error = %v, wantErr %v", err, tt.wantErr)
			}

			got := tt.args.to.Float()
			want := float64(tt.args.from.Uint())
			if got != want {
				t.Errorf("uint2floatMapper() got = %v, want %v", got, want)
			}
		})
	}
}

func TestUint2StringMapper(t *testing.T) {
	var testCases = []struct {
		name          string
		input         reflect.Value
		expectedValue string
		shouldError   bool
	}{
		{"zero", reflect.ValueOf(uint(0)), "0", false},
		{"positive", reflect.ValueOf(uint(12345)), "12345", false},
		{"maximum", reflect.ValueOf(uint(^uint(0))), strconv.FormatUint(uint64(^uint(0)), 10), false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var str string
			err := uint2stringMapper(tc.input, reflect.ValueOf(&str).Elem())

			if (err != nil) != tc.shouldError {
				t.Fatalf("Received unexpected error: '%v'", err)
			}

			if str != tc.expectedValue {
				t.Errorf("Expected '%v', got '%v'", tc.expectedValue, str)
			}
		})
	}
}
