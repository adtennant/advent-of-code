package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"

	"adtennant.dev/aoc/util"
)

func parseNumber(value string) (int, error) {
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

func parseNumberFromLine(re *regexp.Regexp, line string) (int, error) {
	matches := re.FindStringSubmatch(line)
	if len(matches) < 2 {
		return -1, fmt.Errorf("failed to find number: %s", line)
	}

	num, err := parseNumber(re.FindStringSubmatch(line)[1])
	if err != nil {
		return -1, fmt.Errorf("failed to parse number: %s: %w", line, err)
	}

	return num, nil
}

func solve(input string, exp1, exp2 string) (int, error) {
	re1, err := regexp.Compile(exp1)
	if err != nil {
		return -1, fmt.Errorf("failed to compile exp1: %w", err)
	}

	re2, err := regexp.Compile(exp2)
	if err != nil {
		return -1, fmt.Errorf("failed to compile exp2: %w", err)
	}

	sum := 0

	for _, line := range util.Lines(input) {
		first, err := parseNumberFromLine(re1, line)
		if err != nil {
			return -1, fmt.Errorf("failed to parse first number from line: %s: %w", line, err)
		}

		last, err := parseNumberFromLine(re2, line)
		if err != nil {
			return -1, fmt.Errorf("failed to parse last number from line: %s: %w", line, err)
		}

		value := first*10 + last
		sum += value
	}

	return sum, nil
}

func Part1(input string) (int, error) {
	return solve(input, `^\D*(\d).*$`, `^.*(\d)\D*$`)
}

func Part2(input string) (int, error) {
	return solve(input, `.*?(one|two|three|four|five|six|seven|eight|nine|\d).*`, `.*(one|two|three|four|five|six|seven|eight|nine|\d)`)
}

//go:embed input.txt
var input string

func main() {
	util.Main(input, Part1, Part2)
}
