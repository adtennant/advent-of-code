package day17

import readInputAsText

enum class Jet {
    LEFT,
    RIGHT
}

fun jets(input: String) = input
    .map {
        when (it) {
            '<' -> Jet.LEFT
            '>' -> Jet.RIGHT
            else -> error("invalid jet")
        }
    }
    .let { Jets(it) }

data class Jets(val data: List<Jet>) : Iterator<Jet> {
    var jetIndex = 0

    override fun hasNext() = true
    override fun next() = data[jetIndex++ % data.size]
}

data class Point(val x: Int, val y: Int)

typealias Piece = List<String>

fun Piece.points() = flatMapIndexed { y, row ->
    row.withIndex()
        .filter { it.value == '#' }
        .map { (x, _) -> Point(x, y) }
}

class Pieces : Iterator<Piece> {
    var pieceIndex = 0

    override fun hasNext() = true
    override fun next() = when ((pieceIndex++ % 5)) {
        0 -> listOf("####")
        1 -> listOf(".#.", "###", ".#.")
        2 -> listOf("..#", "..#", "###")
        3 -> listOf("#", "#", "#", "#")
        4 -> listOf("##", "##")
        else -> error("unreachable")
    }
}

fun moveIfPossible(cave: MutableList<String>, pos: Point, piece: Piece, offset: Point): Pair<Point, Boolean> {
    val next = Point(pos.x + offset.x, pos.y + offset.y)

    if (piece.points().all { cave.getOrNull(next.y + it.y)?.getOrNull(next.x + it.x) == '.' }) {
        return Pair(next, true)
    }

    return Pair(pos, false)
}

fun dropShape(cave: MutableList<String>, piece: Piece, jets: Jets): Int {
    var pos = Point(2, 0)

    while (true) {
        val offset = when (jets.next()) {
            Jet.LEFT -> Point(-1, 0)
            Jet.RIGHT -> Point(1, 0)
        }
        val (moved, _) = moveIfPossible(cave, pos, piece, offset)
        pos = moved

        val (dropped, ok) = moveIfPossible(cave, pos, piece, Point(0, 1))
        pos = dropped

        if (!ok) {
            piece.points().forEach {
                val chars = cave[pos.y + it.y].toCharArray()
                chars[pos.x + it.x] = '#'
                cave[pos.y + it.y] = chars.joinToString("")
            }
            break
        }
    }

    return pos.y
}

fun detectLoop(
    cave: MutableList<String>,
    height: Long,
    memory: MutableList<Triple<String, Long, Long>>,
    window: Int
): Pair<Long, Long>? {
    val key = cave.take(window).joinToString("\n")
    val states = memory.filter { it.first == key }

    return if (states.size == 2) {
        Pair(states[1].third - states[0].third, states[1].second - states[0].second)
    } else {
        memory.add(Triple(key, height, memory.size.toLong()))
        null
    }
}

fun simulate(input: String, count: Long): Long {
    val jets = jets(input)
    val pieces = Pieces()
    val cave = mutableListOf<String>()
    var memory = mutableListOf<Triple<String, Long, Long>>()
    var foundLoop = false
    var depth = 0L
    var height = 0L
    var i = 0L

    while (i < count) {
        val piece = pieces.next()

        while (depth < height + piece.size + 3) {
            cave.add(0, ".......")
            depth++
        }

        while (depth > height + piece.size + 3) {
            cave.removeAt(0)
            depth--
        }

        val loop = detectLoop(cave, height, memory, 100)

        if (!foundLoop && loop != null) {
            val mul = (count - i) / loop.first
            i += mul * loop.first
            depth += mul * loop.second
            height += mul * loop.second
            foundLoop = true
        }

        val y = dropShape(cave, piece, jets)
        height = maxOf(height, depth - y)

        i++
    }

    return height
}

fun main() {
    fun part1(input: String) = simulate(input, 2022)

    fun part2(input: String) = simulate(input, 1000000000000L)

    val testInput = readInputAsText("day17/Day17_test")
    check(part1(testInput) == 3068L)
    check(part2(testInput) == 1514285714288L)

    val input = readInputAsText("day17/Day17")
    println(part1(input))
    println(part2(input))
}
