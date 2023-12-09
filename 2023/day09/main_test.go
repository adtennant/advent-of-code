package main

import (
	"testing"

	"adtennant.dev/aoc/util"
)

const exampleInput = `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`

func Test_Day9(t *testing.T) {
	util.RunTests(t, util.Part[int]{
		Solution: Part1,
		Tests: []util.Test[int]{
			{
				Input:    exampleInput,
				Expected: 114,
			},
		},
	}, util.Part[int]{
		Solution: Part2,
		Tests: []util.Test[int]{
			{
				Input:    exampleInput,
				Expected: 2,
			},
		},
	})
}
