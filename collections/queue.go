package collections

import "fmt"

type Queue[T any] struct {
	items []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{items: make([]T, 0)}
}

func (q *Queue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
}

func (q *Queue[T]) Add(item T) {
	q.Enqueue(item)
}

func (q *Queue[T]) Dequeue() T {
	if q.IsEmpty() {
		var empty T
		return empty
	}

	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func (q *Queue[T]) Peek() T {
	if q.IsEmpty() {
		var empty T
		return empty
	}

	return q.items[0]
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue[T]) Size() int {
	return len(q.items)
}

func (q *Queue[T]) String() string {
	return fmt.Sprintf("%v", q.items)
}
