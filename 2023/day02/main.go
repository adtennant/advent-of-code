package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"adtennant.dev/aoc/util"
)

type result struct {
	r int
	g int
	b int
}

type game struct {
	id      int
	results []result
}

var idRegexp = regexp.MustCompile(`\d+`)

func parseID(str string) int {
	v, _ := strconv.Atoi(idRegexp.FindString(str))
	return v
}

func parseColor(str, color string) int {
	re := regexp.MustCompile(fmt.Sprintf(`(\d*) %s`, color))
	matches := re.FindStringSubmatch(str)

	if matches == nil {
		return 0
	}

	v, _ := strconv.Atoi(matches[1])
	return v
}

func parseCubes(str string) result {
	return result{
		r: parseColor(str, "red"),
		g: parseColor(str, "green"),
		b: parseColor(str, "blue"),
	}
}

func parseResults(str string) (cubes []result) {
	for _, part := range strings.Split(str, ";") {
		cubes = append(cubes, parseCubes(part))
	}

	return cubes
}

func parseGame(line string) game {
	parts := strings.Split(line, ":")

	return game{
		id:      parseID(parts[0]),
		results: parseResults(parts[1]),
	}
}

func parseGames(input string) (games []game) {
	for _, line := range util.Lines(input) {
		games = append(games, parseGame(line))
	}

	return games
}

func Part1(input string) int {
	games := parseGames(input)

	maxR := 12
	maxG := 13
	maxB := 14

	sum := 0

	for _, game := range games {
		possible := true

		for _, result := range game.results {
			possible = result.r <= maxR && result.g <= maxG && result.b <= maxB

			if !possible {
				break
			}
		}

		if possible {
			sum += game.id
		}
	}

	return sum
}

func Part2(input string) int {
	games := parseGames(input)

	totalPower := 0

	for _, game := range games {
		minR := 0
		minG := 0
		minB := 0

		for _, result := range game.results {
			minR = max(minR, result.r)
			minG = max(minG, result.g)
			minB = max(minB, result.b)
		}

		totalPower += minR * minG * minB
	}

	return totalPower
}

//go:embed input.txt
var input string

func main() {
	input := util.Sanitize(input)

	fmt.Println("Part 1 =", Part1(input))
	fmt.Println("Part 2 =", Part2(input))
}
