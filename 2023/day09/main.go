package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"adtennant.dev/aoc/util"
)

func parseSequence(str string) (seq []int, err error) {
	for _, part := range strings.Split(str, " ") {
		n, err := strconv.Atoi(part)
		if err != nil {
			return nil, err
		}

		seq = append(seq, n)
	}

	return seq, nil
}

func parseSequences(input string) (seqs [][]int, err error) {
	for _, line := range util.Lines(input) {
		seq, err := parseSequence(line)
		if err != nil {
			return nil, fmt.Errorf("failed to parse sequence: %w", err)
		}

		seqs = append(seqs, seq)
	}

	return seqs, nil
}

func solve(seqs [][]int) int {
	sum := 0

	for _, seq := range seqs {
		current := seq

		var stack [][]int
		stack = append(stack, current)

		for {
			diff := make([]int, len(current)-1)

			for i := 0; i < len(current)-1; i++ {
				diff[i] = current[i+1] - current[i]
			}

			done := true

			for _, d := range diff {
				if d != 0 {
					done = false
					break
				}
			}

			if done {
				break
			}

			stack = append(stack, diff)
			current = diff
		}

		next := 0

		for _, diff := range stack {
			next += diff[len(diff)-1]
		}

		sum += next
	}

	return sum
}
func Part1(input string) (int, error) {
	seqs, err := parseSequences(input)
	if err != nil {
		return -1, err
	}

	return solve(seqs), nil
}

func Part2(input string) (int, error) {
	seqs, err := parseSequences(input)
	if err != nil {
		return -1, err
	}

	for i := range seqs {
		slices.Reverse(seqs[i])
	}

	return solve(seqs), nil
}

//go:embed input.txt
var input string

func main() {
	util.Run(Part1, Part2, input)
}
