package functional

func Map[T, R any](items []T, f func(T) R) []R {
	result := make([]R, len(items))
	for i, item := range items {
		result[i] = f(item)
	}
	return result
}

func Filter[T any](items []T, f func(T) bool) []T {
	result := make([]T, 0)
	for _, item := range items {
		if f(item) {
			result = append(result, item)
		}
	}
	return result
}

func Reduce[T, R any](items []T, f func(R, T) R, initial R) R {
	result := initial
	for _, item := range items {
		result = f(result, item)
	}
	return result
}

func ForEach[T any](items []T, f func(T)) {
	for _, item := range items {
		f(item)
	}
}
