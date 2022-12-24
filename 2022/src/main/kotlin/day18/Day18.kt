package day18

import readInput
import takeExact

data class Cube(val x: Int, val y: Int, val z: Int)

fun Cube.neighbours() = listOf(
    Cube(x - 1, y, z),
    Cube(x + 1, y, z),
    Cube(x, y - 1, z),
    Cube(x, y + 1, z),
    Cube(x, y, z - 1),
    Cube(x, y, z + 1),
)

fun String.toCube() = split(",")
    .takeExact(3)
    .let { Cube(it[0].toInt(), it[1].toInt(), it[2].toInt()) }

fun main() {
    fun isExterior(cube: Cube, unique: Set<Cube>): Boolean {
        val minX = unique.minOf { it.x }
        val maxX = unique.maxOf { it.x }
        val minY = unique.minOf { it.y }
        val maxY = unique.maxOf { it.y }
        val minZ = unique.minOf { it.z }
        val maxZ = unique.maxOf { it.z }

        fun checkLimits(cube: Cube) = cube.x in minX..maxX &&
                cube.y in minY..maxY &&
                cube.z in minZ..maxZ

        val visited = mutableSetOf(cube)
        val queue = ArrayDeque(listOf(cube))

        while (queue.isNotEmpty()) {
            val next = queue.removeFirst()

            if (!checkLimits(next)) {
                return true
            }

            val neighbours = next.neighbours().filter { !unique.contains(it) && !visited.contains(it) }
            neighbours.forEach {
                visited.add(it)
                queue.add(it)
            }
        }

        return false
    }

    fun surfaceArea(input: List<String>, countPrecicate: (Cube, Set<Cube>) -> Boolean = { _, _ -> true }): Int {
        val cubes = input.map(String::toCube).toSet()

        return cubes.fold(0) { acc, cube ->
            val surfaces = cube.neighbours()
                .filter { !cubes.contains(it) }

            acc + surfaces.count { countPrecicate(it, cubes) }
        }
    }

    fun part1(input: List<String>) = surfaceArea(input)

    fun part2(input: List<String>) = surfaceArea(input) { it, cubes -> isExterior(it, cubes) }

    val testInput = readInput("day18/Day18_test")
    check(part1(testInput) == 64)
    check(part2(testInput) == 58)

    val input = readInput("day18/Day18")
    println(part1(input))
    println(part2(input))
}
