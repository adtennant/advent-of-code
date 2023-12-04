package main

import (
	_ "embed"
	"fmt"
	"strings"

	"adtennant.dev/aoc/util"
)

type ExtractFunc func(string) (int, bool)

func extractNumber(str string) (int, bool) {
	if str[0] >= '0' && str[0] <= '9' {
		return int(str[0] - '0'), true
	}

	return -1, false
}

func extractNumberOrWord(str string) (int, bool) {
	if num, found := extractNumber(str); found {
		return num, true
	}

	switch {
	case strings.HasPrefix(str, "zero"):
		return 0, true
	case strings.HasPrefix(str, "one"):
		return 1, true
	case strings.HasPrefix(str, "two"):
		return 2, true
	case strings.HasPrefix(str, "three"):
		return 3, true
	case strings.HasPrefix(str, "four"):
		return 4, true
	case strings.HasPrefix(str, "five"):
		return 5, true
	case strings.HasPrefix(str, "six"):
		return 6, true
	case strings.HasPrefix(str, "seven"):
		return 7, true
	case strings.HasPrefix(str, "eight"):
		return 8, true
	case strings.HasPrefix(str, "nine"):
		return 9, true
	default:
		return -1, false
	}
}

func extractFirst(str string, extract ExtractFunc) (int, bool) {
	for i := 0; i < len(str); i++ {
		if num, found := extract(str[i:]); found {
			return num, true
		}
	}

	return -1, false
}

func extractLast(str string, extract ExtractFunc) (int, bool) {
	for i := len(str) - 1; i >= 0; i-- {
		if num, found := extract(str[i:]); found {
			return num, true
		}
	}

	return -1, false
}

func solve(input string, extract ExtractFunc) (int, error) {
	sum := 0

	for _, line := range util.Lines(input) {
		first, found := extractFirst(line, extract)
		if !found {
			return -1, fmt.Errorf("failed to find first number in line: %s", line)
		}

		last, found := extractLast(line, extract)
		if !found {
			return -1, fmt.Errorf("failed to find last number in line: %s", line)
		}

		value := first*10 + last
		sum += value
	}

	return sum, nil
}

func Part1(input string) (int, error) {
	return solve(input, extractNumber)
}

func Part2(input string) (int, error) {
	return solve(input, extractNumberOrWord)
}

//go:embed input.txt
var input string

func main() {
	util.Main(input, Part1, Part2)
}
