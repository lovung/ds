package slice

import (
	"sort"

	"golang.org/x/exp/constraints"
)

func SortAscUniq[T constraints.Ordered](s []T) []T {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	return UniqSorted(s)
}

func SortDescUniq[T constraints.Ordered](s []T) []T {
	sort.Slice(s, func(i, j int) bool {
		return s[i] > s[j]
	})
	return UniqSorted(s)
}

func UniqSorted[T constraints.Ordered](s []T) []T {
	l := 0
	for r := 0; r < len(s); {
		for r < len(s) && s[l] == s[r] {
			r++
		}
		l++
		if r < len(s) {
			s[l] = s[r]
		}
	}
	return s[:l]
}
