package maths

import "golang.org/x/exp/constraints"

func abs[T constraints.Float | constraints.Signed](x T) T {
	if x < 0 {
		return -x
	}
	return x
}
