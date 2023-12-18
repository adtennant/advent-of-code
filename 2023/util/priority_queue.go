package util

import "container/heap"

type item struct {
	value    any
	priority int
}

type priorityqueue []*item

func (pq priorityqueue) Len() int { return len(pq) }

func (pq priorityqueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq priorityqueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *priorityqueue) Push(x any) {
	item := x.(*item)
	*pq = append(*pq, item)
}

func (pq *priorityqueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}

type PriorityQueue[T any] struct {
	inner priorityqueue
}

func NewPriorityQueue[T any]() *PriorityQueue[T] {
	inner := make(priorityqueue, 0)
	heap.Init(&inner)

	return &PriorityQueue[T]{inner}
}

func (pq PriorityQueue[T]) Len() int { return pq.inner.Len() }

func (pq *PriorityQueue[T]) Push(value T, priority int) {
	heap.Push(&pq.inner, &item{
		value,
		priority,
	})
}

func (pq *PriorityQueue[T]) Pop() (T, int) {
	item := heap.Pop(&pq.inner).(*item)
	return item.value.(T), item.priority
}
