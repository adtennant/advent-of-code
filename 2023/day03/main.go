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

func extractPartNumber(str string) (int, int) {
	num := ""
	i := 0

	for ; i < len(str); i++ {
		if unicode.IsDigit(rune(str[i])) {
			num += string(str[i])
		} else {
			break
		}
	}

	n, _ := strconv.Atoi(num)
	return n, i
}

func parseSchematic(input string) (parts []partNumber, symbols map[point]bool, gears map[point]bool) {
	symbols = make(map[point]bool)
	gears = make(map[point]bool)

	for y, line := range util.Lines(input) {
		x := 0

		for x < len(line) {
			if line[x] >= '0' && line[x] <= '9' {
				num, len := extractPartNumber(line[x:])

				part := partNumber{value: num, startX: x, endX: x + len - 1, y: y}
				parts = append(parts, part)

				x += len
			} else {
				if line[x] != '.' {
					if line[x] == '*' {
						gears[point{x, y}] = true
					} else {
						symbols[point{x, y}] = true
					}
				}

				x++
			}
		}
	}

	return parts, symbols, gears
}

func Part1(input string) int {
	parts, symbols, gears := parseSchematic(input)

	sum := 0

	for _, part := range parts {
	loop:
		for y := part.y - 1; y <= part.y+1; y++ {
			for x := part.startX - 1; x <= part.endX+1; x++ {
				if symbols[point{x, y}] || gears[point{x, y}] {
					sum += part.value
					break loop
				}
			}
		}
	}

	return sum
}

func Part2(input string) int {
	parts, _, gears := parseSchematic(input)

	partsByGear := make(map[point][]partNumber)

	for _, part := range parts {
	loop:
		for y := part.y - 1; y <= part.y+1; y++ {
			for x := part.startX - 1; x <= part.endX+1; x++ {
				if gears[point{x, y}] {
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

	return sum
}

//go:embed input.txt
var input string

func main() {
	input := util.Sanitize(input)

	fmt.Println("Part 1 =", Part1(input))
	fmt.Println("Part 2 =", Part2(input))
}
