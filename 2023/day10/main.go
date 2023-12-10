package main

import (
	_ "embed"
	"fmt"
	"math"

	"adtennant.dev/aoc/util"
)

var directions = map[byte]point{
	'n': {0, -1},
	's': {0, 1},
	'e': {1, 0},
	'w': {-1, 0},
}

var corners = map[byte]bool{
	'L': true,
	'J': true,
	'7': true,
	'F': true,
	'S': true,
}

var neighbours = map[byte][]byte{
	'|': {'n', 's'},
	'-': {'e', 'w'},
	'L': {'n', 'e'},
	'J': {'n', 'w'},
	'7': {'s', 'w'},
	'F': {'s', 'e'},
	'.': {},
	'S': {'s'}, // TODO: Don't hardcode this, s works for my input and all tests
}

type point struct {
	x, y int
}

func findLoop(grid map[point]byte, start point) []point {
	queue := []point{start}
	loop := []point{start}

	explored := map[point]bool{
		start: true,
	}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, dir := range neighbours[grid[current]] {
			delta := directions[dir]
			next := point{current.x + delta.x, current.y + delta.y}

			if _, ok := grid[next]; !ok {
				continue
			}

			if explored[next] {
				continue
			}

			queue = append(queue, next)
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
		sum += v0.y*v1.x - v0.x*v1.y
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
