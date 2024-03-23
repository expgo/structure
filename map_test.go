package structure

import (
	"testing"
)

type Foo struct {
	Name string
}

type Bar struct {
	Name string
}

func TestAddTypeAliasMap(t *testing.T) {
	tests := []struct {
		name      string
		runFunc   func()
		wantPanic bool
	}{
		{
			name:      "Adds valid alias successfully",
			runFunc:   func() { AddTypeAliasMap[Foo, Bar]() },
			wantPanic: false,
		},
		{
			name:      "Adding the same alias twice triggers panic",
			runFunc:   func() { AddTypeAliasMap[Foo, Bar]() },
			wantPanic: true,
		},
		{
			name:      "Add different aliases to the same type",
			runFunc:   func() { AddTypeAliasMap[Foo, Bar]() },
			wantPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantPanic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("AddTypeAliasMap should panic when adding the same alias twice")
					}
				}()
			} else {
				defer func() {
					if r := recover(); r != nil {
						t.Errorf("Got unexpected panic: %v", r)
					}
				}()
			}
			tt.runFunc()
		})
	}
}

func TestMap(t *testing.T) {
	var vint int
	err := Map("123", &vint)
	if err != nil {
		t.Error(err)
	}
	if vint != 123 {
		t.Error("map err")
	}
}
