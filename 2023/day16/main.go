package main

import (
	_ "embed"

	"adtennant.dev/aoc/util"
)

type point struct {
	x, y int
}

func parseGrid(input string) (map[point]byte, int, int) {
	lines := util.Lines(input)
	grid := make(map[point]byte)

	for y, line := range lines {
		for x, c := range []byte(line) {
			grid[point{x, y}] = c
		}
	}

	width := len(lines[0])
	height := len(lines)

	return grid, width, height
}

type direction byte

const (
	UP    direction = 'u'
	DOWN  direction = 'd'
	LEFT  direction = 'l'
	RIGHT direction = 'r'
)

var directions = map[direction]point{
	'u': {0, -1},
	'd': {0, 1},
	'l': {-1, 0},
	'r': {1, 0},
}

type beam struct {
	p   point
	dir direction
}

func next(current point, dir direction) beam {
	delta := directions[dir]

	return beam{
		point{current.x + delta.x, current.y + delta.y},
		dir,
	}
}

func castBeam(grid map[point]byte, b beam) int {
	queue := []beam{b}
	seen := make(map[beam]bool)

	energised := make(map[point]bool)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if _, ok := grid[current.p]; !ok {
			continue
		}

		if seen[current] {
			continue
		}

		seen[current] = true
		energised[current.p] = true

		switch grid[current.p] {
		case '.':
			queue = append(queue, next(current.p, current.dir))
		case '/':
			switch current.dir {
			case UP:
				queue = append(queue, next(current.p, RIGHT))
			case DOWN:
				queue = append(queue, next(current.p, LEFT))
			case RIGHT:
				queue = append(queue, next(current.p, UP))
			case LEFT:
				queue = append(queue, next(current.p, DOWN))
			}
		case '\\':
			switch current.dir {
			case UP:
				queue = append(queue, next(current.p, LEFT))
			case DOWN:
				queue = append(queue, next(current.p, RIGHT))
			case RIGHT:
				queue = append(queue, next(current.p, DOWN))
			case LEFT:
				queue = append(queue, next(current.p, UP))
			}
		case '|':
			switch current.dir {
			case LEFT, RIGHT:
				queue = append(queue, next(current.p, UP), next(current.p, DOWN))
			default:
				queue = append(queue, next(current.p, current.dir))
			}
		case '-':
			switch current.dir {
			case UP, DOWN:
				queue = append(queue, next(current.p, LEFT), next(current.p, RIGHT))
			default:
				queue = append(queue, next(current.p, current.dir))
			}
		}
	}

	return len(energised)
}

func Part1(input string) (int, error) {
	grid, _, _ := parseGrid(input)

	return castBeam(grid, beam{point{0, 0}, RIGHT}), nil
}

func Part2(input string) (int, error) {
	grid, width, height := parseGrid(input)

	var startingBeams []beam

	for x := 0; x < width; x++ {
		startingBeams = append(startingBeams, beam{point{x, 0}, DOWN}, beam{point{x, height - 1}, UP})
	}

	for y := 0; y < height; y++ {
		startingBeams = append(startingBeams, beam{point{0, y}, RIGHT}, beam{point{width - 1, y}, LEFT})
	}

	maxEnergised := 0

	for _, b := range startingBeams {
		energised := castBeam(grid, b)
		maxEnergised = max(energised, maxEnergised)
	}

	return maxEnergised, nil
}

//go:embed input.txt
var input string

func main() {
	util.Run(Part1, Part2, input)
}
