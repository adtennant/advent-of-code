package dev.adtennant.adventofcode.day04

import readInput
import takeExact

class Assignment(private val first: IntRange, private val second: IntRange) {
    val hasFullOverlap
        get(): Boolean {
            return first.all { second.contains(it) }
                || second.all { first.contains(it) }
        }

    val hasPartialOverlap
        get(): Boolean {
            return first.any { second.contains(it) }
                || second.any { first.contains(it) }
        }
}

fun String.toIntRange() = split("-")
    .map(String::toInt)
    .windowed(2, 2)
    .takeExact(1)
    .map { IntRange(it[0], it[1]) }
    .first()

fun String.toAssignment() = split(",")
    .map(String::toIntRange)
    .windowed(2, 2)
    .takeExact(1)
    .map { Assignment(it[0], it[1]) }
    .first()

fun main() {
    fun part1(input: List<String>) = input.map { it.toAssignment() }
        .count { it.hasFullOverlap }

    fun part2(input: List<String>) = input.map { it.toAssignment() }
        .count { it.hasPartialOverlap }

    val testInput = readInput("day04/Day04_test")
    check(part1(testInput) == 2)
    check(part2(testInput) == 4)

    val input = readInput("day04/Day04")
    println(part1(input))
    println(part2(input))
}