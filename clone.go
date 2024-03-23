package structure

// Clone is a generic function that creates a new instance of the given type T
// and assigns the value of the input parameter to the newly created instance.
// It then returns the pointer to the newly cloned instance.
//
// Parameters:
// - value: a pointer to the value to be cloned
//
// Returns:
// - a pointer to a newly created instance of type T, with the same value as the input parameter.
func Clone[T any](value *T) *T {
	if value == nil {
		return nil
	}

	t := new(T)
	*t = *value
	return t
}
