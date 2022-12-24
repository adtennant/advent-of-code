package day23

import readInput
import java.util.Collections

data class Point2D(val x: Int, val y: Int)

typealias Elf = Point2D

fun List<String>.toElves() =
    flatMapIndexed { y, line ->
        line.withIndex()
            .filter { (_, value) -> value == '#' }
            .map { (x, _) -> Elf(x, y) }
    }.toSet()

enum class Direction {
    NORTH, SOUTH, WEST, EAST;

    fun target(elf: Elf) = when (this) {
        NORTH -> Point2D(elf.x, elf.y - 1)
        SOUTH -> Point2D(elf.x, elf.y + 1)
        WEST -> Point2D(elf.x - 1, elf.y)
        EAST -> Point2D(elf.x + 1, elf.y)
    }

    fun checks(elf: Elf) = when (this) {
        NORTH -> listOf(elf.x, elf.x - 1, elf.x + 1).map { Point2D(it, elf.y - 1) }
        SOUTH -> listOf(elf.x, elf.x - 1, elf.x + 1).map { Point2D(it, elf.y + 1) }
        WEST -> listOf(elf.y, elf.y - 1, elf.y + 1).map { Point2D(elf.x - 1, it) }
        EAST -> listOf(elf.y, elf.y - 1, elf.y + 1).map { Point2D(elf.x + 1, it) }
    }
}

typealias Proposals = Map<Point2D, List<Elf>>

fun proposals(elves: Set<Elf>, directions: List<Direction>): Proposals =
    elves.fold(mutableMapOf<Point2D, MutableList<Elf>>()) { acc, elf ->
        val allEmpty = directions.flatMap { it.checks(elf) }.all { !elves.contains(it) }
        if (allEmpty) {
            return@fold acc
        }

        for (direction in directions) {
            val options = direction.checks(elf)

            if (options.all { !elves.contains(it) }) {
                val target = direction.target(elf)

                val elves = acc.getOrDefault(target, mutableListOf())
                elves.add(elf)

                acc.put(target, elves)
                break;
            }
        }

        acc
    }

fun move(elves: Set<Elf>, proposals: Proposals) = proposals.keys.fold(elves.toMutableList()) { acc, proposal ->
    val elves = proposals.get(proposal)!!

    if (elves.size == 1) {
        val elf = elves.first()
        acc.remove(elf)
        acc.add(proposal)
    }

    acc
}.toSet()

fun main() {
    fun part1(input: List<String>): Int {
        val (result, _) = (0 until 10)
            .fold(
                input.toElves() to listOf(
                    Direction.NORTH,
                    Direction.SOUTH,
                    Direction.WEST,
                    Direction.EAST
                )
            ) { (elves, directions), _ ->
                val proposals = proposals(elves, directions)
                val next = move(elves, proposals)
                Collections.rotate(directions, -1)
                next to directions
            }

        val minX = result.minOf { it.x }
        val minY = result.minOf { it.y }
        val maxX = result.maxOf { it.x }
        val maxY = result.maxOf { it.y }

        val width = maxX - minX + 1
        val height = maxY - minY + 1

        val area = width * height
        return area - result.size
    }

    fun part2(input: List<String>): Int {
        var elves = input.toElves()
        val directions = mutableListOf(
            Direction.NORTH,
            Direction.SOUTH,
            Direction.WEST,
            Direction.EAST
        )

        var rounds = 0

        while (true) {
            val proposals = proposals(elves, directions)
            val next = move(elves, proposals)
            Collections.rotate(directions, -1)
            rounds += 1

            val moved = next.subtract(elves).size

            if (moved == 0) {
                break
            }

            elves = next
        }

        return rounds
    }

    val testInput = readInput("day23/Day23_test")
    check(part1(testInput) == 110)
    check(part2(testInput) == 20)

    val input = readInput("day23/Day23")
    println(part1(input))
    println(part2(input))
}
