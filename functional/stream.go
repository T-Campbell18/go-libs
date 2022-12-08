package functional

type Stream[T any] struct {
	items []T
}

func NewStream[T any](items []T) *Stream[T] {
	return &Stream[T]{items: items}
}

func (s *Stream[T]) Map(f func(T) T) *Stream[T] {
	result := make([]T, len(s.items))
	for i, item := range s.items {
		result[i] = f(item)
	}
	return NewStream(result)
}

func (s *Stream[T]) Filter(f func(T) bool) *Stream[T] {
	result := make([]T, 0)
	for _, item := range s.items {
		if f(item) {
			result = append(result, item)
		}
	}
	return NewStream(result)
}

func (s *Stream[T]) Reduce(f func(T, T) T, initial T) T {
	result := initial
	for _, item := range s.items {
		result = f(result, item)
	}
	return result
}

func (s *Stream[T]) ForEach(f func(T)) {
	for _, item := range s.items {
		f(item)
	}
}

func (s *Stream[T]) ToSlice() []T {
	return s.items
}

func (s *Stream[T]) MapToInt(f func(T) int) *Stream[int] {
	result := make([]int, len(s.items))
	for i, item := range s.items {
		result[i] = f(item)
	}
	return NewStream(result)
}

func (s *Stream[T]) MapToIntPtr(f func(T) *int) *Stream[*int] {
	result := make([]*int, len(s.items))
	for i, item := range s.items {
		result[i] = f(item)
	}
	return NewStream(result)
}

func (s *Stream[T]) MapToInt64(f func(T) int64) *Stream[int64] {
	result := make([]int64, len(s.items))
	for i, item := range s.items {
		result[i] = f(item)
	}
	return NewStream(result)
}

func (s *Stream[T]) MapToIntPtr64(f func(T) *int64) *Stream[*int64] {
	result := make([]*int64, len(s.items))
	for i, item := range s.items {
		result[i] = f(item)
	}
	return NewStream(result)
}

func (s *Stream[T]) MapToString(f func(T) string) *Stream[string] {
	result := make([]string, len(s.items))
	for i, item := range s.items {
		result[i] = f(item)
	}
	return NewStream(result)
}

func (s *Stream[T]) MapToStringPtr(f func(T) *string) *Stream[*string] {
	result := make([]*string, len(s.items))
	for i, item := range s.items {
		result[i] = f(item)
	}
	return NewStream(result)
}

func (s *Stream[T]) MapToFloat64(f func(T) float64) *Stream[float64] {
	result := make([]float64, len(s.items))
	for i, item := range s.items {
		result[i] = f(item)
	}
	return NewStream(result)
}

func (s *Stream[T]) MapToFloat64Ptr(f func(T) *float64) *Stream[*float64] {
	result := make([]*float64, len(s.items))
	for i, item := range s.items {
		result[i] = f(item)
	}
	return NewStream(result)
}

func (s *Stream[T]) MapToBool(f func(T) bool) *Stream[bool] {
	result := make([]bool, len(s.items))
	for i, item := range s.items {
		result[i] = f(item)
	}
	return NewStream(result)
}

func (s *Stream[T]) MapToBoolPtr(f func(T) *bool) *Stream[*bool] {
	result := make([]*bool, len(s.items))
	for i, item := range s.items {
		result[i] = f(item)
	}
	return NewStream(result)
}

func (s *Stream[T]) MapToObject(f func(T) interface{}) *Stream[interface{}] {
	result := make([]interface{}, len(s.items))
	for i, item := range s.items {
		result[i] = f(item)
	}
	return NewStream(result)
}
