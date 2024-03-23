package structure

import (
	"net/url"
	"reflect"
	"testing"
	"time"
)

// copied from github.com/caarlos0/env
type unmarshaler struct {
	time.Duration
}

// TextUnmarshaler implements encoding.TextUnmarshaler.
func (d *unmarshaler) UnmarshalText(data []byte) (err error) {
	if len(data) != 0 {
		d.Duration, err = time.ParseDuration(string(data))
	} else {
		d.Duration = 0
	}
	return err
}

type PointStruct struct {
	NestedNonDefined struct {
		NonDefined struct {
			String string `env:"STR"`
		} `envPrefix:"NONDEFINED_"`
	} `envPrefix:"PRF_"`
}

// copied from github.com/caarlos0/env
type WireValueExprStruct struct {
	String          string    `wire:"value:${env.STRING}"`
	StringPtr       *string   `wire:"value:${env.STRING}"`
	Strings         []string  `wire:"value:${env.STRINGS}"`
	StringPtrs      []*string `wire:"value:${env.STRINGS}"`
	StringPtrInited *string   `wire:"value:${env.STRINGInited}"`

	Bool     bool    `wire:"value:${env.BOOL}"`
	BoolPtr  *bool   `wire:"value:${env.BOOL}"`
	Bools    []bool  `wire:"value:${env.BOOLS}"`
	BoolPtrs []*bool `wire:"value:${env.BOOLS}"`

	Int     int    `wire:"value:${env.INT}"`
	IntPtr  *int   `wire:"value:${env.INT}"`
	Ints    []int  `wire:"value:${env.INTS}"`
	IntPtrs []*int `wire:"value:${env.INTS}"`

	Int8     int8    `wire:"value:${env.INT8}"`
	Int8Ptr  *int8   `wire:"value:${env.INT8}"`
	Int8s    []int8  `wire:"value:${env.INT8S}"`
	Int8Ptrs []*int8 `wire:"value:${env.INT8S}"`

	Int16     int16    `wire:"value:${env.INT16}"`
	Int16Ptr  *int16   `wire:"value:${env.INT16}"`
	Int16s    []int16  `wire:"value:${env.INT16S}"`
	Int16Ptrs []*int16 `wire:"value:${env.INT16S}"`

	Int32     int32    `wire:"value:${env.INT32}"`
	Int32Ptr  *int32   `wire:"value:${env.INT32}"`
	Int32s    []int32  `wire:"value:${env.INT32S}"`
	Int32Ptrs []*int32 `wire:"value:${env.INT32S}"`

	Int64     int64    `wire:"value:${env.INT64}"`
	Int64Ptr  *int64   `wire:"value:${env.INT64}"`
	Int64s    []int64  `wire:"value:${env.INT64S}"`
	Int64Ptrs []*int64 `wire:"value:${env.INT64S}"`

	Uint     uint    `wire:"value:${env.UINT}"`
	UintPtr  *uint   `wire:"value:${env.UINT}"`
	Uints    []uint  `wire:"value:${env.UINTS}"`
	UintPtrs []*uint `wire:"value:${env.UINTS}"`

	Uint8     uint8    `wire:"value:${env.UINT8}"`
	Uint8Ptr  *uint8   `wire:"value:${env.UINT8}"`
	Uint8s    []uint8  `wire:"value:${env.UINT8S}"`
	Uint8Ptrs []*uint8 `wire:"value:${env.UINT8S}"`

	Uint16     uint16    `wire:"value:${env.UINT16}"`
	Uint16Ptr  *uint16   `wire:"value:${env.UINT16}"`
	Uint16s    []uint16  `wire:"value:${env.UINT16S}"`
	Uint16Ptrs []*uint16 `wire:"value:${env.UINT16S}"`

	Uint32     uint32    `wire:"value:${env.UINT32}"`
	Uint32Ptr  *uint32   `wire:"value:${env.UINT32}"`
	Uint32s    []uint32  `wire:"value:${env.UINT32S}"`
	Uint32Ptrs []*uint32 `wire:"value:${env.UINT32S}"`

	Uint64     uint64    `wire:"value:${env.UINT64}"`
	Uint64Ptr  *uint64   `wire:"value:${env.UINT64}"`
	Uint64s    []uint64  `wire:"value:${env.UINT64S}"`
	Uint64Ptrs []*uint64 `wire:"value:${env.UINT64S}"`

	Float32     float32    `wire:"value:${env.FLOAT32}"`
	Float32Ptr  *float32   `wire:"value:${env.FLOAT32}"`
	Float32s    []float32  `wire:"value:${env.FLOAT32S}"`
	Float32Ptrs []*float32 `wire:"value:${env.FLOAT32S}"`

	Float64     float64    `wire:"value:${env.FLOAT64}"`
	Float64Ptr  *float64   `wire:"value:${env.FLOAT64}"`
	Float64s    []float64  `wire:"value:${env.FLOAT64S}"`
	Float64Ptrs []*float64 `wire:"value:${env.FLOAT64S}"`

	Duration     time.Duration    `wire:"value:${env.DURATION}"`
	Durations    []time.Duration  `wire:"value:${env.DURATIONS}"`
	DurationPtr  *time.Duration   `wire:"value:${env.DURATION}"`
	DurationPtrs []*time.Duration `wire:"value:${env.DURATIONS}"`

	Unmarshaler     unmarshaler    `wire:"value:${env.UNMARSHALER}"`
	UnmarshalerPtr  *unmarshaler   `wire:"value:${env.UNMARSHALER}"`
	Unmarshalers    []unmarshaler  `wire:"value:${env.UNMARSHALERS}"`
	UnmarshalerPtrs []*unmarshaler `wire:"value:${env.UNMARSHALERS}"`

	URL     url.URL    `wire:"value:${env.URL}"`
	URLPtr  *url.URL   `wire:"value:${env.URL}"`
	URLs    []url.URL  `wire:"value:${env.URLS}"`
	URLPtrs []*url.URL `wire:"value:${env.URLS}"`

	StringWithdefault string `wire:"value:${env.DATABASE_URL}" envDefault:"postgres://localhost:5432/db"`

	CustomSeparator []string `wire:"value:${env.SEPSTRINGS}" envSeparator:":"`

	NonDefined struct {
		String string `wire:"value:${env.NONDEFINED_STR}"`
	}

	NestedNonDefined struct {
		NonDefined struct {
			String string `wire:"value:${env.STR}"`
		} `envPrefix:"NONDEFINED_"`
	} `envPrefix:"PRF_"`

	NotAnEnv   string
	unexported string `wire:"value:${env.FOO}"`

	PointStruct
	PS         PointStruct
	PPS        *PointStruct
	PPSNotInit *PointStruct

	SlicePointStruct   []PointStruct
	SlicePPointStruct  []*PointStruct
	PSlicePointStruct  *[]PointStruct
	PSlicePPointStruct *[]PointStruct
}

