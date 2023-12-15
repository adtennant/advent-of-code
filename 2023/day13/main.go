package main

import (
	_ "embed"
	"fmt"
	"math/bits"
	"strconv"
	"strings"

	"adtennant.dev/aoc/util"
)

type pattern struct {
	rows []uint
	cols []uint
}

func convertToBits(rowOrCol string) (uint, error) {
	bits, err := strconv.ParseInt(strings.ReplaceAll(strings.ReplaceAll(rowOrCol, ".", "0"), "#", "1"), 2, 32)
	if err != nil {
		return 0, err
	}

	return uint(bits), nil
}

func extractPattern(lines []string) (pattern, error) {
	var rbits []uint

	for _, line := range lines {
		if line == "" {
			break
		}

		bits, err := convertToBits(line)
		if err != nil {
			return pattern{}, fmt.Errorf("failed to parse row: %w", err)
		}

		rbits = append(rbits, uint(bits))
	}

	var cbits []uint

	for x := 0; x < len(lines[0]); x++ {
		var col []byte

		for i := 0; i < len(rbits); i++ {
			col = append(col, lines[i][x])
		}

		bits, err := convertToBits(string(col))
		if err != nil {
			return pattern{}, fmt.Errorf("failed to parse column: %w", err)
		}

		cbits = append(cbits, uint(bits))
	}

	return pattern{rbits, cbits}, nil
}

func parsePatterns(input string) (patterns []pattern, err error) {
	lines := util.Lines(input)

	for i := 0; i < len(lines); i++ {
		pattern, err := extractPattern(lines[i:])
		if err != nil {
			return nil, fmt.Errorf("failed to parse pattern: %w", err)
		}

		patterns = append(patterns, pattern)

		i += len(pattern.rows)
	}

	return patterns, nil
}

func sumMirrorIndexes(rowsOrCols []uint, targetDiff int) int {
	sum := 0

	for i := 0; i < len(rowsOrCols)-1; i++ {
		width := min(i+1, len(rowsOrCols)-i-1)
		gap := 1

		diff := 0

		for j := i; j > i-width; j-- {
			diff += bits.OnesCount(rowsOrCols[j] ^ rowsOrCols[j+gap])
			gap += 2
		}

		if diff == targetDiff {
			sum += i + 1
		}
	}

	return sum
}

func solve(patterns []pattern, targetDiff int) int {
	cols := 0
	rows := 0

	for _, p := range patterns {
		cols += sumMirrorIndexes(p.cols, targetDiff)
		rows += sumMirrorIndexes(p.rows, targetDiff)
	}

	return rows*100 + cols
}

func Part1(input string) (int, error) {
	patterns, err := parsePatterns(input)
	if err != nil {
		return -1, err
	}

	return solve(patterns, 0), nil
}

func Part2(input string) (int, error) {
	patterns, err := parsePatterns(input)
	if err != nil {
		return -1, err
	}

	return solve(patterns, 1), nil
}

//go:embed input.txt
var input string

func main() {
	util.Run(Part1, Part2, input)
}
