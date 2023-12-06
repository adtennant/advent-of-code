package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"

	"adtennant.dev/aoc/util"
)

type mappingRange struct {
	dest int64
	src  int64
	len  int64
}

type almanac struct {
	seeds    []int64
	mappings [][]mappingRange
}

func parseSeeds(line string) (seeds []int64, err error) {
	parts := strings.Split(line, ": ")

	for _, s := range strings.Split(parts[1], " ") {
		v, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return nil, err
		}

		seeds = append(seeds, v)
	}

	return seeds, nil
}

func extractRanges(lines []string) (ranges []mappingRange, err error) {
	for _, line := range lines[1:] {
		if strings.TrimSpace(line) == "" {
			break
		}

		parts := strings.Split(line, " ")

		dest, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			return ranges, fmt.Errorf("failed to parse dest: %w", err)
		}

		src, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			return ranges, fmt.Errorf("failed to parse src: %w", err)
		}

		len, err := strconv.ParseInt(parts[2], 10, 64)
		if err != nil {
			return ranges, fmt.Errorf("failed to parse len: %w", err)
		}

		ranges = append(ranges, struct {
			dest int64
			src  int64
			len  int64
		}{dest, src, len})
	}

	return ranges, nil
}

func parseAlmanac(input string) (almanac, error) {
	lines := util.Lines(input)

	var seeds []int64
	var mappings [][]mappingRange

	for i := 0; i < len(lines); i++ {
		if strings.HasPrefix(lines[i], "seeds") {
			seeds, _ = parseSeeds(lines[i])
		} else if strings.HasSuffix(lines[i], "map:") {
			ranges, err := extractRanges(lines[i:])
			if err != nil {
				return almanac{}, fmt.Errorf("failed to extract ranges: %w", err)
			}

			mappings = append(mappings, ranges)

			i += len(ranges)
		}
	}

	return almanac{seeds, mappings}, nil
}

func seedToLocation(seed int64, mappings [][]mappingRange) int64 {
	for _, ranges := range mappings {
		for _, r := range ranges {
			if seed >= r.src && seed < r.src+r.len {
				seed = seed + (r.dest - r.src)
				break
			}
		}
	}

	return seed
}

func locationToSeed(location int64, mappings [][]mappingRange) int64 {
	for i := len(mappings) - 1; i >= 0; i-- {
		for _, r := range mappings[i] {
			if location >= r.dest && location < r.dest+r.len {
				location = location + (r.src - r.dest)
				break
			}
		}
	}

	return location
}

func Part1(input string) (int64, error) {
	almanac, err := parseAlmanac(input)
	if err != nil {
		return -1, err
	}

	minLocation := int64(math.MaxInt64)

	for _, seed := range almanac.seeds {
		minLocation = min(minLocation, seedToLocation(seed, almanac.mappings))
	}

	return minLocation, nil
}

func Part2(input string) (int64, error) {
	almanac, err := parseAlmanac(input)
	if err != nil {
		return -1, err
	}

	for i := int64(0); i < math.MaxInt64; i++ {
		seed := locationToSeed(i, almanac.mappings)

		for j := 0; j < len(almanac.seeds); j += 2 {
			start := almanac.seeds[j]
			len := almanac.seeds[j+1]

			if seed >= start && seed < start+len {
				return i, nil
			}
		}
	}

	return -1, fmt.Errorf("no seed found")
}

//go:embed input.txt
var input string

func main() {
	util.Run(Part1, Part2, input)
}
