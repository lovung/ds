package slice

import "golang.org/x/exp/constraints"

func Reverse[T any](s []T) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// MaxValueAndIndex return max value and the index of first max
func MaxValueAndIndex[T constraints.Ordered](s []T) (max T, idx int) {
	if len(s) == 0 {
		return max, -1
	}
	max = s[0]
	idx = 0
	for i, v := range s {
		if v > max {
			max = v
			idx = i
		}
	}
	return max, idx
}
