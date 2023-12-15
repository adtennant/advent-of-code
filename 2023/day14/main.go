package main

import (
	_ "embed"

	"adtennant.dev/aoc/util"
)

type Grid [][]byte

func (g Grid) String() string {
	str := ""

	for _, row := range g {
		str += string(row)
	}

	return str
}

func parseGrid(input string) Grid {
	var grid [][]byte

	for _, line := range util.Lines(input) {
		grid = append(grid, []byte(line))
	}

	return grid
}

func tiltNorth(grid [][]byte) {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == 'O' {
				for i := y - 1; i >= 0 && grid[i][x] == '.'; i-- {
					grid[i+1][x], grid[i][x] = grid[i][x], grid[i+1][x]
				}
			}
		}
	}
}

func tiltWest(grid [][]byte) {
	for col := 0; col < len(grid[0]); col++ {
		for row, line := range grid {
			if line[col] == 'O' {
				for i := col - 1; i >= 0 && grid[row][i] == '.'; i-- {
					grid[row][i+1], grid[row][i] = grid[row][i], grid[row][i+1]
				}
			}
		}
	}

}

func tiltSouth(grid [][]byte) {
	for row := len(grid) - 1; row >= 0; row-- {
		for col, val := range grid[row] {
			if val == 'O' {
				for i := row + 1; i < len(grid) && grid[i][col] == '.'; i++ {
					grid[i-1][col], grid[i][col] = grid[i][col], grid[i-1][col]
				}
			}
		}
	}
}

func tiltEast(grid [][]byte) {
	for col := len(grid[0]) - 1; col >= 0; col-- {
		for row, line := range grid {
			if line[col] == 'O' {
				for i := col + 1; i < len(grid[0]) && grid[row][i] == '.'; i++ {
					grid[row][i-1], grid[row][i] = grid[row][i], grid[row][i-1]
				}
			}
		}
	}
}

func spinCycle(grid [][]byte) {
	tiltNorth(grid)
	tiltWest(grid)
	tiltSouth(grid)
	tiltEast(grid)
}

func countLoad(grid [][]byte) int {
	result := 0

	for y := len(grid) - 1; y >= 0; y-- {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == 'O' {
				result += len(grid) - y
			}
		}
	}

	return result
}

func Part1(input string) (int, error) {
	grid := parseGrid(input)
	tiltNorth(grid)

	return countLoad(grid), nil
}

const targetCycles = 1000000000

func Part2(input string) (int, error) {
	grid := parseGrid(input)
	states := make(map[string]int)

	for cycle := 0; cycle < targetCycles; cycle++ {
		spinCycle(grid)

		key := grid.String()

		if start, ok := states[key]; ok {
			len := cycle - start

			for j := 0; j < (targetCycles-cycle)%len-1; j++ {
				spinCycle(grid)
			}

			break
		}

		states[key] = cycle
	}

	return countLoad(grid), nil
}

//go:embed input.txt
var input string

func main() {
	util.Run(Part1, Part2, input)
}
