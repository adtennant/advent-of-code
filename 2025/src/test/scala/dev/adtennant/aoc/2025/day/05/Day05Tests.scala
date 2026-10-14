package dev.adtennant.aoc.`2025`.day.`05`

import org.scalatest.funsuite.AnyFunSuite

val input =
  """3-5
10-14
16-20
12-18

1
5
8
11
17
32"""

class Day05Tests extends AnyFunSuite {
  test("Part1") {
    val result = Day05.solvePart1(input)
    assert(result == 3)
  }

  test("Part2") {
    val result = Day05.solvePart2(input)
    assert(result == 0)
  }
}
