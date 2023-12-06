package main

import (
	"testing"

	"adtennant.dev/aoc/util"
)

func Test_Day1(t *testing.T) {
	util.RunTests(t, util.Part[int]{
		Solution: Part1,
		Tests: []util.Test[int]{
			{
				Input: `1abc2
				pqr3stu8vwx
				a1b2c3d4e5f
				treb7uchet`,
				Expected: 142,
			},
		},
	}, util.Part[int]{
		Solution: Part2,
		Tests: []util.Test[int]{
			{
				Input: `1abc2
				pqr3stu8vwx
				a1b2c3d4e5f
				treb7uchet`,
				Expected: 142,
			},
		},
	})
}
