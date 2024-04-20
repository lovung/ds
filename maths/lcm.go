package maths

import "golang.org/x/exp/constraints"

// find Least Common Multiple (LCM) via GCD
func LCM[T constraints.Integer](integers ...T) T {
	if len(integers) == 0 {
		return 0
	}
	if len(integers) == 1 {
		return integers[0]
	}
	a := integers[0]
	b := integers[1]
	result := a * b / gcd(a, b)

	for i := 2; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
