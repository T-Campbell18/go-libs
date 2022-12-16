package math

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Float | constraints.Integer
}

const PI = 3.14159265358979323846264338327950288419716939937510582097494459
const E = 2.71828182845904523536028747135266249775724709369995957496696763

func Min[T Number](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Max[T Number](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func Abs[T Number](a T) T {
	if a < 0 {
		return -a
	}
	return a
}

func Ceil[T constraints.Float](a T) T {
	return T(int(a) + 1)
}

func Floor[T constraints.Float](a T) T {
	return T(int(a))
}

func Round[T constraints.Float](a T) T {
	return T(int(a + 0.5))
}

func Pow[T Number](x T, n int) T {
	var pow int64 = int64(n)
	if pow < 0 {
		pow = -pow
		x = 1 / x
	}

	var ans T = 1
	var curr T = x
	for i := pow; i > 0; i /= 2 {
		if i%2 == 1 {
			ans *= curr
		}
		curr *= curr
	}
	return ans
}
