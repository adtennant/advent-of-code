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

type Grid struct {
	cells  map[Point[int]]byte
	width  int
	height int
}

func (g *Grid) Find(c byte) []Point[int] {
	var cells []Point[int]

	for p, cell := range g.cells {
		if cell == c {
			cells = append(cells, p)
		}
	}

	return cells
}

func (g *Grid) Get(p Point[int]) byte {
	return g.cells[p]
}

func (g *Grid) TryGet(p Point[int]) (byte, bool) {
	c, ok := g.cells[p]
	return c, ok
}

func (g *Grid) Width() int {
	return g.width
}

func (g *Grid) Height() int {
	return g.height
}

func ParseGrid(input string) *Grid {
	lines := Lines(input)
	cells := make(map[Point[int]]byte)

	for y, line := range lines {
		for x, c := range []byte(line) {
			cells[Point[int]{X: x, Y: y}] = c
		}
	}

	width := len(lines[0])
	height := len(lines)

	return &Grid{cells, width, height}
}
