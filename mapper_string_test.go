package structure

import (
	"reflect"
	"testing"
)

func TestString2boolMapper(t *testing.T) {
	tests := []struct {
		name    string
		from    string
		want    bool
		wantErr bool
	}{
		{
			name: "TrueLowerCase",
			from: "true",
			want: true,
		},
		{
			name: "FalseLowerCase",
			from: "false",
			want: false,
		},
		{
			name: "TrueUpperCase",
			from: "TRUE",
			want: true,
		},
		{
			name: "FalseUpperCase",
			from: "FALSE",
			want: false,
		},
		{
			name:    "RandomString",
			from:    "RandomString",
			want:    false,
			wantErr: true,
		},
		{
			name: "EmptyString",
			from: "",
			want: false,
		},
		{
			name: "WhiteSpaceString",
			from: "     ",
			want: false,
		},
		{
			name: "MixedCaseTrue",
			from: "t",
			want: true,
		},
		{
			name: "MixedCaseFalse",
			from: "f",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			from := reflect.ValueOf(tt.from)
			var to reflect.Value
			var toType bool
			to = reflect.ValueOf(&toType).Elem()
			err := string2boolMapper(from, to)

			if (err != nil) != tt.wantErr {
				t.Errorf("string2boolMapper() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if to.Bool() != tt.want {
				t.Errorf("string2boolMapper() = %v, want %v", to.Bool(), tt.want)
			}
		})
	}
}

func TestString2intMapper(t *testing.T) {
	tests := []struct {
		name  string
		from  reflect.Value
		to    reflect.Value
		err   bool
		value int64
	}{
		{
			name:  "Test regular conversion",
			from:  reflect.ValueOf("5"),
			to:    reflect.ValueOf(new(int64)).Elem(),
			err:   false,
			value: 5,
		},
		{
			name:  "Test empty string conversion",
			from:  reflect.ValueOf(" "),
			to:    reflect.ValueOf(new(int64)).Elem(),
			err:   false,
			value: 0,
		},
		{
			name:  "Test invalid string conversion",
			from:  reflect.ValueOf("AA"),
			to:    reflect.ValueOf(new(int64)).Elem(),
			err:   true,
			value: 0,
		},
		{
			name:  "Test extremely large number",
			from:  reflect.ValueOf("9223372036854775807"),
			to:    reflect.ValueOf(new(int64)).Elem(),
			err:   false,
			value: 9223372036854775807,
		},
		{
			name:  "Test number too large to hold",
			from:  reflect.ValueOf("9223372036854775808"),
			to:    reflect.ValueOf(new(int64)).Elem(),
			err:   true,
			value: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := string2intMapper(tc.from, tc.to)

			if tc.err {
				if err == nil {
					t.Errorf("%s: expected error but got no error", tc.name)
				}
				return
			}

			if err != nil {
				t.Errorf("%s: unexpected error - %s", tc.name, err)
				return
			}

			if tc.to.Int() != tc.value {
				t.Errorf("%s: got %d but expected %d", tc.name, tc.to.Int(), tc.value)
			}
		})
	}
}

func Test_string2uintMapper(t *testing.T) {
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
			name: "EmptyString",
			args: args{
				from: reflect.ValueOf(""),
				to:   reflect.ValueOf(new(uint64)).Elem(),
			},
			wantErr: false,
		},
		{
			name: "WhitespaceString",
			args: args{
				from: reflect.ValueOf("  \t "),
				to:   reflect.ValueOf(new(uint64)).Elem(),
			},
			wantErr: false,
		},
		{
			name: "ValidUintString",
			args: args{
				from: reflect.ValueOf("123456"),
				to:   reflect.ValueOf(new(uint64)).Elem(),
			},
			wantErr: false,
		},
		{
			name: "InvalidUintString",
			args: args{
				from: reflect.ValueOf("abc"),
				to:   reflect.ValueOf(new(uint64)).Elem(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := string2uintMapper(tt.args.from, tt.args.to); (err != nil) != tt.wantErr {
				t.Errorf("string2uintMapper() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestString2floatMapper(t *testing.T) {
	var tests = []struct {
		name      string
		from      reflect.Value
		to        reflect.Value
		want      float64
		expectErr bool
	}{
		{
			name:      "normal conversion",
			from:      reflect.ValueOf("123"),
			to:        reflect.ValueOf(new(float64)).Elem(),
			want:      123,
			expectErr: false,
		},
		{
			name:      "conversion with space",
			from:      reflect.ValueOf("    -456    "),
			to:        reflect.ValueOf(new(float64)).Elem(),
			want:      -456,
			expectErr: false,
		},
		{
			name:      "removing trailing zero",
			from:      reflect.ValueOf("789.500"),
			to:        reflect.ValueOf(new(float64)).Elem(),
			want:      789.5,
			expectErr: false,
		},
		{
			name:      "empty conversion",
			from:      reflect.ValueOf("   "),
			to:        reflect.ValueOf(new(float64)).Elem(),
			want:      0,
			expectErr: false,
		},
		{
			name:      "errorful conversion",
			from:      reflect.ValueOf("abc"),
			to:        reflect.ValueOf(new(float64)).Elem(),
			expectErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := string2floatMapper(tt.from, tt.to)
			if tt.expectErr {
				if err == nil {
					t.Errorf("expected error but received none")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: got %v", err)
				}
				if got := tt.to.Float(); got != tt.want {
					t.Errorf("got: %v, want: %v", got, tt.want)
				}
			}
		})
	}
}

func TestString2StringMapper(t *testing.T) {
	tests := []struct {
		name string
		from string
		want string
	}{
		{"EmptyStringTest", "", ""},
		{"SingleWordTest", "word", "word"},
		{"MultiWordTest", "multiple words", "multiple words"},
		{"SpecialCharTest", "!$^&*()", "!$^&*()"},
		{"MixedCharTest", "Mix3d Ch@r$", "Mix3d Ch@r$"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			from := reflect.ValueOf(tt.from)
			to := reflect.New(reflect.TypeOf(tt.from)).Elem()

			err := string2stringMapper(from, to)

			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			got := to.String()
			if got != tt.want {
				t.Errorf("string2stringMapper() = %v, want %v", got, tt.want)
			}
		})
	}
}
