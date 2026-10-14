package dev.adtennant.aoc.`2025`.day.`05`

import scala.io.Source._
import dev.adtennant.aoc.`2025`.Solution
import scala.util.matching.Regex
import scala.annotation.tailrec
import scala.collection.immutable.NumericRange.Inclusive

case class Range(from: Long, to: Long):
  def contains(value: Long): Boolean = value >= from && value <= to

def parseRangesAndIds(input: String): (Seq[Range], Seq[Long]) =
  val Array(rulesS, updatesS) = input.split("\n\n")

  val rules = rulesS.linesIterator
    .map(_.split('-').map(_.toLong))
    .map { case Array(left, right) => Range(left, right) }
    .toSeq

  val updates = updatesS.linesIterator
    .map(_.toLong)
    .toSeq

  (rules, updates)

def isFresh(ingredient: Long, ranges: Seq[Range]): Boolean =
  ranges.exists(range => range.contains(ingredient))

object Day05 extends Solution[Long]:
  def day = 5

  def solvePart1(input: String): Long =
    val (ranges, updates) = parseRangesAndIds(input)
    updates.count(ingredient => isFresh(ingredient, ranges))

  def solvePart2(input: String): Long =
    val (ranges, _) = parseRangesAndIds(input)
    ranges.toIterable
      .sortBy(_.from)
      .foldLeft(Seq.empty[Range]) { (merged, range) =>
        val Range(from, to) = range

        merged.lastOption match {
          case Some(Range(previousFrom, previousTo)) if from <= previousTo =>
            merged.dropRight(1) :+ Range(previousFrom, Math.max(previousTo, to))
          case _ =>
            merged :+ range
        }
      }
      .map(range => range.to - range.from + 1)
      .sum

@main
def main: Unit =
  Day05.run
