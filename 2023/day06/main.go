package main

import (
	_ "embed"
	"fmt"
	"math"
	"regexp"
	"strconv"

	"adtennant.dev/aoc/util"
)

type race struct {
	time     int
	distance int
}

type ParseLine func(string) ([]int, error)

func parseRaces(input string, parseLine ParseLine) (races []race, err error) {
	lines := util.Lines(input)

	times, err := parseLine(lines[0])
	if err != nil {
		return nil, fmt.Errorf("failed to parse times: %w", err)
	}

	distances, err := parseLine(lines[1])
	if err != nil {
		return nil, fmt.Errorf("failed to parse distances: %w", err)
	}

	if len(times) != len(distances) {
		return nil, fmt.Errorf("mismatch between number of times and distances")
	}

	for i := 0; i < len(times); i++ {
		races = append(races, race{
			time:     times[i],
			distance: distances[i],
		})
	}

	return races, nil
}

func countWins(r race) int {
	a := float64(-1)
	b := float64(r.time)
	c := -float64(r.distance)

	d := b*b - 4*a*c

	lower := (-b + math.Sqrt(d)) / (2 * a)
	upper := (-b - math.Sqrt(d)) / (2 * a)

	return int(math.Ceil(upper-1) - math.Floor(lower+1) + 1)
}

func solve(input string, parseLine ParseLine) (int, error) {
	races, err := parseRaces(input, parseLine)
	if err != nil {
		return -1, err
	}

	product := 1

	for _, r := range races {
		product *= countWins(r)
	}

	return product, nil
}

var numberRegexp = regexp.MustCompile(`\b(\d+)\b`)

func Part1(input string) (int, error) {
	return solve(input, func(str string) (nums []int, err error) {
		matches := numberRegexp.FindAllStringSubmatch(str, -1)

		for _, m := range matches {
			n, _ := strconv.Atoi(m[0])
			nums = append(nums, n)
		}

		return nums, nil
	})
}

func Part2(input string) (int, error) {
	return solve(input, func(str string) (nums []int, err error) {
		matches := numberRegexp.FindAllStringSubmatch(str, -1)
		num := ""

		for _, m := range matches {
			num += m[0]
		}

		n, _ := strconv.Atoi(num)
		return []int{n}, nil
	})
}

//go:embed input.txt
var input string

func main() {
	util.Run(Part1, Part2, input)
}
