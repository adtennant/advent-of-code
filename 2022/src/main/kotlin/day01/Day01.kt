package day01

import readInput
import split

fun main() {
    fun getTotals(input: List<String>) = input
        .split(String::isNotEmpty)
        .map {
            it.map(String::toInt)
                .sum()
        }

    fun part1(input: List<String>) = getTotals(input)
        .max()

    fun part2(input: List<String>) = getTotals(input)
        .sortedDescending()
        .take(3)
        .sum()

    val testInput = readInput("day01/Day01_test")
    check(part1(testInput) == 24000)
    check(part2(testInput) == 45000)

    val input = readInput("day01/Day01")
    println(part1(input))
    println(part2(input))
}
