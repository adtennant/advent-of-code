package util

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Float | constraints.Integer
}

type Point[T Number] struct {
	X, Y T
}

func (p Point[T]) Add(v Point[T]) Point[T] {
	return Point[T]{p.X + v.X, p.Y + v.Y}
}

type Direction byte

const (
	UP    Direction = 'U'
	DOWN  Direction = 'D'
	LEFT  Direction = 'L'
	RIGHT Direction = 'R'
)

func Delta[T constraints.Signed](dir Direction) Point[T] {
	switch dir {
	case UP:
		return Point[T]{0, -1}
	case DOWN:
		return Point[T]{0, 1}
	case LEFT:
		return Point[T]{-1, 0}
	case RIGHT:
		return Point[T]{1, 0}
	}

	return Point[T]{0, 0}
}
