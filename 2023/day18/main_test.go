package main

import (
	"testing"

	"adtennant.dev/aoc/util"
)

const exampleInput = `R 6 (#70c710)
D 5 (#0dc571)
L 2 (#5713f0)
D 2 (#d2c081)
R 2 (#59c680)
D 2 (#411b91)
L 5 (#8ceee2)
U 2 (#caa173)
L 1 (#1b58a2)
U 2 (#caa171)
R 2 (#7807d2)
U 3 (#a77fa3)
L 2 (#015232)
U 2 (#7a21e3)`

func Test_Day18(t *testing.T) {
	util.RunTests(t, util.Part[int64]{
		Solution: Part1,
		Tests: []util.Test[int64]{
			{
				Input:    exampleInput,
				Expected: 62,
			},
		},
	}, util.Part[int64]{
		Solution: Part2,
		Tests: []util.Test[int64]{
			{
				Input:    exampleInput,
				Expected: 952408144115,
			},
		},
	})
}
