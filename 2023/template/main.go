package main

import (
	_ "embed"

	"adtennant.dev/aoc/util"
)

func Part1(input string) (int, error) {
	return 0, nil
}

func Part2(input string) (int, error) {
	return 0, nil
}

//go:embed input.txt
var input string

func main() {
	util.Main(input, Part1, Part2)
}
