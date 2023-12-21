package main

import (
	_ "embed"

	"adtennant.dev/aoc/util"
)

type point = util.Point[int]

type state struct {
	pos   point
	steps int
}

func mod(a, b int) int {
	m := a % b
	if a < 0 && b < 0 {
		m -= b
	}
	if a < 0 && b > 0 {
		m += b
	}
	return m
}

func findPlots(grid *util.Grid, start point, steps int) *util.Set[point] {
	q := util.NewQueue[state]()
	q.Push(state{start, 0})

	seen := util.NewSet[state]()

	plots := util.NewSet[point]()

	for q.Len() > 0 {
		current := q.Pop()

		if current.steps == steps {
			plots.Insert(current.pos)
			continue
		}

		if seen.Contains(current) {
			continue
		}

		seen.Insert(current)

		for _, dir := range []util.Direction{util.UP, util.DOWN, util.LEFT, util.RIGHT} {
			delta := util.GetDelta[int](dir)

			next := state{current.pos.Translate(delta), current.steps + 1}

			if v := grid.Get(point{X: mod(next.pos.X, grid.Width()), Y: mod(next.pos.Y, grid.Height())}); v == '#' {
				continue
			}

			q.Push(next)
		}
	}

	return plots
}

func Part1(steps int) func(input string) (int, error) {
	return func(input string) (int, error) {
		grid := util.ParseGrid(input)
		start := grid.Find('S')[0]

		plots := findPlots(grid, start, steps)

		return plots.Len(), nil
	}
}

// TODO: Attempt to make this work for the example, is it even possible to make it work generically?
func Part2(steps int) func(input string) (int, error) {
	return func(input string) (int, error) {
		grid := util.ParseGrid(input)
		start := grid.Find('S')[0]

		var poly []int

		for i := 0; i < 3; i++ {
			poly = append(poly, findPlots(grid, start, (steps%grid.Width())+i*grid.Width()).Len())
		}

		x := steps % grid.Width()

		b0 := poly[0]
		b1 := poly[1] - poly[0]
		b2 := poly[2] - poly[1]

		return b0 + (b1 * x) + (x*(x-1)/2)*(b2-b1), nil
	}
}

//go:embed input.txt
var input string

func main() {
	util.Run[int](Part1(64), Part2(26501365), input)
}
