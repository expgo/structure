package structure

import (
	"reflect"
	"testing"
)

func TestFloat2BoolMapper(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		from     reflect.Value
		expected bool
	}{
		{
			name:     "MapZeroFloat",
			from:     reflect.ValueOf(0.0),
			expected: false,
		},
		{
			name:     "MapNegativeFloat",
			from:     reflect.ValueOf(-23.32),
			expected: true,
		},
		{
			name:     "MapPositiveFloat",
			from:     reflect.ValueOf(2.3),
			expected: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var to = reflect.ValueOf(new(bool)).Elem()
			err := float2boolMapper(tt.from, to)
			if err != nil {
				t.Errorf("float2boolMapper() error = %v", err)
				return
			}
			if to.Bool() != tt.expected {
				t.Errorf("float2boolMapper() = %v, expected %v", to.Bool(), tt.expected)
			}
		})
	}
}

func TestFloat2IntMapper(t *testing.T) {
	type args struct {
		from reflect.Value
		to   reflect.Value
	}
	int1 := int64(1)
	float1 := float64(1)

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "TestWithPositiveNumber",
			args: args{
				from: reflect.ValueOf(float1),
				to:   reflect.ValueOf(&int1).Elem(),
			},
			wantErr: false,
		},
		{
			name: "TestWithNegativeNumber",
			args: args{
				from: reflect.ValueOf(-float1),
				to:   reflect.ValueOf(&int1).Elem(),
			},
			wantErr: false,
		},
		{
			name: "TestWithZero",
			args: args{
				from: reflect.ValueOf(float64(0)),
				to:   reflect.ValueOf(&int1).Elem(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := float2intMapper(tt.args.from, tt.args.to); (err != nil) != tt.wantErr {
				t.Errorf("float2intMapper() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFloat2uintMapper(t *testing.T) {
	tests := []struct {
		name  string
		input float64
		want  uint64
		err   error
	}{
		{
			name:  "PositiveFloat",
			input: 123.45,
			want:  123,
			err:   nil,
		},
		{
			name:  "ZeroFloat",
			input: 0.0,
			want:  0,
			err:   nil,
		},
		{
			name:  "NegativeFloat",
			input: -123.45,
			want:  18446744073709551493, // The result of uint64 negative conversion
			err:   nil,
		},
		{
			name:  "MaxFloat",
			input: 1.7976931348623157e+308,
			want:  9223372036854775808, // The result of uint64 max float conversion
			err:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			from := reflect.ValueOf(tt.input)
			to := reflect.New(reflect.TypeOf(tt.want)).Elem()

			err := float2uintMapper(from, to)

			if err != tt.err {
				t.Errorf("float2uintMapper() error = %v, wantErr %v", err, tt.err)
				return
			}

			if to.Uint() != tt.want {
				t.Errorf("float2uintMapper() = %v, want %v", to.Uint(), tt.want)
			}
		})
	}
}

func TestFloat2FloatMapper(t *testing.T) {
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
			name: "SimpleTest1",
			args: args{
				from: reflect.ValueOf(3.1415),
				to:   reflect.ValueOf(new(float64)).Elem(),
			},
			wantErr: false,
		},
		{
			name: "SimpleTest2",
			args: args{
				from: reflect.ValueOf(0.0),
				to:   reflect.ValueOf(new(float64)).Elem(),
			},
			wantErr: false,
		},
		{
			name: "NegativeTest1",
			args: args{
				from: reflect.ValueOf(-2.7182),
				to:   reflect.ValueOf(new(float64)).Elem(),
			},
			wantErr: false,
		},
		{
			name: "NegativeInfinityTest",
			args: args{
				from: reflect.ValueOf(float64(-1) / float64(0.5)),
				to:   reflect.ValueOf(new(float64)).Elem(),
			},
			wantErr: false,
		},
		{
			name: "PositiveInfinityTest",
			args: args{
				from: reflect.ValueOf(float64(1) / float64(0.5)),
				to:   reflect.ValueOf(new(float64)).Elem(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := float2floatMapper(tt.args.from, tt.args.to); (err != nil) != tt.wantErr {
				t.Errorf("float2floatMapper() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got, want := tt.args.to.Float(), tt.args.from.Float(); got != want {
				t.Errorf("float2floatMapper() = %v, want %v", got, want)
			}
		})
	}
}

func TestFloat2stringMapper(t *testing.T) {
	tests := []struct {
		name  string
		float float64
		want  string
	}{
		{"Zero value", 0.0, "0"},
		{"Positive value", 10.0, "10"},
		{"Negative value", -10.0, "-10"},
		{"Fractional value", 2.5, "2.5"},
		{"Small value", 0.0000001, "1e-07"},
		{"Large value", 1e50, "1e+50"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			from := reflect.ValueOf(tt.float)
			to := reflect.New(reflect.TypeOf(""))

			err := float2stringMapper(from, to.Elem())
			if err != nil {
				t.Errorf("float2stringMapper() error = %v", err)
				return
			}

			if to.Elem().String() != tt.want {
				t.Errorf("float2stringMapper() got = %v, want = %v", to.Elem().String(), tt.want)
			}
		})
	}
}
