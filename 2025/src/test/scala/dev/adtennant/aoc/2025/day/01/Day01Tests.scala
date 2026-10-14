package dev.adtennant.aoc.`2025`.day.`01`

import org.scalatest.funsuite.AnyFunSuite

val input = """L68
L30
R48
L5
R60
L55
L1
L99
R14
L82"""

class Day01Tests extends AnyFunSuite {
  test("Part1") {
    val result = Day01.solvePart1(input)
    assert(result == 3)
  }

  test("Part2") {
    val result = Day01.solvePart2(input)
    assert(result == 6)
  }
}
