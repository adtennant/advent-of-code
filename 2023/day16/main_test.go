package main

import (
	"testing"

	"adtennant.dev/aoc/util"
)

const exampleInput = `.|...\....
|.-.\.....
.....|-...
........|.
..........
.........\
..../.\\..
.-.-/..|..
.|....-|.\
..//.|....`

func Test_Day16(t *testing.T) {
	util.RunTests(t, util.Part[int]{
		Solution: Part1,
		Tests: []util.Test[int]{
			{
				Input:    exampleInput,
				Expected: 46,
			},
		},
	}, util.Part[int]{
		Solution: Part2,
		Tests: []util.Test[int]{
			{
				Input:    exampleInput,
				Expected: 51,
			},
		},
	})
}
