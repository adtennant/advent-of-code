package main

import (
	"testing"

	"adtennant.dev/aoc/util"
)

const exampleInput = ``

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
			Expected: 0,
		},
	}.Run(t, Part1)
}

func Test_Part2(t *testing.T) {
	util.Tests[int]{
		{
			Input:    exampleInput,
			Expected: 0,
		},
	}.Run(t, Part2)
}
