package main

import (
	_ "embed"

	"adtennant.dev/aoc/util"
)

func parseGrid(input string) (map[point]int, int, int) {
	grid := make(map[point]int)
	lines := util.Lines(input)

	for y, line := range lines {
		for x, c := range []byte(line) {
			grid[point{X: x, Y: y}] = int(c - '0')
		}
	}

	width := len(lines[0])
	height := len(lines)

	return grid, width, height
}

type point = util.Point[int]

type state struct {
	pos   point
	dir   util.Direction
	steps int
}

var neighbours = map[util.Direction][]util.Direction{
	util.UP:    {util.UP, util.LEFT, util.RIGHT},
	util.DOWN:  {util.DOWN, util.LEFT, util.RIGHT},
	util.LEFT:  {util.LEFT, util.UP, util.DOWN},
	util.RIGHT: {util.RIGHT, util.UP, util.DOWN},
}

func f(grid map[point]int, start, end point, minSteps, maxSteps int) int {
	q := util.NewPriorityQueue[state]()
	q.Push(state{start, util.RIGHT, 1}, 0)
	q.Push(state{start, util.DOWN, 1}, 0)

	seen := util.NewSet[state]()

	for q.Len() > 0 {
		current, cost := q.Pop()

		if current.pos == end && current.steps >= minSteps-1 {
			return cost
		}

		var candidates []util.Direction

		if current.steps < minSteps-1 {
			candidates = append(candidates, current.dir)
		} else {
			candidates = neighbours[current.dir]
		}

		for _, dir := range candidates {
			steps := 0

			if dir == current.dir {
				steps = current.steps + 1
			}

			if steps >= maxSteps {
				continue
			}

			delta := util.GetDelta[int](dir)
			next := current.pos.Translate(delta) //point{current.pos.x + delta.x, current.pos.y + delta.y}

			if _, ok := grid[next]; !ok {
				continue
			}

			key := state{current.pos, dir, steps}

			if seen.Contains(key) {
				continue
			}

			seen.Insert(key)

			q.Push(state{next, dir, steps}, cost+grid[next])
		}
	}

	return 0
}

func Part1(input string) (int, error) {
	grid, width, height := parseGrid(input)
	return f(grid, point{X: 0, Y: 0}, point{X: width - 1, Y: height - 1}, 0, 3), nil
}

func Part2(input string) (int, error) {
	grid, width, height := parseGrid(input)
	return f(grid, point{X: 0, Y: 0}, point{X: width - 1, Y: height - 1}, 4, 10), nil
}

//go:embed input.txt
var input string

func main() {
	util.Run(Part1, Part2, input)
}
