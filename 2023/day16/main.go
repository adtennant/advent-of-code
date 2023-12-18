package main

import (
	_ "embed"

	"adtennant.dev/aoc/util"
)

type point = util.Point[int]

func parseGrid(input string) (map[point]byte, int, int) {
	lines := util.Lines(input)
	grid := make(map[point]byte)

	for y, line := range lines {
		for x, c := range []byte(line) {
			grid[point{X: x, Y: y}] = c
		}
	}

	width := len(lines[0])
	height := len(lines)

	return grid, width, height
}

type beam struct {
	pos point
	dir util.Direction
}

func next(current point, dir util.Direction) beam {
	delta := util.GetDelta[int](dir)

	return beam{
		current.Translate(delta),
		dir,
	}
}

func castBeam(grid map[point]byte, b beam) int {
	q := util.NewQueue[beam]()
	q.Push(b)

	seen := util.NewSet[beam]()
	energised := util.NewSet[point]()

	for q.Len() > 0 {
		current := q.Pop()

		if _, ok := grid[current.pos]; !ok {
			continue
		}

		if seen.Contains(current) {
			continue
		}

		seen.Insert(current)
		energised.Insert(current.pos)

		switch grid[current.pos] {
		case '.':
			q.Push(next(current.pos, current.dir))
		case '/':
			switch current.dir {
			case util.UP:
				q.Push(next(current.pos, util.RIGHT))
			case util.DOWN:
				q.Push(next(current.pos, util.LEFT))
			case util.RIGHT:
				q.Push(next(current.pos, util.UP))
			case util.LEFT:
				q.Push(next(current.pos, util.DOWN))
			}
		case '\\':
			switch current.dir {
			case util.UP:
				q.Push(next(current.pos, util.LEFT))
			case util.DOWN:
				q.Push(next(current.pos, util.RIGHT))
			case util.RIGHT:
				q.Push(next(current.pos, util.DOWN))
			case util.LEFT:
				q.Push(next(current.pos, util.UP))
			}
		case '|':
			switch current.dir {
			case util.LEFT, util.RIGHT:
				q.Push(next(current.pos, util.UP), next(current.pos, util.DOWN))
			default:
				q.Push(next(current.pos, current.dir))
			}
		case '-':
			switch current.dir {
			case util.UP, util.DOWN:
				q.Push(next(current.pos, util.LEFT), next(current.pos, util.RIGHT))
			default:
				q.Push(next(current.pos, current.dir))
			}
		}
	}

	return energised.Len()
}

func Part1(input string) (int, error) {
	grid, _, _ := parseGrid(input)

	return castBeam(grid, beam{point{X: 0, Y: 0}, util.RIGHT}), nil
}

func Part2(input string) (int, error) {
	grid, width, height := parseGrid(input)

	var startingBeams []beam

	for x := 0; x < width; x++ {
		startingBeams = append(
			startingBeams,
			beam{point{X: x, Y: 0}, util.DOWN},
			beam{point{X: x, Y: height - 1}, util.UP},
		)
	}

	for y := 0; y < height; y++ {
		startingBeams = append(
			startingBeams,
			beam{point{X: 0, Y: y}, util.RIGHT},
			beam{point{X: width - 1, Y: y}, util.LEFT},
		)
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
