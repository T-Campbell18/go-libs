package functional

type MappableStream[T, R any] struct {
	items []T
}

func NewMappableStream[T, R any](items []T) *MappableStream[T, R] {
	return &MappableStream[T, R]{items: items}
}

func ToMappableStream[T, U any](stream *Stream[T]) *MappableStream[T, U] {
	return NewMappableStream[T, U](stream.ToSlice())
}

func (s *MappableStream[T, R]) Map(f func(T) R) *Stream[R] {
	result := make([]R, len(s.items))
	for i, item := range s.items {
		result[i] = f(item)
	}
	return NewStream(result)
}
