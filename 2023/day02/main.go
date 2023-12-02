package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"adtennant.dev/aoc/util"
)

type Cubes struct {
	r int
	g int
	b int
}

type Game struct {
	id    int
	cubes []Cubes
}

var idRegexp = regexp.MustCompile(`\d+`)

func ParseID(str string) int {
	v, _ := strconv.Atoi(idRegexp.FindString(str))
	return v
}

func ParseColor(str, color string) int {
	re := regexp.MustCompile(fmt.Sprintf(`(\d*) %s`, color))
	matches := re.FindStringSubmatch(str)

	if matches == nil {
		return 0
	}

	v, _ := strconv.Atoi(matches[1])
	return v
}

func ParseCubes(str string) Cubes {
	return Cubes{
		r: ParseColor(str, "red"),
		g: ParseColor(str, "green"),
		b: ParseColor(str, "blue"),
	}
}

func ParseResults(str string) (cubes []Cubes) {
	for _, part := range strings.Split(str, ";") {
		cubes = append(cubes, ParseCubes(part))
	}

	return cubes
}

func ParseGame(line string) Game {
	parts := strings.Split(line, ":")

	return Game{
		id:    ParseID(parts[0]),
		cubes: ParseResults(parts[1]),
	}
}

func ParseGames(input string) (games []Game) {
	for _, line := range util.Lines(input) {
		games = append(games, ParseGame(line))
	}

	return games
}

func Part1(input string) int {
	games := ParseGames(input)

	maxR := 12
	maxG := 13
	maxB := 14

	sum := 0

	for _, game := range games {
		possible := true

		for _, result := range game.cubes {
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
	games := ParseGames(input)

	totalPower := 0

	for _, game := range games {
		minR := 0
		minG := 0
		minB := 0

		for _, result := range game.cubes {
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
