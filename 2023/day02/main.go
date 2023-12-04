package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"unicode"

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

func parseID(str string) (int, error) {
	v, err := strconv.Atoi(strings.TrimLeftFunc(str, func(r rune) bool {
		return !unicode.IsNumber(r)
	}))
	if err != nil {
		return -1, err
	}

	return v, nil
}

func parseResult(str string) (result, error) {
	colors := strings.Split(str, ",")
	res := result{}

	for _, color := range colors {
		parts := strings.Split(strings.TrimSpace(color), " ")
		if len(parts) != 2 {
			return res, fmt.Errorf("invalid format")
		}

		num, err := strconv.Atoi(parts[0])
		if err != nil {
			return res, fmt.Errorf("invalid num: %s", parts[0])
		}

		switch parts[1] {
		case "red":
			res.r = num
		case "green":
			res.g = num
		case "blue":
			res.b = num
		default:
			return res, fmt.Errorf("invalid color: %s", parts[1])
		}
	}

	return res, nil
}

func parseResults(str string) (results []result, err error) {
	for _, part := range strings.Split(str, ";") {
		result, err := parseResult(part)
		if err != nil {
			return nil, fmt.Errorf("failed to parse result: %w", err)
		}

		results = append(results, result)
	}

	return results, nil
}

func parseGame(line string) (game, error) {
	parts := strings.Split(line, ":")
	if len(parts) != 2 {
		return game{}, fmt.Errorf("invalid format")
	}

	id, err := parseID(parts[0])
	if err != nil {
		return game{}, fmt.Errorf("failed to parse ID: %w", err)
	}

	results, err := parseResults(parts[1])
	if err != nil {
		return game{}, fmt.Errorf("failed to parse result: %w", err)
	}

	return game{
		id,
		results,
	}, nil
}

func parseGames(input string) (games []game, err error) {
	for _, line := range util.Lines(input) {
		game, err := parseGame(line)
		if err != nil {
			return nil, fmt.Errorf("failed to parse game from line: %s: %w", line, err)
		}

		games = append(games, game)
	}

	return games, nil
}

func Part1(input string) (int, error) {
	games, err := parseGames(input)
	if err != nil {
		return -1, err
	}

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

	return sum, nil
}

func Part2(input string) (int, error) {
	games, err := parseGames(input)
	if err != nil {
		return -1, err
	}

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

	return totalPower, nil
}

//go:embed input.txt
var input string

func main() {
	util.Main(input, Part1, Part2)
}
