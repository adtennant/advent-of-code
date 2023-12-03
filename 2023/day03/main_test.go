package main

import (
	"testing"

	"adtennant.dev/aoc/util"
)

const exampleInput = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

func Test_Part1(t *testing.T) {
	util.Tests{
		{
			Input:    exampleInput,
			Expected: 4361,
		},
	}.Run(t, Part1)
}

func Test_Part2(t *testing.T) {
	util.Tests{
		{
			Input:    exampleInput,
			Expected: 467835,
		},
	}.Run(t, Part2)
}
