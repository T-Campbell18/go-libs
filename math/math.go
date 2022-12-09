package math

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Float | constraints.Integer
}

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
