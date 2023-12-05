package util

import (
	"fmt"
	"strings"
)

type Result interface {
	int | int64
}

type Solution[T Result] func(string) (T, error)

func sanitize(input string) string {
	var trimmed []string

	for _, line := range Lines(input) {
		trimmed = append(trimmed, strings.TrimSpace(line))
	}

	return strings.Join(trimmed, "\n")
}

func Main[T Result](rawInput string, part1, part2 Solution[T]) {
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