type WalkSliceStruct struct {
	PointStruct
	SlicePointStructInit   []PointStruct
	SlicePPointStructInit  []*PointStruct
	SlicePointStruct       []PointStruct
	SlicePPointStruct      []*PointStruct
	PSlicePointStructInit  *[]PointStruct
	PSlicePPointStructInit *[]*PointStruct
	PSlicePointStruct      *[]PointStruct
	PSlicePPointStruct     *[]*PointStruct
}

func TestWalkSliceStruct(t *testing.T) {

	init1 := make([]PointStruct, 2)
	init2 := make([]*PointStruct, 2)
	init2 = append(init2, &PointStruct{})
	wk := &WalkSliceStruct{
		SlicePointStructInit:   make([]PointStruct, 2),
		SlicePPointStructInit:  make([]*PointStruct, 2),
		PSlicePointStructInit:  &init1,
		PSlicePPointStructInit: &init2,
	}
	wk.SlicePPointStructInit = append(wk.SlicePPointStructInit, &PointStruct{})

	count := 0
	err := WalkField(wk, func(fieldValue reflect.Value, structField reflect.StructField, rootValues []reflect.Value) error {
		count++
		println(count, GetFieldPath(structField, rootValues), fieldValue.Kind().String())
		return nil
	})

	if err != nil {
		t.Error(err)
	}
}

func TestWalk(t *testing.T) {

	count := 0
	var hello = "abc"
	wk := &WireValueExprStruct{
		StringPtrInited: &hello,
		PPS:             &PointStruct{},
	}
	err := WalkField(wk, func(fieldValue reflect.Value, structField reflect.StructField, rootValues []reflect.Value) error {
		count++
		println(count, GetFieldPath(structField, rootValues), fieldValue.Kind().String())
		return nil
	})

	if err != nil {
		t.Error(err)
	}
}

func TestWalkWithTag(t *testing.T) {

	count := 0
	var hello = "abc"
	wk := &WireValueExprStruct{
		StringPtrInited: &hello,
		PPS:             &PointStruct{},
	}

	err := WalkWithTagNames(wk, []string{"env"}, func(fieldValue reflect.Value, structField reflect.StructField, rootValues []reflect.Value, tags map[string]string) error {
		count++
		println(count, GetFieldPath(structField, rootValues), fieldValue.Kind().String(), tags)
		return nil
	})

	if err != nil {
		t.Error(err)
	}
}
