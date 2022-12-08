package functional

type BiStream[T, R any] struct {
	items []T
}

func NewBiStream[T, R any](items []T) *BiStream[T, R] {
	return &BiStream[T, R]{items: items}
}

func ToBiStream[T, U any](stream *Stream[T]) *BiStream[T, U] {
	return NewBiStream[T, U](stream.ToSlice())
}

func (s *BiStream[T, R]) Map(f func(T) R) *Stream[R] {
	result := make([]R, len(s.items))
	for i, item := range s.items {
		result[i] = f(item)
	}
	return NewStream(result)
}

func (s *BiStream[T, R]) Reduce(f func(R, T) R, initial R) R {
	result := initial
	for _, item := range s.items {
		result = f(result, item)
	}
	return result
}
