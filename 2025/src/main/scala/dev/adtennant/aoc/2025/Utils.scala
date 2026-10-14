package dev.adtennant.aoc.`2025`

import scala.io.Source._
import java.time.Instant
import java.time.Duration

type Result[Output] = (Output, Double)

def benchmark[Output](f: () => Output): Result[Output] =
  val start = Instant.now
  val result = f()
  val end = Instant.now

  val diff = Duration.between(start, end)
  (result, diff.toMillis.toDouble / 1000.0)

def printHeader(day: Int): Unit =
  println(s"# Day ${day}")

def printResult[Output](part: Int, result: Result[Output]): Unit =
  println(s"Part $part = ${result._1} (in ${result._2}s)")

trait Solution[Output]:
  def day: Int

  def solvePart1(input: String): Output
  def solvePart2(input: String): Output

  def run: Unit =
    printHeader(day)

    val input = fromResource(f"day$day%02d.txt").mkString

    val result1 = benchmark(() => solvePart1(input))
    printResult(1, result1)

    val result2 = benchmark(() => solvePart2(input))
    printResult(2, result2)

    println()
