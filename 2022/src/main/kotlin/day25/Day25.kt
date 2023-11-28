package day25

import readInput

val from = mapOf('=' to -2L, '-' to -1L, '0' to 0L, '1' to 1L, '2' to 2L)
val to = mapOf(0L to '0', 1L to '1', 2L to '2', 3L to '=', 4L to '-')

fun Long.Companion.fromSNAFU(snafu: String, base: Long) = snafu.fold(0L) { number, c ->
    number * base + from[c]!!
}

fun Long.toSNAFU(base: Long) = generateSequence(this) { (it + 2) / base }
    .takeWhile { it != 0L }
    .map { to[(it % base)]!! }
    .joinToString("")
    .reversed()

fun main() {
    fun part1(input: List<String>) =
        input.fold(0L) { result, snafu -> result + Long.fromSNAFU(snafu, 5) }
            .toSNAFU(5)

    fun part2(input: List<String>) = Unit

    val testInput = readInput("day25/Day25_test")
    check(part1(testInput) == "2=-1=0")
    // check(part2(testInput) == 54)

    val input = readInput("day25/Day25")
    println(part1(input))
    // println(part2(input))
}
