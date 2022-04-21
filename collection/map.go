package collection

type Map[K comparable, V any] map[K]V

func (m Map[K, V]) Clear() {
	for k := range m {
		delete(m, k)
	}
}

func (m Map[K, V]) Has(key K) bool {
	_, ok := m[key]
	return ok
}

func (m Map[K, V]) Len() int {
	return len(m)
}

func (m Map[K, V]) IsEmpty() bool {
	return len(m) == 0
}

func (m Map[K, V]) Keys() Vec[K] {
	v := make(Vec[K], 0, len(m))
	for k := range m {
		v = append(v, k)
	}
	return v
}

func (m Map[K, V]) Values() Vec[V] {
	v := make(Vec[V], 0, len(m))
	for _, val := range m {
		v = append(v, val)
	}
	return v
}
