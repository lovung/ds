package maths

import "golang.org/x/exp/constraints"

func gcd[T constraints.Integer](a, b T) T {
	for b != 0 {
		b, a = a%b, b
	}
	return a
}

func GCD[T constraints.Integer](numbers ...T) T {
	if len(numbers) == 0 {
		return 0
	}
	result := numbers[0]
	for i := 1; i < len(numbers); i++ {
		result = gcd(result, numbers[i])
	}
	return result
}
