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

const exampleInput = `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`

func Test_Part1(t *testing.T) {
	util.Tests[int64]{
		{
			Input:    exampleInput,
			Expected: 35,
		},
	}.Run(t, Part1)
}

func Test_Part2(t *testing.T) {
	util.Tests[int64]{
		{
			Input:    exampleInput,
			Expected: 46,
		},
	}.Run(t, Part2)
}
