package dev.adtennant.adventofcode.day12

import indexesOf
import readInput
import java.lang.Thread.yield
import java.util.LinkedList

typealias Point = Pair<Int, Int>

val Point.x get() = first
val Point.y get() = second

operator fun Point.plus(other: Point) = Point(x + other.x, y + other.y)

class Grid(input: List<String>) {
    private val data = input.flatMap { it.toCharArray().toList() }
    private val width = input.first().length
    private val height = input.size

    fun find(c: Char) = data.indexOf(c)
        .let { Point(it % width, it / width) }

    fun findAll(c: Char) = data.indexesOf(c)
        .map { Point(it % width, it / width) }

    private fun getHeight(point: Point) = data[point.y * width + point.x].let {
        when (it) {
            'S' -> 'a'
            'E' -> 'z'
            else -> it
        }
    }

    private fun inBounds(point: Point) = point.x in 0 until width && point.y in 0 until height

    fun findPath(from: Point = find('S'), to: Point = find('E')): List<Point> {
        val q = LinkedList<Point>()
        q.add(from)

        val explored = mutableMapOf<Point, Point?>(from to null)

        while (!q.isEmpty()) {
            val v = q.remove()

            if (v == to) {
                return buildList {
                    var current = to

                    while (current != from) {
                        add(current)
                        current = explored[current]!!
                    }
                }.reversed()
            }

            for (w in neighbours(v)) {
                if (!explored.containsKey(w)) {
                    explored[w] = v
                    q.add(w)
                }
            }
        }

        return emptyList()
    }

    private fun neighbours(point: Point): List<Point> = getHeight(point)
        .let { current ->
            listOf(
                Point(0, -1),
                Point(0, 1),
                Point(-1, 0),
                Point(1, 0)
            )
                .map { point + it }
                .filter { inBounds(it) }
                .filter {
                    val neighbour = getHeight(it)
                    (neighbour.code - current.code) <= 1
                }
        }
}

fun main() {
    fun part1(input: List<String>) = Grid(input).findPath().size

    fun part2(input: List<String>) = Grid(input).let { grid ->
        val starts = grid.findAll('a') + grid.find('S')
        starts.map { start -> grid.findPath(start) }
            .filter { it.isNotEmpty() }
            .minOf { it.size }
    }

    val testInput = readInput("day12/Day12_test")
    check(part1(testInput) == 31)
    check(part2(testInput) == 29)

    val input = readInput("day12/Day12")
    println(part1(input))
    println(part2(input))
}
