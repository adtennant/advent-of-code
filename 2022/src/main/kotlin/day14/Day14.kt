package day14

import readInput
import takeExact

enum class Cell {
    ROCK, SAND, AIR;

    override fun toString() = when (this) {
        ROCK -> "#"
        SAND -> "o"
        AIR -> "."
    }
}

typealias Point = Pair<Int, Int>

val Point.x get() = first
val Point.y get() = second

class Cave(input: List<String>, hasFloor: Boolean) : Iterator<Point> {
    private val cells = run {
        val rocks = input.map { line ->
            line.split(" -> ")
                .map {
                    it.split(",")
                        .takeExact(2)
                        .let { parts -> Point(parts[0].toInt(), parts[1].toInt()) }
                }
        }.flatMap { line ->
            line.windowed(2).flatMap {
                if (it[0].x == it[1].x) {
                    // Vertical
                    val from = minOf(it[0].y, it[1].y)
                    val to = maxOf(it[0].y, it[1].y)
                    (from..to).map { y -> Point(it[0].x, y) }
                } else {
                    // Horizontal
                    val from = minOf(it[0].x, it[1].x)
                    val to = maxOf(it[0].x, it[1].x)
                    (from..to).map { x -> Point(x, it[0].y) }
                }
            }
        }

        val maxX = rocks.maxOf { it.x }
        val maxY = rocks.maxOf { it.y }

        // Make the map massive to give sand plenty of space to spread
        val cells = MutableList(maxY * 2) { List(maxX * 2) { Cell.AIR } }

        cells.mapIndexed { y, row ->
            MutableList(row.size) { x ->
                if (rocks.contains(Point(x, y))) {
                    Cell.ROCK
                } else if (hasFloor && y == maxY + 2) {
                    Cell.ROCK
                } else {
                    Cell.AIR
                }
            }
        }
    }

    private fun getOrNull(x: Int, y: Int) = cells.getOrNull(y)?.getOrNull(x)

    override fun hasNext() = !cells.last().contains(Cell.SAND) &&
        getOrNull(500, 0) != Cell.SAND

    override fun next(): Point {
        var x = 500
        var y = 0

        while (y < cells.size) {
            when (Cell.AIR) {
                getOrNull(x, y + 1) -> {
                    y++
                }
                getOrNull(x - 1, y + 1) -> {
                    y++
                    x--
                }
                getOrNull(x + 1, y + 1) -> {
                    y++
                    x++
                }
                else -> {
                    cells[y][x] = Cell.SAND
                    return Point(x, y)
                }
            }
        }

        error("unreachable")
    }
}

fun main() {
    fun part1(input: List<String>) = Cave(input, false)
        .asSequence()
        .count() - 1 // Ignore the grain that fell into the void

    fun part2(input: List<String>) = Cave(input, true)
        .asSequence()
        .count()

    val testInput = readInput("day14/Day14_test")
    check(part1(testInput) == 24)
    check(part2(testInput) == 93)

    val input = readInput("day14/Day14")
    println(part1(input))
    println(part2(input))
}
