package collections

import "fmt"

type Map[T comparable, V any] struct {
	items map[T]V
}

func NewMap[T comparable, V any]() *Map[T, V] {
	return &Map[T, V]{items: make(map[T]V)}
}

func (m *Map[T, V]) Put(key T, value V) {
	m.items[key] = value
}

func (m *Map[T, V]) Get(key T) V {
	return m.items[key]
}

func (m *Map[T, V]) Remove(key T) {
	delete(m.items, key)
}

func (m *Map[T, V]) Contains(key T) bool {
	_, ok := m.items[key]
	return ok
}

func (m *Map[T, V]) IsEmpty() bool {
	return len(m.items) == 0
}

func (m *Map[T, V]) Size() int {
	return len(m.items)
}

func (m *Map[T, V]) Keys() []T {
	keys := make([]T, len(m.items))
	i := 0
	for k := range m.items {
		keys[i] = k
		i++
	}
	return keys
}

func (m *Map[T, V]) Values() []V {
	values := make([]V, len(m.items))
	i := 0
	for _, v := range m.items {
		values[i] = v
		i++
	}
	return values
}

func (m *Map[T, V]) GetOrDefault(key T, defaultValue V) V {
	if m.Contains(key) {
		return m.Get(key)
	}
	return defaultValue
}

func (m *Map[T, V]) String() string {
	return fmt.Sprintf("%v", m.items)
}
