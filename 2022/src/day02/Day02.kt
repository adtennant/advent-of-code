package day02

import readInput

enum class Outcome {
    WIN, DRAW, LOSE;

    val points
        get() = when (this) {
            WIN -> 6
            DRAW -> 3
            LOSE -> 0
        }
}

fun String.toOutcome() = when (this) {
    "X" -> Outcome.LOSE
    "Y" -> Outcome.DRAW
    "Z" -> Outcome.WIN
    else -> throw Error()
}

enum class Choice {
    ROCK, PAPER, SCISSORS;

    val points
        get() = when (this) {
            ROCK -> 1
            PAPER -> 2
            SCISSORS -> 3
        }
}

fun String.toChoice() = when (this) {
    "A", "X" -> Choice.ROCK
    "B", "Y" -> Choice.PAPER
    "C", "Z" -> Choice.SCISSORS
    else -> throw Error()
}

class DynamicRound(private val player: Choice, private val opponent: Choice) {
    private val outcome
        get() =
            when (player) {
                Choice.ROCK -> when (opponent) {
                    Choice.ROCK -> Outcome.DRAW
                    Choice.PAPER -> Outcome.LOSE
                    Choice.SCISSORS -> Outcome.WIN
                }
                Choice.PAPER -> when (opponent) {
                    Choice.ROCK -> Outcome.WIN
                    Choice.PAPER -> Outcome.DRAW
                    Choice.SCISSORS -> Outcome.LOSE
                }
                Choice.SCISSORS -> when (opponent) {
                    Choice.ROCK -> Outcome.LOSE
                    Choice.PAPER -> Outcome.WIN
                    Choice.SCISSORS -> Outcome.DRAW
                }
            }

    val score get() = player.points + outcome.points
}

fun String.toDynamicRound() =
    split(" ", limit = 2).windowed(2, 2).firstNotNullOf { DynamicRound(it[1].toChoice(), it[0].toChoice()) }

class FixedRound(private val opponent: Choice, private val outcome: Outcome) {
    private val player
        get() =
            when (opponent) {
                Choice.ROCK -> when (outcome) {
                    Outcome.WIN -> Choice.PAPER
                    Outcome.DRAW -> Choice.ROCK
                    Outcome.LOSE -> Choice.SCISSORS
                }
                Choice.PAPER -> when (outcome) {
                    Outcome.WIN -> Choice.SCISSORS
                    Outcome.DRAW -> Choice.PAPER
                    Outcome.LOSE -> Choice.ROCK
                }
                Choice.SCISSORS -> when (outcome) {
                    Outcome.WIN -> Choice.ROCK
                    Outcome.DRAW -> Choice.SCISSORS
                    Outcome.LOSE -> Choice.PAPER
                }
            }

    val score get() = player.points + outcome.points
}

fun String.toFixedRound() =
    split(" ", limit = 2).windowed(2, 2).firstNotNullOf { FixedRound(it[0].toChoice(), it[1].toOutcome()) }

fun main() {
    fun part1(input: List<String>) = input.map { it.toDynamicRound() }.sumOf { it.score }

    fun part2(input: List<String>) = input.map { it.toFixedRound() }.sumOf { it.score }

    val testInput = readInput("day02/Day02_test")
    check(part1(testInput) == 15)
    check(part2(testInput) == 12)

    val input = readInput("day02/Day02")
    println(part1(input))
    println(part2(input))
}
