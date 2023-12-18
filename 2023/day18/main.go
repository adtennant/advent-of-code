package main

import (
	_ "embed"
	"regexp"
	"strconv"

	"adtennant.dev/aoc/util"
)

type entry struct {
	dir      util.Direction
	distance int64
}

var entryRegexp = regexp.MustCompile(`([UDLR]) ([0-9]+).*`)

func parsePlan(input string) (plan []entry, err error) {
	for _, line := range util.Lines(input) {
		matches := entryRegexp.FindStringSubmatch(line)
		distance, _ := strconv.Atoi(matches[2])

		plan = append(plan, entry{
			dir:      util.Direction(matches[1][0]),
			distance: int64(distance),
		})
	}

	return plan, nil
}

var hexEntryRegexp = regexp.MustCompile(`.*\(#(.*)\)`)

func parsePlanFromHex(input string) (plan []entry, err error) {
	for _, line := range util.Lines(input) {
		matches := hexEntryRegexp.FindStringSubmatch(line)
		distance, _ := strconv.ParseInt(matches[1][:5], 16, 64)

		var dir util.Direction

		switch matches[1][5] {
		case '0':
			dir = util.RIGHT
		case '1':
			dir = util.DOWN
		case '2':
			dir = util.LEFT
		case '3':
			dir = util.UP
		}

		plan = append(plan, entry{
			dir:      dir,
			distance: int64(distance),
		})
	}

	return plan, nil
}

type point = util.Point[int64]

func solve(plan []entry) int64 {
	start := point{X: 0, Y: 0}

	a := int64(0)

	for _, e := range plan {
		delta := util.GetDelta[int64](e.dir)
		end := start.Translate(delta.Scale(e.distance))

		a += (start.Y + end.Y) * (start.X - end.X)
		a += e.distance

		start = end
	}

	return (a / 2) + 1
}

func Part1(input string) (int64, error) {
	plan, _ := parsePlan(input)
	return solve(plan), nil
}

func Part2(input string) (int64, error) {
	plan, _ := parsePlanFromHex(input)
	return solve(plan), nil
}

//go:embed input.txt
var input string

func main() {
	util.Run(Part1, Part2, input)
}
