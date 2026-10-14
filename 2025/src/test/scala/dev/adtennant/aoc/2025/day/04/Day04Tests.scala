package dev.adtennant.aoc.`2025`.day.`04`

import org.scalatest.funsuite.AnyFunSuite

val input =
  """..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@."""

class Day04Tests extends AnyFunSuite {
  test("Part1") {
    val result = Day04.solvePart1(input)
    assert(result == 13)
  }

  test("Part2") {
    val result = Day04.solvePart2(input)
    assert(result == 43)
  }
}
