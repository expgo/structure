package structure

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func init() {
	RegisterMapper[string, bool](string2boolMapper)
	RegisterMapper[string, int](string2intMapper)
	RegisterMapper[string, uint](string2uintMapper)
	RegisterMapper[string, float64](string2floatMapper)
	RegisterMapper[string, string](string2stringMapper)
}

func string2boolMapper(from reflect.Value, to reflect.Value) error {
	str := strings.TrimSpace(from.String())
	b, err := strconv.ParseBool(str)
	if err == nil {
		to.SetBool(b)
	} else if strings.TrimSpace(str) == "" {
		to.SetBool(false)
	} else {
		return fmt.Errorf("cannot parse '%s' as bool: %s", str, err)
	}
	return nil
}

func string2intMapper(from reflect.Value, to reflect.Value) error {
	str := strings.TrimSpace(from.String())
	if str == "" {
		str = "0"
	}

	i, err := strconv.ParseInt(str, 0, to.Type().Bits())
	if err == nil {
		to.SetInt(i)
	} else {
		return fmt.Errorf("cannot parse '%s' as int: %s", str, err)
	}

	return nil
}

func string2uintMapper(from reflect.Value, to reflect.Value) error {
	str := strings.TrimSpace(from.String())
	if str == "" {
		str = "0"
	}

	ui, err := strconv.ParseUint(str, 0, to.Type().Bits())
	if err == nil {
		to.SetUint(ui)
	} else {
		return fmt.Errorf("cannot parse '%s' as uint: %s", str, err)
	}

	return nil
}

func string2floatMapper(from reflect.Value, to reflect.Value) error {
	str := strings.TrimSpace(from.String())
	if str == "" {
		str = "0"
	}

	f, err := strconv.ParseFloat(str, to.Type().Bits())
	if err == nil {
		to.SetFloat(f)
	} else {
		return fmt.Errorf("cannot parse '%s' as float: %s", str, err)
	}

	return nil
}

func string2stringMapper(from reflect.Value, to reflect.Value) error {
	to.SetString(from.String())
	return nil
}
