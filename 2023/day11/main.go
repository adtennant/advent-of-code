package main

import (
	_ "embed"
	"math"

	"adtennant.dev/aoc/util"
)

type point = util.Point[int64]

type image struct {
	galaxies  []point
	emptyRows map[int64]bool
	emptyCols map[int64]bool
}

func parseImage(input string) image {
	lines := util.Lines(input)

	var galaxies []point

	for y, line := range lines {
		for x, c := range []byte(line) {
			if c == '#' {
				galaxies = append(galaxies, point{int64(x), int64(y)})
			}
		}
	}

	emptyRows := make(map[int64]bool)

	for y, line := range lines {
		empty := true

		for _, c := range []byte(line) {
			if c == '#' {
				empty = false
				break
			}
		}

		if empty {
			emptyRows[int64(y)] = true
		}
	}

	emptyCols := make(map[int64]bool)

	for x := 0; x < len(lines[0]); x++ {
		empty := true

		for y := 0; y < len(lines); y++ {
			if lines[y][x] == '#' {
				empty = false
				break
			}
		}

		if empty {
			emptyCols[int64(x)] = true
		}
	}

	return image{galaxies, emptyRows, emptyCols}
}

func findDistance(a, b point, emptyRows map[int64]bool, emptyCols map[int64]bool, expansion int64) int64 {
	dx := int64(math.Abs(float64(a.X - b.X)))
	dy := int64(math.Abs(float64(a.Y - b.Y)))

	for x := min(a.X, b.X); x < max(a.X, b.X); x++ {
		if emptyCols[x] {
			dx += int64(expansion)
		}
	}

	for y := min(a.Y, b.Y); y < max(a.Y, b.Y); y++ {
		if emptyRows[y] {
			dy += int64(expansion)
		}
	}

	return dx + dy
}

func findTotalDistance(image image, expansion int64) int64 {
	total := int64(0)

	for i, g1 := range image.galaxies {
		for j := i + 1; j < len(image.galaxies); j++ {
			g2 := image.galaxies[j]

			total += findDistance(g1, g2, image.emptyRows, image.emptyCols, expansion)
		}
	}

	return total
}

func Part1(input string) (int64, error) {
	image := parseImage(input)

	return findTotalDistance(image, 1), nil
}

func Part2(input string) (int64, error) {
	image := parseImage(input)

	return findTotalDistance(image, 999999), nil
}

//go:embed input.txt
var input string

func main() {
	util.Run(Part1, Part2, input)
}
