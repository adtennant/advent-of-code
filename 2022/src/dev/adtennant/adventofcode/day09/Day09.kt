package dev.adtennant.adventofcode.day09

import readInput
import kotlin.math.absoluteValue

enum class Direction {
    UP, DOWN, LEFT, RIGHT
}

fun String.toDirection() = when (this) {
    "U" -> Direction.UP
    "D" -> Direction.DOWN
    "L" -> Direction.LEFT
    "R" -> Direction.RIGHT
    else -> error("invalid direction")
}

data class Move(val direction: Direction, val distance: Int)

fun String.toMove() = split(" ")
    .let { Move(it[0].toDirection(), it[1].toInt()) }

data class Point(val x: Int, val y: Int) {
    fun isTouching(other: Point) = (x - other.x).absoluteValue < 2 && (y - other.y).absoluteValue < 2

    fun moveTowards(other: Point) = Point(
        when {
            x == other.x -> x
            x <= other.x -> x + 1
            else -> x - 1
        },
        when {
            y == other.y -> y
            y <= other.y -> y + 1
            else -> y - 1
        }
    )
}

data class Rope(val knots: List<Point>) {
    constructor(length: Int) : this(List(length) { Point(0, 0) })

    fun apply(move: Move) = (0 until move.distance)
        .fold(listOf(this)) { acc, _ ->
            val prev = acc.last()
            acc + prev.move(move.direction)
        }
        .drop(1)

    private fun move(direction: Direction): Rope {
        val first = when (direction) {
            Direction.UP -> Point(knots[0].x, knots[0].y - 1)
            Direction.DOWN -> Point(knots[0].x, knots[0].y + 1)
            Direction.LEFT -> Point(knots[0].x - 1, knots[0].y)
            Direction.RIGHT -> Point(knots[0].x + 1, knots[0].y)
        }

        return knots
            .drop(1)
            .fold(listOf(first)) { rope, point ->
                val prev = rope.last()
                val next = if (!point.isTouching(prev)) {
                    point.moveTowards(prev)
                } else {
                    point
                }

                rope + next
            }
            .let { Rope(it) }
    }
}

fun main() {
    fun countTailPositions(input: List<String>, length: Int) = input.map(String::toMove)
        .fold(listOf(Rope(length))) { acc, move ->
            val next = acc.last().apply(move)
            acc + next
        }
        .distinctBy { it.knots.last() }
        .count()

    fun part1(input: List<String>) = countTailPositions(input, 2)

    fun part2(input: List<String>) = countTailPositions(input, 10)

    val testInput = readInput("day09/Day09_test")
    check(part1(testInput) == 13)
    check(part2(testInput) == 1)

    val testInput2 = readInput("day09/Day09_test_part2")
    check(part2(testInput2) == 36)

    val input = readInput("day09/Day09")
    println(part1(input))
    println(part2(input))
}
