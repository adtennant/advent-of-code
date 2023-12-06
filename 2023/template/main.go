package main

import (
	_ "embed"

	"adtennant.dev/aoc/util"
)

func Part1(input string) (int, error) {
	return -1, nil
}

func Part2(input string) (int, error) {
	return -1, nil
}

//go:embed input.txt
var input string

func main() {
	util.Run(Part1, Part2, input)
}
