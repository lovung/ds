package maths

import "golang.org/x/exp/constraints"

func Sum[T constraints.Ordered](a ...T) T {
	var sum T
	for i := range a {
		sum += a[i]
	}
	return sum
}
