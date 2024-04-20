package maths

import "golang.org/x/exp/constraints"

func ABS[T constraints.Float | constraints.Integer](x T) T {
	if x < 0 {
		return -x
	}
	return x
}
