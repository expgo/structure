package structure

import (
	"encoding"
	"errors"
	"fmt"
	"reflect"
	"time"
)

func init() {
	RegisterMapper[string, encoding.TextUnmarshaler](string2TextUnmarshalerMapper)
	RegisterMapper[string, encoding.BinaryUnmarshaler](string2BinaryUnmarshalerMapper)
	RegisterMapper[string, time.Duration](string2durationMapper)
}

func string2TextUnmarshalerMapper(from reflect.Value, to reflect.Value) error {
	textUnmarshaler, ok := to.Interface().(encoding.TextUnmarshaler)
	if !ok {
		if to.CanAddr() {
			textUnmarshaler, ok = to.Addr().Interface().(encoding.TextUnmarshaler)
		}
	}

	if ok {
		err := textUnmarshaler.UnmarshalText([]byte(from.String()))
		if err != nil {
			return errors.New(fmt.Sprintf("unable to unmarshal text: %v", err))
		}
		return nil
	}

	return errors.New(fmt.Sprintf("type %s could not convert to TextUnmarshaler", to.Type()))
}

func string2BinaryUnmarshalerMapper(from reflect.Value, to reflect.Value) error {
	binaryUnmarshaler, ok := to.Interface().(encoding.BinaryUnmarshaler)
	if !ok {
		if to.CanAddr() {
			binaryUnmarshaler, ok = to.Addr().Interface().(encoding.BinaryUnmarshaler)
		}
	}

	if ok {
		err := binaryUnmarshaler.UnmarshalBinary([]byte(from.String()))
		if err != nil {
			return errors.New(fmt.Sprintf("unable to unmarshal binary: %v", err))
		}

		return nil
	}

	return errors.New(fmt.Sprintf("type %s could not convert to BinaryUnmarshaler", to.Type()))
}

func string2durationMapper(from reflect.Value, to reflect.Value) error {
	s, err := time.ParseDuration(from.String())
	if err != nil {
		return errors.New(fmt.Sprintf("unable to parse duration: %v", err))
	}
	to.Set(reflect.ValueOf(s))
	return nil
}
