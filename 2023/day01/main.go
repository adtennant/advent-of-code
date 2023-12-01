package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"

	"adtennant.dev/aoc/util"
)

func Parse(value string) (int, error) {
	switch value {
	case "one":
		return 1, nil
	case "two":
		return 2, nil
	case "three":
		return 3, nil
	case "four":
		return 4, nil
	case "five":
		return 5, nil
	case "six":
		return 6, nil
	case "seven":
		return 7, nil
	case "eight":
		return 8, nil
	case "nine":
		return 9, nil
	default:
		return strconv.Atoi(value)
	}
}

func Solve(input string, exp1, exp2 string) int {
	re1 := regexp.MustCompile(exp1)
	re2 := regexp.MustCompile(exp2)

	sum := 0

	for _, line := range util.Lines(input) {
		if len(line) == 0 {
			continue
		}

		first, _ := Parse(re1.FindStringSubmatch(line)[1])
		last, _ := Parse(re2.FindStringSubmatch(line)[1])

		value := first*10 + last
		sum += value
	}

	return sum
}

func Part1(input string) int {
	return Solve(input, `^\D*(\d).*$`, `^.*(\d)\D*$`)
}

func Part2(input string) int {
	return Solve(input, `.*?(one|two|three|four|five|six|seven|eight|nine|\d).*`, `.*(one|two|three|four|five|six|seven|eight|nine|\d)`)
}

//go:embed input.txt
var input string

func main() {
	input := util.Sanitize(input)

	fmt.Println("Part 1 =", Part1(input))
	fmt.Println("Part 2 =", Part2(input))
}
