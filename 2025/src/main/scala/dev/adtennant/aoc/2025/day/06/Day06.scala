package dev.adtennant.aoc.`2025`.day.`06`

import scala.io.Source._
import dev.adtennant.aoc.`2025`.Solution
import scala.util.matching.Regex
import scala.annotation.tailrec
import scala.collection.immutable.NumericRange.Inclusive

enum Operation:
  case Add, Multiply

case class Problem(nums: Seq[Long], op: Operation):
  def evaluate: Long =
    op match
      case Operation.Add      => nums.sum
      case Operation.Multiply => nums.product

def parseProblems(input: String): Seq[Problem] =
  val nums = input.linesIterator
    .take(input.linesIterator.size - 1)
    .map(_.trim.split("\\s+").map(_.toLong))
    .toSeq

  val ops = input.linesIterator
    .drop(input.linesIterator.size - 1)
    .next()
    .trim
    .split("\\s+")
    .map {
      case "+" => Operation.Add
      case "*" => Operation.Multiply
    }

  nums(0).zipWithIndex
    .map((num, i) => Problem(nums.map(_(i)), ops(i)))
    .toSeq

def split[T](elems: Seq[T], p: T => Boolean): Seq[Seq[T]] =
  elems match {
    case Seq()                            => Seq()
    case Seq(head, tail @ _*) if !p(head) => Seq(head) +: split(tail, p)
    case xs =>
      val (prefix, rest) = xs span p
      prefix +: split(rest, p)
  }

def parseProblems2(input: String): Seq[Seq[String]] =
  split(
    input.linesIterator.toSeq.transpose.map(_.mkString),
    _.trim.isEmpty
  )

object Day06 extends Solution[Long]:
  def day = 6

  def solvePart1(input: String): Long =
    val problems = parseProblems(input)
    problems.map(_.evaluate).sum

  def solvePart2(input: String): Long =
    val problems2 = parseProblems2(input)
    problems2.foreach(println)
    0

@main
def main: Unit =
  Day06.run
