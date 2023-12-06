package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"

	"adtennant.dev/aoc/util"
)

type scratchcard struct {
	winningNumbers []int
	numbers        []int
}

func parseNumbers(str string) (numbers []int, err error) {
	nums := strings.Split(str, " ")

	for _, num := range nums {
		if strings.TrimSpace(num) == "" {
			continue
		}

		v, err := strconv.Atoi(strings.TrimSpace(num))
		if err != nil {
			return nil, err
		}

		numbers = append(numbers, v)
	}

	return numbers, nil
}

func parseScratchcard(line string) (scratchcard, error) {
	parts := strings.Split(line, ":")
	if len(parts) != 2 {
		return scratchcard{}, fmt.Errorf("invalid format")
	}

	parts = strings.Split(parts[1], "|")
	if len(parts) != 2 {
		return scratchcard{}, fmt.Errorf("invalid format")
	}

	winningNumbers, err := parseNumbers(strings.TrimSpace(parts[0]))
	if err != nil {
		return scratchcard{}, fmt.Errorf("failed to parse winning numbers: %w", err)
	}

	numbers, err := parseNumbers(strings.TrimSpace(parts[1]))
	if err != nil {
		return scratchcard{}, fmt.Errorf("failed to parse numbers: %w", err)
	}

	return scratchcard{
		winningNumbers,
		numbers,
	}, nil
}

func parseScratchcards(input string) (cards []scratchcard, err error) {
	for _, line := range util.Lines(input) {
		card, err := parseScratchcard(line)
		if err != nil {
			return nil, fmt.Errorf("failed to parse scratchcard from line: %s: %w", line, err)
		}

		cards = append(cards, card)
	}

	return cards, nil
}

func intersection(s1, s2 []int) (i []int) {
	lookup := make(map[int]bool)

	for _, v := range s1 {
		lookup[v] = true
	}

	for _, v := range s2 {
		if lookup[v] {
			i = append(i, v)
		}
	}

	return i
}

func Part1(input string) (int, error) {
	cards, err := parseScratchcards(input)
	if err != nil {
		return -1, err
	}

	points := 0

	for _, card := range cards {
		wins := intersection(card.numbers, card.winningNumbers)

		if len(wins) > 0 {
			points += int(1 * math.Pow(2, float64(len(wins)-1)))
		}
	}

	return points, nil
}

func Part2(input string) (int, error) {
	cards, err := parseScratchcards(input)
	if err != nil {
		return -1, err
	}

	copies := make([]int, len(cards))

	for i, card := range cards {
		wins := intersection(card.numbers, card.winningNumbers)

		for j := i + 1; j <= i+len(wins); j++ {
			copies[j] += 1 + copies[i]
		}
	}

	total := len(cards)

	for _, copy := range copies {
		total += copy
	}

	return total, nil
}

//go:embed input.txt
var input string

func main() {
	util.Run(Part1, Part2, input)
}
