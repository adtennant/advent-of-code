package dev.adtennant.adventofcode.day12

import readInput
import java.util.LinkedList

data class Cell(val x: Int, val y: Int, val c: Char) {
    val height
        get() = when (c) {
            'S' -> 'a'
            'E' -> 'z'
            else -> c
        }.code
}

class Map(input: List<String>) {
    private val grid = input.flatMapIndexed { y, row -> row.mapIndexed { x, c -> Cell(x, y, c) } }

    fun find(c: Char) = grid.find { it.c == c }
    fun findAll(c: Char) = grid.filter { it.c == c }
    private fun getOrNull(x: Int, y: Int) = grid.firstOrNull { it.x == x && it.y == y }

    private fun neighbours(cell: Cell): List<Cell> = listOfNotNull(
        getOrNull(cell.x, cell.y - 1),
        getOrNull(cell.x, cell.y + 1),
        getOrNull(cell.x - 1, cell.y),
        getOrNull(cell.x + 1, cell.y)
    ).filter { neighbour -> (neighbour.height - cell.height) <= 1 }

    fun findPath(from: Cell = find('S')!!, to: Cell = find('E')!!): List<Cell> {
        val q = LinkedList<Cell>()
        q.add(from)

        val explored = mutableMapOf<Cell, Cell?>(from to null)

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
}

fun main() {
    fun part1(input: List<String>) = Map(input).findPath().size

    fun part2(input: List<String>) = Map(input).let { grid ->
        val starts = grid.findAll('a') + grid.find('S')!!
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
