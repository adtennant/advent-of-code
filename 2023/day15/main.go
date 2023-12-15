package main

import (
	_ "embed"
	"regexp"
	"slices"
	"strings"

	"adtennant.dev/aoc/util"
)

func hash(str string) int {
	hash := 0

	for _, c := range []byte(str) {
		hash += int(c)
		hash *= 17
		hash %= 256
	}

	return hash
}

func Part1(input string) (int, error) {
	sum := 0

	for _, step := range strings.Split(input, ",") {
		sum += hash(step)
	}

	return sum, nil
}

type step struct {
	label string
	op    byte
	focus int
}

var stepRegexp = regexp.MustCompile("([A-Za-z]*)([=-])([0-9]*)")

func parseSteps(input string) []step {
	var steps []step

	for _, s := range strings.Split(input, ",") {
		matches := stepRegexp.FindStringSubmatch(s)
		focus := -1

		if matches[2][0] == '=' {
			focus = int(matches[3][0] - '0')
		}

		steps = append(steps, step{
			matches[1],
			matches[2][0],
			focus,
		})
	}

	return steps
}

type lens struct {
	label string
	focus int
}

func apply(boxes map[int][]lens, s step) {
	box := hash(s.label)
	findLens := func(l lens) bool {
		return l.label == s.label
	}

	switch s.op {
	case '-':
		boxes[box] = slices.DeleteFunc(boxes[box], findLens)
	case '=':
		i := slices.IndexFunc(boxes[box], findLens)

		if i > -1 {
			boxes[box][i] = lens{s.label, s.focus}
		} else {
			boxes[box] = append(boxes[box], lens{s.label, s.focus})
		}
	}
}

func Part2(input string) (int, error) {
	steps := parseSteps(input)
	boxes := make(map[int][]lens)

	for _, step := range steps {
		apply(boxes, step)
	}

	sum := 0

	for i, box := range boxes {
		for slot, lens := range box {
			sum += (i + 1) * (slot + 1) * lens.focus
		}
	}

	return sum, nil
}

//go:embed input.txt
var input string

func main() {
	util.Run(Part1, Part2, input)
}
