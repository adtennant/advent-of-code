package dev.adtennant.aoc.`2025`

import dev.adtennant.aoc.`2025`.Solution

import dev.adtennant.aoc.`2025`.day.`01`.Day01
import dev.adtennant.aoc.`2025`.day.`02`.Day02
import dev.adtennant.aoc.`2025`.day.`03`.Day03
import dev.adtennant.aoc.`2025`.day.`04`.Day04
import dev.adtennant.aoc.`2025`.day.`05`.Day05

val solutions: Seq[Solution[?]] =
  Seq(Day01, Day02, Day03, Day04, Day05)

@main
def main(): Unit =
  solutions.foreach(_.run)
