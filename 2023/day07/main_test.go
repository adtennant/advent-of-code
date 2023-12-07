package main

import (
	"testing"

	"adtennant.dev/aoc/util"
)

const exampleInput = `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`

func Test_Day7(t *testing.T) {
	util.RunTests(t, util.Part[int]{
		Solution: Part1,
		Tests: []util.Test[int]{
			{
				Input:    exampleInput,
				Expected: 6440,
			},
		},
	}, util.Part[int]{
		Solution: Part2,
		Tests: []util.Test[int]{
			{
				Input:    exampleInput,
				Expected: 5905,
			},
		},
	})
}
