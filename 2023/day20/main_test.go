package main

import (
	"testing"

	"adtennant.dev/aoc/util"
)

func Test_Day20(t *testing.T) {
	util.RunTests(t, util.Part[int]{
		Solution: Part1,
		Tests: []util.Test[int]{
			{
				Input: `broadcaster -> a, b, c
				%a -> b
				%b -> c
				%c -> inv
				&inv -> a`,
				Expected: 32000000,
			},
			{
				Input: `broadcaster -> a
				%a -> inv, con
				&inv -> b
				%b -> con
				&con -> output`,
				Expected: 11687500,
			},
		},
	}, util.Part[int]{
		Solution: Part2,
		Tests:    nil,
	})
}
