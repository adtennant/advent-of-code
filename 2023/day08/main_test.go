package main

import (
	"testing"

	"adtennant.dev/aoc/util"
)

const exampleInput = ``

func Test_Day8(t *testing.T) {
	util.RunTests(t, util.Part[int]{
		Solution: Part1,
		Tests: []util.Test[int]{
			{
				Input: `RL

				AAA = (BBB, CCC)
				BBB = (DDD, EEE)
				CCC = (ZZZ, GGG)
				DDD = (DDD, DDD)
				EEE = (EEE, EEE)
				GGG = (GGG, GGG)
				ZZZ = (ZZZ, ZZZ)`,
				Expected: 2,
			},
			{
				Input: `LLR

				AAA = (BBB, BBB)
				BBB = (AAA, ZZZ)
				ZZZ = (ZZZ, ZZZ)`,
				Expected: 6,
			},
		},
	}, util.Part[int]{
		Solution: Part2,
		Tests: []util.Test[int]{
			{
				Input: `LR

				11A = (11B, XXX)
				11B = (XXX, 11Z)
				11Z = (11B, XXX)
				22A = (22B, XXX)
				22B = (22C, 22C)
				22C = (22Z, 22Z)
				22Z = (22B, 22B)
				XXX = (XXX, XXX)`,
				Expected: 6,
			},
		},
	})
}
