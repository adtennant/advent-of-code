package dev.adtennant.adventofcode.day03

import readInput
import takeExact

val Char.priority
    get() = if (isLowerCase()) {
        code - 96
    } else {
        code - 38
    }

class Rucksack(private val contents: CharArray) {
    val commonItem
        get(): Char {
            val half = contents.size / 2
            val first = contents.take(half)
            val second = contents.takeLast(half)

            return first.intersect(second.toSet())
                .takeExact(1)
                .first()
        }
}

class RucksackGroup(private val first: CharArray, private val second: CharArray, private val third: CharArray) {
    val commonItem
        get() = first.intersect(second.toSet())
            .intersect(third.toSet())
            .takeExact(1)
            .first()
}

fun main() {
    fun part1(input: List<String>) = input.map(String::toCharArray)
        .map { Rucksack(it) }
        .sumOf { it.commonItem.priority }

    fun part2(input: List<String>) = input.map(String::toCharArray)
        .windowed(3, 3)
        .map { RucksackGroup(it[0], it[1], it[2]) }
        .sumOf { it.commonItem.priority }

    val testInput = readInput("day03/Day03_test")
    check(part1(testInput) == 157)
    check(part2(testInput) == 70)

    val input = readInput("day03/Day03")
    println(part1(input))
    println(part2(input))
}
