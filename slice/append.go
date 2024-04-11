package slice

func PushFront[T any](s []T, item T) []T {
	s = append(s, item)
	copy(s[1:], s)
	s[0] = item
	return s
}
