package main

import (
	"testing"

	"adtennant.dev/aoc/util"
)

func Test_Part1(t *testing.T) {
	util.Tests[int]{
		{
			Input: `1abc2
			pqr3stu8vwx
			a1b2c3d4e5f
			treb7uchet`,
			Expected: 142,
		},
	}.Run(t, Part1)
}

func Test_Part2(t *testing.T) {
	util.Tests[int]{
		{
			Input: `two1nine
			eightwothree
			abcone2threexyz
			xtwone3four
			4nineeightseven2
			zoneight234
			7pqrstsixteen`,
			Expected: 281,
		},
	}.Run(t, Part2)
}
