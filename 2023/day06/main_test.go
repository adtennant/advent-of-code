package main

import (
	"testing"

	"adtennant.dev/aoc/util"
)

func Test_Part1(t *testing.T) {
	util.Tests[int]{
		{
			Input: `Time:      7  15   30
			Distance:  9  40  200`,
			Expected: 288,
		},
	}.Run(t, Part1)
}

func Test_Part2(t *testing.T) {
	util.Tests[int]{
		{
			Input: `Time:      7  15   30
			Distance:  9  40  200`,
			Expected: 71503,
		},
	}.Run(t, Part2)
}
