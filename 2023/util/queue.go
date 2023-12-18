package util

type Queue[T any] []T

func NewQueue[T any]() *Queue[T] {
	q := Queue[T](make([]T, 0))
	return &q
}

func (q *Queue[T]) Push(value ...T) {
	*q = append(*q, value...)
}

func (q *Queue[T]) Pop() T {
	old := *q
	popped := old[0]
	*q = old[1:]
	return popped
}

func (q Queue[T]) Len() int {
	return len(q)
}
