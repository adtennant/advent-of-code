package dev.adtennant.aoc.`2025`.day.`02`

import scala.io.Source._
import dev.adtennant.aoc.`2025`.Solution
import scala.util.matching.Regex

case class ProductID(value: Long):
  def toLong: Long = value
  override def toString: String = value.toString

case class Range(start: ProductID, end: ProductID):
  def toSeq = (start.toLong to end.toLong).map(ProductID(_))

def parseRanges(input: String): Seq[Range] =
  input
    .split(",")
    .collect { case s"$start-$end" =>
      Range(ProductID(start.toLong), ProductID(end.toLong))
    }
    .toSeq

object Day02 extends Solution[Long]:
  def day = 2

  def solvePart1(input: String): Long =
    def isValid(id: ProductID): Boolean =
      val str = id.toString
      val start = str.take(str.length / 2)
      val end = str.drop(str.length / 2)
      start != end

    val ranges = parseRanges(input)
    ranges
      .flatMap(_.toSeq)
      .filter(!isValid(_))
      .map(_.toLong)
      .sum

  def solvePart2(input: String): Long =
    def isValid(id: ProductID): Boolean =
      val str = id.toString
      (1 to str.length / 2).forall(n => str != str.take(n) * (str.length / n))

    val ranges = parseRanges(input)
    ranges
      .flatMap(_.toSeq)
      .filter(!isValid(_))
      .map(_.toLong)
      .sum

@main
def main: Unit =
  Day02.run
