package util

type Set[T comparable] map[T]bool

func NewSet[T comparable]() *Set[T] {
	s := Set[T](make(map[T]bool))
	return &s
}

func (s *Set[T]) Insert(value ...T) {
	for _, v := range value {
		(*s)[v] = true
	}
}

func (s *Set[T]) Contains(value T) bool {
	return (*s)[value]
}

func (s Set[T]) Len() int {
	return len(s)
}
