package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"adtennant.dev/aoc/util"
)

type record struct {
	springs []byte
	groups  []int
}

func parseRecord(str string) (record, error) {
	parts := strings.Split(str, " ")
	if len(parts) != 2 {
		return record{}, fmt.Errorf("invalid format")
	}

	springs := []byte(parts[0])

	var groups []int

	for _, n := range strings.Split(parts[1], ",") {
		group, err := strconv.Atoi(n)
		if err != nil {
			return record{}, err
		}

		groups = append(groups, group)
	}

	return record{springs, groups}, nil
}

func parseRecords(input string) (records []record, err error) {
	for _, line := range util.Lines(input) {
		record, err := parseRecord(line)
		if err != nil {
			return nil, fmt.Errorf("failed to parse record: %w", err)
		}

		records = append(records, record)
	}

	return records, nil
}

func unfold(r record) record {
	var springs []byte
	var groups []int

	for i := 0; i < 5; i++ {
		springs = append(springs, r.springs...)
		springs = append(springs, '?')

		groups = append(groups, r.groups...)
	}

	return record{
		springs[:len(springs)-1],
		groups,
	}
}

var cache = make(map[string]int)

func findArrangements(springs []byte, groups []int) int {
	key := fmt.Sprintf("%v-%v", springs, groups)

	if v, ok := cache[key]; ok {
		return v
	}

	count := func() int {
		if len(springs) == 0 {
			if len(groups) == 0 {
				return 1
			} else {
				return 0
			}
		} else if springs[0] == '.' {
			return findArrangements(springs[1:], groups)
		} else if springs[0] == '?' {
			operational := make([]byte, len(springs))
			copy(operational, springs)
			operational[0] = '#'

			return findArrangements(operational, groups) + findArrangements(springs[1:], groups)
		} else if len(groups) == 0 && slices.Contains(springs, '#') || len(springs) < groups[0] || slices.Contains(springs[:groups[0]], '.') {
			return 0
		} else if len(springs) == groups[0] {
			if len(groups) == 1 {
				return 1
			} else {
				return 0
			}
		} else if springs[groups[0]] == '#' {
			return 0
		} else {
			return findArrangements(springs[groups[0]+1:], groups[1:])
		}
	}()

	cache[key] = count
	return count
}

func Part1(input string) (int, error) {
	records, err := parseRecords(input)
	if err != nil {
		return -1, err
	}

	sum := 0

	for _, r := range records {
		sum += findArrangements(r.springs, r.groups)
	}

	return sum, nil
}

func Part2(input string) (int, error) {
	records, err := parseRecords(input)
	if err != nil {
		return -1, err
	}

	sum := 0

	for _, r := range records {
		u := unfold(r)
		sum += findArrangements(u.springs, u.groups)
	}

	return sum, nil
}

//go:embed input.txt
var input string

func main() {
	util.Run(Part1, Part2, input)
}
