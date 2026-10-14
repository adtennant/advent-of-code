package main

import (
	_ "embed"
	"fmt"

	"adtennant.dev/aoc/util"
)

type point = util.Point[int]

func findStart(grid *util.Grid) point {
	for x := 0; x < grid.Width(); x++ {
		if grid.Get(point{X: x, Y: 0}) == '.' {
			return point{X: x, Y: 0}
		}
	}

	return point{X: -1, Y: -1}
}

func findEnd(grid *util.Grid) point {
	for x := 0; x < grid.Width(); x++ {
		if grid.Get(point{X: x, Y: grid.Height() - 1}) == '.' {
			return point{X: x, Y: grid.Height() - 1}
		}
	}

	return point{X: -1, Y: -1}
}

var neighbours = map[byte][]util.Direction{
	'.': {util.UP, util.DOWN, util.LEFT, util.RIGHT},
	'^': {util.UP},
	'>': {util.RIGHT},
	'v': {util.DOWN},
	'<': {util.LEFT},
}

func findLongest(grid *util.Grid, start point, end point) int {
	q := util.NewQueue[point]()
	q.Push(start)

	// seen := util.NewSet[point]()
	cost := map[point]int{
		start: 0,
	}

	for q.Len() > 0 {
		current := q.Pop()

		if current == end {
			continue
		}

		/*if seen.Contains(current) {
			continue
		}

		seen.Insert(current)*/

		for _, dir := range neighbours[grid.Get(current)] {
			delta := util.GetDelta[int](dir)

			next := current.Translate(delta)
			nextCost := cost[current] + 1

			if _, ok := grid.TryGet(next); !ok {
				continue
			}

			if _, ok := cost[next]; !ok || nextCost > cost[next] {
				cost[next] = nextCost
				q.Push(next)
			}
		}
	}

	return cost[end]
}

func Part1(input string) (int, error) {
	grid := util.ParseGrid(input)

	start := findStart(grid)
	end := findEnd(grid)

	fmt.Println(start, end)

	return findLongest(grid, start, end), nil
}

func Part2(input string) (int, error) {
	return -1, nil
}

//go:embed input.txt
var input string

func main() {
	util.Run(Part1, Part2, input)
}
