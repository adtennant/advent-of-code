package main

import (
	_ "embed"
	"strings"

	"adtennant.dev/aoc/util"
)

type node struct {
	left  string
	right string
}

func parseNode(str string) (string, node, error) {
	parts := strings.Split(str, " = ")

	name := parts[0]

	parts = strings.Split(parts[1], ", ")
	left := strings.TrimPrefix(parts[0], "(")
	right := strings.TrimSuffix(parts[1], ")")

	return name, node{left, right}, nil
}

func parse(input string) ([]byte, map[string]node) {
	lines := util.Lines(input)

	instrs := []byte(lines[0])

	network := make(map[string]node)

	for i := 2; i < len(lines); i++ {
		name, node, _ := parseNode(lines[i])

		network[name] = node
	}

	return instrs, network
}

func Part1(input string) (int, error) {
	instrs, network := parse(input)

	current := "AAA"
	steps := 0

	for current != "ZZZ" {
		instr := instrs[steps%len(instrs)]

		switch instr {
		case 'L':
			current = network[current].left
		case 'R':
			current = network[current].right
		}

		steps++
	}

	return steps, nil
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
	instrs, network := parse(input)

	var ghosts []string

	for n := range network {
		if strings.HasSuffix(n, "A") {
			ghosts = append(ghosts, n)
		}
	}

	var shortest []int

	for _, ghost := range ghosts {
		steps := 0

		for !strings.HasSuffix(ghost, "Z") {
			instr := instrs[steps%len(instrs)]

			switch instr {
			case 'L':
				ghost = network[ghost].left
			case 'R':
				ghost = network[ghost].right
			}

			steps++
		}

		shortest = append(shortest, steps)
	}

	return lcm(shortest...), nil
}

//go:embed input.txt
var input string

func main() {
	util.Run(Part1, Part2, input)
}
