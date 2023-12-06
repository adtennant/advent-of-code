package util

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

type Result interface {
	int | int64
}

type Test[T Result] struct {
	Name     string
	Input    string
	Expected T
}

type Solution[T Result] func(string) (T, error)

type Part[T Result] struct {
	Solution Solution[T]
	Tests    []Test[T]
}

func sanitize(input string) string {
	var trimmed []string

	for _, line := range Lines(input) {
		trimmed = append(trimmed, strings.TrimSpace(line))
	}

	return strings.Join(trimmed, "\n")
}

func runSolution[T Result](solution Solution[T], rawInput string, part int) {
	input := sanitize(rawInput)

	start := time.Now()
	result, err := solution(input)
	if err != nil {
		fmt.Printf("Part %d failed: %v", part, err)
	}

	elasped := time.Since(start)

	fmt.Printf("Part %d = %v (in %.3fms)\n", part, result, float64(elasped.Microseconds())/1000)
}

func Run[T Result](solution1, solution2 Solution[T], rawInput string) {
	runSolution(solution1, rawInput, 1)
	runSolution(solution2, rawInput, 2)
}

func runTests[T Result](t *testing.T, tests []Test[T], solution Solution[T], part int) {
	for _, tt := range tests {
		name := fmt.Sprintf("Part %d", part)

		if tt.Name != "" {
			name += fmt.Sprintf(" - %s", tt.Name)
		}

		t.Run(name, func(t *testing.T) {
			input := sanitize(tt.Input)
			actual, err := solution(input)

			if err != nil {
				t.Fatalf("%v", err)
			}

			if actual != tt.Expected {
				t.Fatalf("actual: %v, expected: %v", actual, tt.Expected)
			}
		})
	}
}

func RunTests[T Result](t *testing.T, part1, part2 Part[T]) {
	runTests(t, part1.Tests, part1.Solution, 1)
	runTests(t, part2.Tests, part2.Solution, 2)
}
