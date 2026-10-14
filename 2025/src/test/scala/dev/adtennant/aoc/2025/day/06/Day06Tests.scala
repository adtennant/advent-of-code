package dev.adtennant.aoc.`2025`.day.`06`

import org.scalatest.funsuite.AnyFunSuite

val input =
  """123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  """

class Day06Tests extends AnyFunSuite {
  test("Part1") {
    val result = Day06.solvePart1(input)
    assert(result == 4277556)
  }

  test("Part2") {
    val result = Day06.solvePart2(input)
    assert(result == 0)
  }
}
