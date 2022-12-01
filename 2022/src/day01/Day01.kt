package day01

import readInput

fun main() {
    fun getTotals(input: List<String>): List<Int> {
        val elves = input.joinToString("\n").split("\n\n");
        return elves.map { it.split("\n").map(String::toInt).sum() }
    }

    fun part1(input: List<String>) = getTotals(input).max()

    fun part2(input: List<String>) = getTotals(input).sortedDescending().take(3).sum()

    val testInput = readInput("day01/Day01_test")
    check(part1(testInput) == 24000)
    check(part2(testInput) == 45000)

    val input = readInput("day01/Day01")
    println(part1(input))
    println(part2(input))
}
