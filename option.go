package structure

type Option struct {
	ZeroFields           bool
	WeaklyTypedInput     bool
	StringSplitSeparator string

	// If ErrorUnused is true, then it is an error for there to exist
	// keys in the original map that were unused in the decoding process
	// (extra keys).
	ErrorUnused bool

	// If ErrorUnset is true, then it is an error for there to exist
	// fields in the result that were not set in the decoding process
	// (extra fields). This only applies to decoding to a struct. This
	// will affect all nested structs as well.
	ErrorUnset bool

	// Squash will squash embedded structs.  A squash tag may also be
	// added to an individual struct field using a tag.  For example:
	//
	//  type Parent struct {
	//      Child `mapping:",squash"`
	//  }
	Squash bool

	// The tag name that mapping reads for field names. This
	// defaults to "mapping"
	TagName string

	// IgnoreUntaggedFields ignores all struct fields without explicit
	// TagName, comparable to `mapstructure:"-"` as default behaviour.
	IgnoreUntaggedFields bool

	// MatchName is the function used to match the map key to the struct
	// field name or tag. Defaults to `strings.EqualFold`. This can be used
	// to implement case-sensitive tag values, support snake casing, etc.
	MatchName func(mapKey, fieldName string) bool
}
