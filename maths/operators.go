package maths

import "golang.org/x/exp/constraints"

func Sum[T constraints.Ordered](a ...T) T {
	var sum T
	for i := range a {
		sum += a[i]
	}
	return sum
}

func GCD[T constraints.Integer](a, b T) T {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}
	if a > b {
		return GCD(a%b, b)
	}
	return GCD(a, b%a)
}
