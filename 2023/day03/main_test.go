package main

import (
	"testing"

	"adtennant.dev/aoc/util"
)

func Benchmark_Part1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part1(input)
	}
}

func Benchmark_Part2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part2(input)
	}
}

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
	util.Tests[int]{
		{
			Input:    exampleInput,
			Expected: 4361,
		},
	}.Run(t, Part1)
}

func Test_Part2(t *testing.T) {
	util.Tests[int]{
		{
			Input:    exampleInput,
			Expected: 467835,
		},
	}.Run(t, Part2)
}
