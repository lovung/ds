package slice

func RotateLeft[T any](items []T, k int) {
	temp := make([]T, k)
	copy(temp, items)
	copy(items, items[k:])
	copy(items[len(items)-k:], temp)
}

func RotateRight[T any](items []T, k int) {
	RotateLeft(items, len(items)-k)
}

