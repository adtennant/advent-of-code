package dev.adtennant.aoc.`2025`.day.`03`

import org.scalatest.funsuite.AnyFunSuite

val input =
  """987654321111111
811111111111119
234234234234278
818181911112111"""

class Day03Tests extends AnyFunSuite {
  test("Part1") {
    val result = Day03.solvePart1(input)
    assert(result == 357)
  }

  test("Part2") {
    val result = Day03.solvePart2(input)
    assert(result == 3121910778619L)
  }
}
