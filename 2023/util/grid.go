package util

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Float | constraints.Integer
}

type Point[T Number] struct {
	X, Y T
}

func (p Point[T]) Translate(v Vector[T]) Point[T] {
	return Point[T]{p.X + v.X, p.Y + v.Y}
}

type Vector[T Number] struct {
	X, Y T
}

func (v Vector[T]) Scale(scale T) Vector[T] {
	return Vector[T]{v.X * scale, v.Y * scale}
}

type Direction byte

const (
	UP    Direction = 'U'
	DOWN  Direction = 'D'
	LEFT  Direction = 'L'
	RIGHT Direction = 'R'
)

func GetDelta[T constraints.Signed](dir Direction) Vector[T] {
	switch dir {
	case UP:
		return Vector[T]{0, -1}
	case DOWN:
		return Vector[T]{0, 1}
	case LEFT:
		return Vector[T]{-1, 0}
	case RIGHT:
		return Vector[T]{1, 0}
	}

	return Vector[T]{0, 0}
}
