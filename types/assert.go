package types

// Assertion tries to assert from interface{} to other type
func Assertion[T any](v any) (T, bool) {
	t, ok := v.(T)
	return t, ok
}

// Switch to new type
func Switch[T any](v any) (T, bool) {
	switch v := v.(type) {
	case T:
		return v, true
	default:
		var zero T
		return zero, false
	}
}
