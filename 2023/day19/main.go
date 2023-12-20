package main

import (
	_ "embed"
	"regexp"
	"strconv"
	"strings"

	"adtennant.dev/aoc/util"
)

const (
	LT = '<'
	GT = '>'
)

type rule struct {
	field    byte
	operator byte
	value    int
	next     string
}

var ruleRegexp = regexp.MustCompile("([xmas])([<>])([0-9]+):([a-zAR]+)")

func parseRule(str string) rule {
	matches := ruleRegexp.FindStringSubmatch(str)
	value, _ := strconv.Atoi(matches[3])

	return rule{
		field:    matches[1][0],
		operator: matches[2][0],
		value:    value,
		next:     matches[4],
	}
}

type workflow struct {
	name      string
	rules     []rule
	otherwise string
}

func parseWorkflow(str string) workflow {
	parts := strings.Split(str, "{")

	name := parts[0]
	parts = strings.Split(strings.TrimSuffix(parts[1], "}"), ",")

	var rules []rule

	for _, part := range parts[:len(parts)-1] {
		rule := parseRule(part)
		rules = append(rules, rule)
	}

	otherwise := parts[len(parts)-1]

	return workflow{name, rules, otherwise}
}

type part struct {
	x, m, a, s int
}

var partRegexp = regexp.MustCompile("{x=([0-9]+),m=([0-9]+),a=([0-9]+),s=([0-9]+)}")

func parsePart(str string) part {
	matches := partRegexp.FindStringSubmatch(str)

	x, _ := strconv.Atoi(matches[1])
	m, _ := strconv.Atoi(matches[2])
	a, _ := strconv.Atoi(matches[3])
	s, _ := strconv.Atoi(matches[4])

	return part{x, m, a, s}
}

func parseInput(input string) (map[string]workflow, []part) {
	lines := util.Lines(input)
	i := 0

	workflows := make(map[string]workflow)

	for ; i < len(lines); i++ {
		if lines[i] == "" {
			break
		}

		workflow := parseWorkflow(lines[i])
		workflows[workflow.name] = workflow
	}

	i++

	var parts []part

	for ; i < len(lines); i++ {
		part := parsePart(lines[i])
		parts = append(parts, part)
	}

	return workflows, parts
}

func compare(r rule, p part) (string, bool) {
	var value int

	switch r.field {
	case 'x':
		value = p.x
	case 'm':
		value = p.m
	case 'a':
		value = p.a
	case 's':
		value = p.s
	}

	switch r.operator {
	case LT:
		return r.next, value < r.value
	case GT:
		return r.next, value > r.value
	}

	return "", false
}

func evaluate(w workflow, p part) string {
	for _, r := range w.rules {
		dest, ok := compare(r, p)

		if ok {
			return dest
		}
	}

	return w.otherwise
}

func Part1(input string) (int, error) {
	workflows, parts := parseInput(input)

	sum := 0

	for _, p := range parts {
		current := "in"

		for current != "A" && current != "R" {
			current = evaluate(workflows[current], p)
		}

		if current == "A" {
			sum += p.x + p.m + p.a + p.s
		}
	}

	return sum, nil
}

var rangeIndexes = map[byte]int{
	'x': 0,
	'm': 1,
	'a': 2,
	's': 3,
}

func findCombinations(workflows map[string]workflow, current string, ranges [4][2]int) int {
	if current == "A" {
		sum := 1

		for _, r := range ranges {
			sum *= r[1] - r[0] + 1
		}

		return sum
	}

	if current == "R" {
		return 0
	}

	count := 0
	workflow := workflows[current]

	for _, r := range workflow.rules {
		var nextRanges [4][2]int
		copy(nextRanges[:], ranges[:])

		i := rangeIndexes[r.field]

		if r.operator == LT {
			ranges[i][0] = r.value

			nextRanges[i][1] = r.value - 1
			count += findCombinations(workflows, r.next, nextRanges)
		} else {
			ranges[i][1] = r.value

			nextRanges[i][0] = r.value + 1
			count += findCombinations(workflows, r.next, nextRanges)
		}
	}

	count += findCombinations(workflows, workflow.otherwise, ranges)

	return count
}

func Part2(input string) (int, error) {
	workflows, _ := parseInput(input)

	return findCombinations(workflows, "in", [4][2]int{{1, 4000}, {1, 4000}, {1, 4000}, {1, 4000}}), nil
}

//go:embed input.txt
var input string

func main() {
	util.Run(Part1, Part2, input)
}
