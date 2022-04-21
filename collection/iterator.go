package collection

type Iterator[T any] func() (T, bool)

func (it Iterator[T]) Any(test func(item T) bool) bool {
	for i, ok := it(); ok; i, ok = it() {
		if test(i) {
			return true
		}
	}
	return false
}
