package main

import (
	"testing"

	"adtennant.dev/aoc/util"
)

const exampleInput = `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`

func Benchmark_Part1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part1(exampleInput)
	}
}

func Benchmark_Part2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part2(exampleInput)
	}
}

func Test_Part1(t *testing.T) {
	util.Tests[int]{
		{
			Input:    exampleInput,
			Expected: 142,
		},
	}.Run(t, Part1)
}

func Test_Part2(t *testing.T) {
	util.Tests[int]{
		{
			Input:    exampleInput,
			Expected: 281,
		},
	}.Run(t, Part2)
}
