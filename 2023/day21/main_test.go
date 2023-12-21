package main

import (
	"testing"

	"adtennant.dev/aoc/util"
)

const exampleInput = `...........
.....###.#.
.###.##..#.
..#.#...#..
....#.#....
.##..S####.
.##..#...#.
.......##..
.##.#.####.
.##..##.##.
...........`

func Test_DayX(t *testing.T) {
	util.RunTests(t, util.Part[int]{
		Solution: Part1(6),
		Tests: []util.Test[int]{
			{
				Input:    exampleInput,
				Expected: 16,
			},
		},
	}, util.Part[int]{
		Solution: Part2(5000),
		Tests:    nil,
	})
}
