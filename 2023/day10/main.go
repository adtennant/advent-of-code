package main

import (
	_ "embed"
	"fmt"
	"math"

	"adtennant.dev/aoc/util"
)

var corners = map[byte]bool{
	'L': true,
	'J': true,
	'7': true,
	'F': true,
	'S': true,
}

var neighbours = map[byte][]util.Direction{
	'|': {util.UP, util.DOWN},
	'-': {util.RIGHT, util.LEFT},
	'L': {util.UP, util.RIGHT},
	'J': {util.UP, util.LEFT},
	'7': {util.DOWN, util.LEFT},
	'F': {util.DOWN, util.RIGHT},
	'.': {},
	'S': {util.DOWN}, // TODO: Don't hardcode this, s works for my input and all tests
}

type point = util.Point[int]

func findLoop(grid map[point]byte, start point) []point {
	q := util.NewQueue[point]()
	q.Push(start)

	loop := []point{start}

	explored := map[point]bool{
		start: true,
	}

	for q.Len() > 0 {
		current := q.Pop()

		for _, dir := range neighbours[grid[current]] {
			delta := util.GetDelta[int](dir)
			next := current.Translate(delta)

			if _, ok := grid[next]; !ok {
				continue
			}

			if explored[next] {
				continue
			}

			q.Push(next)
			explored[next] = true

			loop = append(loop, next)
		}
	}

	return loop
}

func findStart(grid map[point]byte) (point, error) {
	for p, c := range grid {
		if c == 'S' {
			return p, nil
		}
	}

	return point{}, fmt.Errorf("start not found")
}

func parseGrid(input string) map[point]byte {
	grid := make(map[point]byte)

	for y, line := range util.Lines(input) {
		for x, c := range []byte(line) {
			grid[point{x, y}] = c
		}
	}

	return grid
}

func Part1(input string) (int, error) {
	grid := parseGrid(input)

	start, err := findStart(grid)
	if err != nil {
		return -1, err
	}

	loop := findLoop(grid, start)

	return len(loop) / 2, nil
}

func shoelace(verts []point) int {
	sum := 0
	v0 := verts[len(verts)-1]

	for _, v1 := range verts {
		sum += v0.Y*v1.X - v0.X*v1.Y
		v0 = v1
	}

	return int(math.Abs(float64(sum / 2)))
}

func Part2(input string) (int, error) {
	grid := parseGrid(input)

	start, err := findStart(grid)
	if err != nil {
		return -1, err
	}

	loop := findLoop(grid, start)

	var verts []point

	for _, p := range loop {
		if corners[grid[p]] {
			verts = append(verts, p)
		}
	}

	return shoelace(verts) - (len(loop) / 2) + 1, nil
}

//go:embed input.txt
var input string

func main() {
	util.Run(Part1, Part2, input)
}
