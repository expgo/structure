package structure

import (
	"reflect"
	"testing"
)

func Test_bool2boolMapper(t *testing.T) {
	type args struct {
		from reflect.Value
		to   reflect.Value
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"BoolToBoolTrue",
			args{from: reflect.ValueOf(true), to: reflect.ValueOf(new(bool)).Elem()},
			false},
		{"BoolToBoolFalse",
			args{from: reflect.ValueOf(false), to: reflect.ValueOf(new(bool)).Elem()},
			false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := bool2boolMapper(tt.args.from, tt.args.to); (err != nil) != tt.wantErr {
				t.Errorf("bool2boolMapper() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got := tt.args.to.Bool(); got != tt.args.from.Bool() {
				t.Errorf("bool2boolMapper() return = %v, want %v", got, tt.args.from.Bool())
			}
		})
	}
}

func TestBool2intMapper(t *testing.T) {
	testCases := []struct {
		name  string
		input bool
		want  int64
	}{
		{"BoolTrue", true, 1},
		{"BoolFalse", false, 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			have := reflect.New(reflect.TypeOf(tc.want)).Elem()
			err := bool2intMapper(reflect.ValueOf(tc.input), have)
			if err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}

			if have.Int() != tc.want {
				t.Errorf("Got %d; wanted %d", have.Int(), tc.want)
			}
		})
	}
}

func Test_bool2uintMapper(t *testing.T) {
	tests := []struct {
		name    string
		from    reflect.Value
		to      reflect.Value
		wantErr bool
	}{
		{
			name:    "Test Bool True",
			from:    reflect.ValueOf(true),
			to:      reflect.ValueOf(new(uint64)).Elem(),
			wantErr: false,
		},
		{
			name:    "Test Bool False",
			from:    reflect.ValueOf(false),
			to:      reflect.ValueOf(new(uint64)).Elem(),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := bool2uintMapper(tt.from, tt.to); (err != nil) != tt.wantErr {
				t.Errorf("bool2uintMapper() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.from.Bool() && tt.to.Uint() != 1 {
				t.Errorf("bool2uintMapper() to = %v, want 1", tt.to.Uint())
			}

			if !tt.from.Bool() && tt.to.Uint() != 0 {
				t.Errorf("bool2uintMapper() to = %v, want 0", tt.to.Uint())
			}
		})
	}
}

func TestBool2FloatMapper(t *testing.T) {
	tests := []struct {
		name  string
		from  bool
		want  float64
		error error
	}{
		{
			name: "Test Case True",
			from: true,
			want: 1,
		},
		{
			name: "Test Case False",
			from: false,
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			from := reflect.ValueOf(tt.from)
			to := reflect.New(reflect.TypeOf(tt.want))
			err := bool2floatMapper(from, to.Elem())

			if err != tt.error {
				t.Errorf("bool2floatMapper() error = %v, wantErr %v",
					err, tt.error)
				return
			}

			if to.Elem().Interface() != tt.want {
				t.Errorf("Expected and output is not the same %v %v",
					to.Elem().Interface(), tt.want)
			}
		})
	}
}

func Test_bool2stringMapper(t *testing.T) {

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
			name: "Bool to String - True",
			args: args{
				from: reflect.ValueOf(true),
				to:   reflect.ValueOf(new(string)).Elem(),
			},
			wantErr: false,
		},

		{
			name: "Bool to String - False",
			args: args{
				from: reflect.ValueOf(false),
				to:   reflect.ValueOf(new(string)).Elem(),
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := bool2stringMapper(tt.args.from, tt.args.to); (err != nil) != tt.wantErr {
				t.Errorf("bool2stringMapper() error = %v, wantErr %v", err, tt.wantErr)
			}

			got := tt.args.to.String()

			var want string
			if tt.args.from.Bool() {
				want = "true"
			} else {
				want = "false"
			}

			if got != want {
				t.Errorf("bool2stringMapper() got = %v, want %v", got, want)
			}
		})
	}
}
