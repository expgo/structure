package structure

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

func init() {
	RegisterKindMapper(reflect.Map, reflect.Map, map2mapMapper)
	RegisterKindMapper(reflect.Map, reflect.Struct, map2structMapper)
}

func map2mapMapper(from reflect.Value, to reflect.Value, _ *Option) error {
	return nil
}

func map2structMapper(from reflect.Value, to reflect.Value, option *Option) error {
	dataValType := from.Type()
	if kind := dataValType.Key().Kind(); kind != reflect.String && kind != reflect.Interface {
		return fmt.Errorf(
			"needs a map with string keys, has '%s' keys", dataValType.Key().Kind())
	}

	dataValKeys := make(map[reflect.Value]struct{})
	dataValKeysUnused := make(map[interface{}]struct{})
	for _, dataValKey := range from.MapKeys() {
		dataValKeys[dataValKey] = struct{}{}
		dataValKeysUnused[dataValKey.Interface()] = struct{}{}
	}

	targetValKeysUnused := make(map[interface{}]struct{})
	errors := make([]error, 0)

	// This slice will keep track of all the structs we'll be decoding.
	// There can be more than one struct if there are embedded structs
	// that are squashed.
	structs := make([]reflect.Value, 1, 5)
	structs[0] = to

	// Compile the list of all the fields that we're going to be decoding
	// from all the structs.
	type field struct {
		field reflect.StructField
		val   reflect.Value
	}

	// remainField is set to a valid field set with the "remain" tag if
	// we are keeping track of remaining values.
	var remainField *field

	fields := []field{}
	for len(structs) > 0 {
		structVal := structs[0]
		structs = structs[1:]

		structType := structVal.Type()

		for i := 0; i < structType.NumField(); i++ {
			fieldType := structType.Field(i)
			fieldVal := structVal.Field(i)
			if fieldVal.Kind() == reflect.Ptr && fieldVal.Elem().Kind() == reflect.Struct {
				// Handle embedded struct pointers as embedded structs.
				fieldVal = fieldVal.Elem()
			}

			// If "squash" is specified in the tag, we squash the field down.
			squash := option.Squash && fieldVal.Kind() == reflect.Struct && fieldType.Anonymous
			remain := false

			// We always parse the tags cause we're looking for other tags too
			tagParts := strings.Split(fieldType.Tag.Get(option.TagName), ",")
			for _, tag := range tagParts[1:] {
				if tag == "squash" {
					squash = true
					break
				}

				if tag == "remain" {
					remain = true
					break
				}
			}

			if squash {
				if fieldVal.Kind() != reflect.Struct {
					errors = appendErrors(errors,
						fmt.Errorf("%s: unsupported type for squash: %s", fieldType.Name, fieldVal.Kind()))
				} else {
					structs = append(structs, fieldVal)
				}
				continue
			}

			// Build our field
			if remain {
				remainField = &field{fieldType, fieldVal}
			} else {
				// Normal struct field, store it away
				fields = append(fields, field{fieldType, fieldVal})
			}
		}
	}

	// for fieldType, field := range fields {
	for _, f := range fields {
		field, fieldValue := f.field, f.val
		fieldName := field.Name

		tagValue := field.Tag.Get(option.TagName)
		tagValue = strings.SplitN(tagValue, ",", 2)[0]
		if tagValue != "" {
			fieldName = tagValue
		}

		rawMapKey := reflect.ValueOf(fieldName)
		rawMapVal := from.MapIndex(rawMapKey)
		if !rawMapVal.IsValid() {
			// Do a slower search by iterating over each key and
			// doing case-insensitive search.
			for dataValKey := range dataValKeys {
				mK, ok := dataValKey.Interface().(string)
				if !ok {
					// Not a string key
					continue
				}

				if option.MatchName(mK, fieldName) {
					rawMapKey = dataValKey
					rawMapVal = from.MapIndex(dataValKey)
					break
				}
			}

			if !rawMapVal.IsValid() {
				// There was no matching key in the map for the value in
				// the struct. Remember it for potential errors and metadata.
				targetValKeysUnused[fieldName] = struct{}{}
				continue
			}
		}

		if !fieldValue.IsValid() {
			// This should never happen
			panic("field is not valid")
		}

		// If we can't set the field, then it is unexported or something,
		// and we just continue onwards.
		if !fieldValue.CanSet() {
			continue
		}

		// Delete the key we're using from the unused map so we stop tracking
		delete(dataValKeysUnused, rawMapKey.Interface())

		//// If the name is empty string, then we're at the root, and we
		//// don't dot-join the fields.
		//if name != "" {
		//	fieldName = name + "." + fieldName
		//}

		if err := MapToValueWithOption(rawMapVal.Interface(), fieldValue, option); err != nil {
			errors = appendErrors(errors, err)
		}
	}

	// If we have a "remain"-tagged field and we have unused keys then
	// we put the unused keys directly into the remain field.
	if remainField != nil && len(dataValKeysUnused) > 0 {
		// Build a map of only the unused values
		remain := map[interface{}]interface{}{}
		for key := range dataValKeysUnused {
			remain[key] = from.MapIndex(reflect.ValueOf(key)).Interface()
		}

		// Decode it as-if we were just decoding this map onto our map.
		if err := MapToValueWithOption(remain, remainField.val, option); err != nil {
			errors = appendErrors(errors, err)
		}

		// Set the map to nil so we have none so that the next check will
		// not error (ErrorUnused)
		dataValKeysUnused = nil
	}

	if option.ErrorUnused && len(dataValKeysUnused) > 0 {
		keys := make([]string, 0, len(dataValKeysUnused))
		for rawKey := range dataValKeysUnused {
			keys = append(keys, rawKey.(string))
		}
		sort.Strings(keys)

		err := fmt.Errorf("invalid keys: %s", strings.Join(keys, ", "))
		errors = appendErrors(errors, err)
	}

	if option.ErrorUnset && len(targetValKeysUnused) > 0 {
		keys := make([]string, 0, len(targetValKeysUnused))
		for rawKey := range targetValKeysUnused {
			keys = append(keys, rawKey.(string))
		}
		sort.Strings(keys)

		err := fmt.Errorf("unset fields: %s", strings.Join(keys, ", "))
		errors = appendErrors(errors, err)
	}

	if len(errors) > 0 {
		return &Error{errors}
	}

	return nil
}
