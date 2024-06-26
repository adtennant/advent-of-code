package main

import (
	"testing"

	"adtennant.dev/aoc/util"
)

func Test_Day10(t *testing.T) {
	util.RunTests(t, util.Part[int]{
		Solution: Part1,
		Tests: []util.Test[int]{
			{
				Input: `.....
				.S-7.
				.|.|.
				.L-J.
				.....`,
				Expected: 4,
			},
			{
				Input: `..F7.
				.FJ|.
				SJ.L7
				|F--J
				LJ...`,
				Expected: 8,
			},
		},
	}, util.Part[int]{
		Solution: Part2,
		Tests: []util.Test[int]{
			{
				Input: `...........
				.S-------7.
				.|F-----7|.
				.||.....||.
				.||.....||.
				.|L-7.F-J|.
				.|..|.|..|.
				.L--J.L--J.
				...........`,
				Expected: 4,
			},
			{
				Input: `.F----7F7F7F7F-7....
				.|F--7||||||||FJ....
				.||.FJ||||||||L7....
				FJL7L7LJLJ||LJ.L-7..
				L--J.L7...LJS7F-7L7.
				....F-J..F7FJ|L7L7L7
				....L7.F7||L7|.L7L7|
				.....|FJLJ|FJ|F7|.LJ
				....FJL-7.||.||||...
				....L---J.LJ.LJLJ...`,
				Expected: 8,
			},
			{
				Input: `FF7FSF7F7F7F7F7F---7
				L|LJ||||||||||||F--J
				FL-7LJLJ||||||LJL-77
				F--JF--7||LJLJ7F7FJ-
				L---JF-JLJ.||-FJLJJ7
				|F|F-JF---7F7-L7L|7|
				|FFJF7L7F-JF7|JL---7
				7-L-JL7||F7|L7F-7F7|
				L.L7LFJ|||||FJL7||LJ
				L7JLJL-JLJLJL--JLJ.L`,
				Expected: 10,
			},
		},
	})
}
