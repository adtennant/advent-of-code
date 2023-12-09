package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"unicode"

	"adtennant.dev/aoc/util"
)

type point struct {
	x int
	y int
}

type partNumber struct {
	value        int
	startX, endX int
	y            int
}

type schematic struct {
	parts   []partNumber
	symbols map[point]bool
	gears   map[point]bool
}

func extractPartNumber(str string) (int, error) {
	num := ""

	for i := range str {
		if unicode.IsDigit(rune(str[i])) {
			num += string(str[i])
		} else {
			break
		}
	}

	n, err := strconv.Atoi(num)
	if err != nil {
		return -1, err
	}

	return n, nil
}

func parseSchematic(input string) (s schematic, err error) {
	s.symbols = make(map[point]bool)
	s.gears = make(map[point]bool)

	for y, line := range util.Lines(input) {
		x := 0

		for x < len(line) {
			if line[x] >= '0' && line[x] <= '9' {
				num, err := extractPartNumber(line[x:])
				len := len(strconv.Itoa(num))

				if err != nil {
					return s, fmt.Errorf("failed to extract part number: %w", err)
				}

				part := partNumber{value: num, startX: x, endX: x + len - 1, y: y}
				s.parts = append(s.parts, part)

				x += len
			} else {
				if line[x] != '.' {
					if line[x] == '*' {
						s.gears[point{x, y}] = true
					} else {
						s.symbols[point{x, y}] = true
					}
				}

				x++
			}
		}
	}

	return s, nil
}

func Part1(input string) (int, error) {
	schematic, err := parseSchematic(input)
	if err != nil {
		return -1, err
	}

	sum := 0

	for _, part := range schematic.parts {
	loop:
		for y := part.y - 1; y <= part.y+1; y++ {
			for x := part.startX - 1; x <= part.endX+1; x++ {
				if schematic.symbols[point{x, y}] || schematic.gears[point{x, y}] {
					sum += part.value
					break loop
				}
			}
		}
	}

	return sum, nil
}

func Part2(input string) (int, error) {
	schematic, err := parseSchematic(input)
	if err != nil {
		return -1, err
	}

	partsByGear := make(map[point][]partNumber)

	for _, part := range schematic.parts {
	loop:
		for y := part.y - 1; y <= part.y+1; y++ {
			for x := part.startX - 1; x <= part.endX+1; x++ {
				if schematic.gears[point{x, y}] {
					partsByGear[point{x, y}] = append(partsByGear[point{x, y}], part)
					break loop
				}
			}
		}
	}

	sum := 0

	for _, parts := range partsByGear {
		if len(parts) == 2 {
			sum += parts[0].value * parts[1].value
		}
	}

	return sum, nil
}

//go:embed input.txt
var input string

func main() {
	util.Run(Part1, Part2, input)
}
