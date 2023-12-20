package main

import (
	_ "embed"
	"strings"

	"adtennant.dev/aoc/util"
)

type pulse struct {
	src   string
	dest  string
	value bool
}

type module interface {
	Destinations() []string
	Handle(p pulse) *bool
}

type base struct {
	dests []string
}

func (m *base) Destinations() []string {
	return m.dests
}

type flipflop struct {
	base
	on bool
}

func (m *flipflop) Handle(p pulse) *bool {
	if p.value {
		return nil
	} else {
		m.on = !m.on
		return &m.on
	}
}

type conjunction struct {
	base
	memory map[string]bool
}

func (m *conjunction) Handle(p pulse) *bool {
	m.memory[p.src] = p.value

	sent := false

	for _, value := range m.memory {
		if !value {
			sent = true
			break
		}
	}

	return &sent
}

type broadcaster struct {
	base
}

func (m *broadcaster) Handle(p pulse) *bool {
	return &p.value
}

func send(modules map[string]module, p pulse) []pulse {
	module := modules[p.dest]

	if module == nil {
		return nil
	}

	output := module.Handle(p)

	if output == nil {
		return nil
	}

	var sent []pulse

	for _, t := range module.Destinations() {
		sent = append(sent, pulse{
			src:   p.dest,
			dest:  t,
			value: *output,
		})
	}

	return sent
}

func parseModule(str string) (string, module) {
	parts := strings.Split(str, " -> ")
	targets := strings.Split(parts[1], ", ")

	switch {
	case parts[0][0] == '%':
		return parts[0][1:], &flipflop{base{targets}, false}
	case parts[0][0] == '&':
		return parts[0][1:], &conjunction{base{targets}, make(map[string]bool)}
	case parts[0] == "broadcaster":
		return parts[0], &broadcaster{base{targets}}
	default:
		panic("")
	}
}

func parseModules(input string) map[string]module {
	modules := make(map[string]module)

	for _, line := range util.Lines(input) {
		name, module := parseModule(line)
		modules[name] = module
	}

	for name, module := range modules {
		for _, d := range module.Destinations() {
			dest := modules[d]

			if m, ok := dest.(*conjunction); ok {
				m.memory[name] = false
			}
		}
	}

	return modules
}

func Part1(input string) (int, error) {
	modules := parseModules(input)

	low := 0
	high := 0

	for i := 0; i < 1000; i++ {
		q := util.NewQueue[pulse]()
		q.Push(pulse{"button", "broadcaster", false})

		for q.Len() > 0 {
			pulse := q.Pop()

			if pulse.value {
				high += 1
			} else {
				low += 1
			}

			sent := send(modules, pulse)
			q.Push(sent...)
		}
	}

	return low * high, nil
}

func findInputs(modules map[string]module, module string) []string {
	var inputs []string

	for name, m := range modules {
		for _, dest := range m.Destinations() {
			if dest == module {
				inputs = append(inputs, name)
			}
		}
	}

	return inputs
}

func Part2(input string) (int, error) {
	modules := parseModules(input)

	rxInput := findInputs(modules, "rx")[0] // bh
	rxInputInputs := findInputs(modules, rxInput)

	foundCycles := make(map[string]int)
	i := 1

	for len(foundCycles) != len(rxInputInputs) {
		q := util.NewQueue[pulse]()
		q.Push(pulse{"button", "broadcaster", false})

		for q.Len() > 0 {
			pulse := q.Pop()

			sent := send(modules, pulse)
			q.Push(sent...)

			for _, name := range rxInputInputs {
				_, ok := foundCycles[name]

				if !ok && modules[rxInput].(*conjunction).memory[name] {
					foundCycles[name] = i
				}
			}
		}

		i++
	}

	var cycles []int

	for _, cycle := range foundCycles {
		cycles = append(cycles, cycle)
	}

	return util.LCM(cycles...), nil
}

//go:embed input.txt
var input string

func main() {
	util.Run(Part1, Part2, input)
}
