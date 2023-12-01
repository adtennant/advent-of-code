package main

import (
	_ "embed"
	"fmt"

	"adtennant.dev/aoc/util"
)

func Part1(input string) int {
	return 0
}

func Part2(input string) int {
	return 0
}

//go:embed input.txt
var input string

func main() {
	input := util.Sanitize(input)

	fmt.Println("Part 1 =", Part1(input))
	fmt.Println("Part 2 =", Part2(input))
}
