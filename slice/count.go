package slice

import "golang.org/x/exp/constraints"

func Count[T constraints.Ordered](s []T, target T) int {
	cnt := 0
	for i := range s {
		if target == s[i] {
			cnt++
		}
	}
	return cnt
}

func CountIf[T any](s []T, checkFn func(T) bool) int {
	cnt := 0
	for i := range s {
		if checkFn(s[i]) {
			cnt++
		}
	}
	return cnt
}
