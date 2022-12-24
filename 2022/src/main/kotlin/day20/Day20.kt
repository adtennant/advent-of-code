package day20

import readInput

// Wrap in a class for equality checks to match instances not values
class Number(val value: Long)

fun String.toNumber(key: Long) = Number(toLong() * key)

fun mix(numbers: List<Number>, times: Int) = (0 until times)
    .fold(numbers.toMutableList()) { acc, _ ->
        numbers.fold(acc) { mixed, num ->
            val currentIndex = mixed.indexOf(num)
            val newIndex = (currentIndex + num.value).mod(numbers.size - 1)
            mixed.removeAt(currentIndex)
            mixed.add(newIndex, num)

            mixed
        }
    }

fun solve(input: List<String>, times: Int = 1, key: Long = 1): Long {
    val numbers = input.map { it.toNumber(key) }
    val mixed = mix(numbers, times)

    val firstZero = mixed.indexOfFirst { it.value == 0L }
    return listOf(1000, 2000, 3000)
        .sumOf { mixed[(firstZero + it) % mixed.size].value }
}

fun main() {
    fun part1(input: List<String>) = solve(input)

    fun part2(input: List<String>) = solve(input, 10, 811589153)

    val testInput = readInput("day20/Day20_test")
    check(part1(testInput) == 3L)
    check(part2(testInput) == 1623178306L)

    val input = readInput("day20/Day20")
    println(part1(input))
    println(part2(input))
}
