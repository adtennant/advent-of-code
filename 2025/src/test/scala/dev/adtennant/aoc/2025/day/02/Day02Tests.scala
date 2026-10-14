package dev.adtennant.aoc.`2025`.day.`02`

import org.scalatest.funsuite.AnyFunSuite

val input =
  """11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"""

class Day02Tests extends AnyFunSuite {
  test("Part1") {
    val result = Day02.solvePart1(input)
    assert(result == 1227775554L)
  }

  test("Part2") {
    val result = Day02.solvePart2(input)
    assert(result == 4174379265L)
  }
}
