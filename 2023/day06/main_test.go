package main

import (
	"testing"

	"adtennant.dev/aoc/util"
)

const exampleInput = `Time:      7  15   30
Distance:  9  40  200`

func Test_Day6(t *testing.T) {
	util.RunTests(t, util.Part[int]{
		Solution: Part1,
		Tests: []util.Test[int]{
			{
				Input:    exampleInput,
				Expected: 288,
			},
		},
	}, util.Part[int]{
		Solution: Part2,
		Tests: []util.Test[int]{
			{
				Input:    exampleInput,
				Expected: 71503,
			},
		},
	})
}
