package structure

import (
	"reflect"
	"strconv"
	"testing"
)

func TestInt2BoolMapper(t *testing.T) {
	tests := []struct {
		name  string
		input int64
		want  bool
	}{
		{"zero", 0, false},
		{"positive", 2, true},
		{"negative", -2, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			from := reflect.ValueOf(tt.input)
			var to bool
			toV := reflect.ValueOf(&to).Elem()
			err := int2boolMapper(from, toV)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}
			if to != tt.want {
				t.Errorf("Unexpected value: got %v, want %v", to, tt.want)
			}
		})
	}
}

func TestInt2IntMapper(t *testing.T) {
	tests := []struct {
		name   string
		from   int64
		to     int64
		expect int64
	}{
		{"positive value", 10, 0, 10},
		{"negative value", -5, 0, -5},
		{"zero", 0, 0, 0},
		{"boundary value", 9223372036854775807, 0, 9223372036854775807},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fromVal := reflect.ValueOf(tt.from)
			toVal := reflect.ValueOf(&tt.to).Elem()
			err := int2intMapper(fromVal, toVal)
			if err != nil {
				t.Errorf("Got error %v", err)
				return
			}
			if toVal.Int() != tt.expect {
				t.Errorf("Expected %d, got %d", tt.expect, toVal.Int())
			}
		})
	}
}

func TestInt2UintMapper(t *testing.T) {
	value := int64(-10)
	tests := []struct {
		name    string
		from    int64
		want    uint64
		wantErr bool
	}{
		{
			name:    "Positive Case",
			from:    10,
			want:    10,
			wantErr: false,
		},
		{
			name:    "Zero Case",
			from:    0,
			want:    0,
			wantErr: false,
		},
		{
			name:    "Negative Case",
			from:    -10,
			want:    uint64(value), // This will cause an overflow, resulting into a big positive number
			wantErr: false,
		},
		{
			name:    "Max Int64 Case",
			from:    1<<63 - 1,
			want:    1<<63 - 1,
			wantErr: false,
		},
		{
			name:    "Min Int64 Case",
			from:    -1 << 63,
			want:    1 << 63,
			wantErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			from := reflect.ValueOf(tc.from)
			to := reflect.New(reflect.TypeOf(tc.want)).Elem()
			gotErr := int2uintMapper(from, to)
			if (gotErr != nil) != tc.wantErr {
				t.Errorf("int2uintMapper() error = %v, wantErr %v", gotErr, tc.wantErr)
				return
			}
			if got := to.Uint(); got != tc.want {
				t.Errorf("int2uintMapper() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestInt2floatMapper(t *testing.T) {
	cases := []struct {
		name     string
		input    int64
		expected float64
	}{
		{
			name:     "Zero",
			input:    0,
			expected: 0,
		},
		{
			name:     "PositiveInteger",
			input:    10,
			expected: 10.0,
		},
		{
			name:     "NegativeInteger",
			input:    -10,
			expected: -10.0,
		},
		{
			name:     "MaxInt64",
			input:    9223372036854775807,
			expected: 9223372036854775807,
		},
		{
			name:     "MinInt64",
			input:    -9223372036854775808,
			expected: -9223372036854775808,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			from := reflect.ValueOf(tc.input)
			to := reflect.New(reflect.TypeOf(tc.expected)).Elem()

			err := int2floatMapper(from, to)
			if err != nil {
				t.Fatalf("expected no error but got: %v", err)
			}

			result := to.Float()
			if result != tc.expected {
				t.Errorf("expected %v but got %v", tc.expected, result)
			}
		})
	}
}

func TestInt2stringMapper(t *testing.T) {
	testCases := []struct {
		name     string
		from     reflect.Value
		expected reflect.Value
	}{
		{
			name:     "TestZero",
			from:     reflect.ValueOf(0),
			expected: reflect.ValueOf("0"),
		},
		{
			name:     "TestPositiveInt",
			from:     reflect.ValueOf(123),
			expected: reflect.ValueOf("123"),
		},
		{
			name:     "TestNegativeInt",
			from:     reflect.ValueOf(-456),
			expected: reflect.ValueOf("-456"),
		},
		{
			name:     "TestMaxInt",
			from:     reflect.ValueOf(int(^uint(0) >> 1)),
			expected: reflect.ValueOf(strconv.Itoa(int(^uint(0) >> 1))),
		},
		{
			name:     "TestMinInt",
			from:     reflect.ValueOf(-int(^uint(0)>>1) - 1),
			expected: reflect.ValueOf(strconv.Itoa(-int(^uint(0)>>1) - 1)),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			to := reflect.New(reflect.TypeOf("")).Elem()
			err := int2stringMapper(tc.from, to)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}
			if to.String() != tc.expected.String() {
				t.Errorf("%v: actual value %v, expected %v", tc.name, to, tc.expected)
			}
		})
	}
}
