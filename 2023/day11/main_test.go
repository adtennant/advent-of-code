package main

import (
	"testing"

	"adtennant.dev/aoc/util"
)

const exampleInput = `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`

func Test_Day11(t *testing.T) {
	util.RunTests(t, util.Part[int64]{
		Solution: Part1,
		Tests: []util.Test[int64]{
			{
				Input:    exampleInput,
				Expected: 374,
			},
		},
	}, util.Part[int64]{
		Solution: Part2,
		Tests: []util.Test[int64]{
			{
				Input:    exampleInput,
				Expected: 82000210,
			},
		},
	})
}
