package dev.adtennant.aoc.`2025`.day.`03`

import scala.io.Source._
import dev.adtennant.aoc.`2025`.Solution
import scala.util.matching.Regex
import scala.annotation.tailrec

case class Bank(batteries: Seq[Int]):
  def joltage(length: Int, result: Long = 0): Long = {
    if length > 0 then {
      val digit =
        batteries
          .dropRight(length - 1)
          .max

      Bank(
        batteries
          .dropWhile(_ != digit)
          .drop(1)
      ).joltage(
        length - 1,
        result * 10 + digit
      )
    } else result
  }

def parseBanks(input: String): Seq[Bank] =
  input.linesIterator
    .map(
      _.split("")
        .map(_.toInt)
        .toSeq
    )
    .map(Bank(_))
    .toSeq

object Day03 extends Solution[Long]:
  def day = 3

  def solvePart1(input: String): Long =
    val banks = parseBanks(input)
    banks.map(_.joltage(2)).sum

  def solvePart2(input: String): Long =
    val banks = parseBanks(input)
    banks.map(_.joltage(12)).sum

@main
def main: Unit =
  Day03.run
