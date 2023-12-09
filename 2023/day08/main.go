package main

import (
	_ "embed"
	"fmt"
	"strings"

	"adtennant.dev/aoc/util"
)

type node struct {
	left  string
	right string
}

func parseNode(str string) (string, node, error) {
	parts := strings.Split(str, " = ")
	if len(parts) != 2 {
		return "", node{}, fmt.Errorf("invalid node format")
	}

	name := parts[0]

	parts = strings.Split(parts[1], ", ")
	if len(parts) != 2 {
		return "", node{}, fmt.Errorf("invalid node format")
	}

	left := strings.TrimPrefix(parts[0], "(")
	right := strings.TrimSuffix(parts[1], ")")

	return name, node{left, right}, nil
}

func parse(input string) ([]byte, map[string]node, error) {
	lines := util.Lines(input)

	instrs := []byte(lines[0])
	network := make(map[string]node)

	for i := 2; i < len(lines); i++ {
		name, node, err := parseNode(lines[i])
		if err != nil {
			return nil, nil, fmt.Errorf("failed to parse node: %w", err)
		}

		network[name] = node
	}

	return instrs, network, nil
}

func countSteps(instrs []byte, network map[string]node, start string, done func(string) bool) int {
	current := start
	steps := 0

	for !done(current) {
		instr := instrs[steps%len(instrs)]

		switch instr {
		case 'L':
			current = network[current].left
		case 'R':
			current = network[current].right
		}

		steps++
	}

	return steps
}

func Part1(input string) (int, error) {
	instrs, network, err := parse(input)
	if err != nil {
		return -1, err
	}

	return countSteps(instrs, network, "AAA", func(current string) bool {
		return current == "ZZZ"
	}), nil
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

func lcm(x ...int) int {
	if len(x) == 1 {
		return x[0]
	} else if len(x) > 2 {
		return lcm(x[0], lcm(x[1:]...))
	}

	return x[0] * x[1] / gcd(x[0], x[1])
}

func Part2(input string) (int, error) {
	instrs, network, err := parse(input)
	if err != nil {
		return -1, err
	}

	var ghosts []string

	for n := range network {
		if strings.HasSuffix(n, "A") {
			ghosts = append(ghosts, n)
		}
	}

	var shortest []int

	for _, ghost := range ghosts {
		steps := countSteps(instrs, network, ghost, func(current string) bool {
			return strings.HasSuffix(current, "Z")
		})

		shortest = append(shortest, steps)
	}

	return lcm(shortest...), nil
}

//go:embed input.txt
var input string

func main() {
	util.Run(Part1, Part2, input)
}
