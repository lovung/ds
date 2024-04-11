package slice

func RemoveHead[T any](items []T, k int) []T {
	if k > len(items) {
		return make([]T, 0)
	}
	return items[k:]
}

func RemoveTail[T any](items []T, k int) []T {
	if k > len(items) {
		return make([]T, 0)
	}
	return items[:len(items)-k]
}

func RemoveAt[T any](a []T, i int) []T {
	if i >= 0 && i < len(a)-1 {
		copy(a[i:], a[i+1:])
	}
	return a[:len(a)-1]
}
