package main

import (
	"testing"

	"adtennant.dev/aoc/util"
)

const exampleInput = `???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1`

func Test_Day12(t *testing.T) {
	util.RunTests(t, util.Part[int]{
		Solution: Part1,
		Tests: []util.Test[int]{
			{
				Input:    exampleInput,
				Expected: 21,
			},
		},
	}, util.Part[int]{
		Solution: Part2,
		Tests: []util.Test[int]{
			{
				Input:    exampleInput,
				Expected: 525152,
			},
		},
	})
}
