package util

import (
	"fmt"
	"strings"
)

type Solution func(string) (int, error)

func sanitize(input string) string {
	var trimmed []string

	for _, line := range Lines(input) {
		trimmed = append(trimmed, strings.TrimSpace(line))
	}

	return strings.Join(trimmed, "\n")
}

func Main(rawInput string, part1, part2 Solution) {
	input := sanitize(rawInput)

	result1, err := part1(input)
	if err != nil {
		panic(err)
	}

	fmt.Println("Part 1 =", result1)

	result2, err := part2(input)
	if err != nil {
		panic(err)
	}

	fmt.Println("Part 2 =", result2)
}
