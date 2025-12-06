package util

import "cmp"

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func Min[T cmp.Ordered](a, b T) T {
	if a < b {
		return a
	} else {
		return b
	}
}

func Max[V cmp.Ordered](a, b V) V {
	if a > b {
		return a
	} else {
		return b
	}
}
