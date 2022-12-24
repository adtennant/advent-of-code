package day24

import readInput

data class Point2D(val x: Int, val y: Int)

data class BlizzardState(val cells: Map<Point2D, List<Char>>, val width: Int, val height: Int) {
    fun get(point: Point2D) = cells[point]

    fun next() = cells
        .flatMap { (position, items) ->
            items
                .map { item ->
                    when (item) {
                        '^' -> position.copy(y = if (position.y == 1) height - 2 else position.y - 1)
                        'v' -> position.copy(y = if (position.y == height - 2) 1 else position.y + 1)
                        '<' -> position.copy(x = if (position.x == 1) width - 2 else position.x - 1)
                        '>' -> position.copy(x = if (position.x == width - 2) 1 else position.x + 1)
                        '#' -> position
                        else -> error("invalid thing")
                    } to item
                }
        }
        .groupBy({ it.first }, { it.second })
        .let { BlizzardState(it, width, height) }
}

data class State(val position: Point2D, val time: Int)

class Blizzard(input: List<String>) {
    private val states = buildList {
        val cells = input.flatMapIndexed { y, row ->
            row.withIndex()
                .filter { it.value != '.' }
                .map { (x, c) -> Point2D(x, y) to listOf(c) }
        }.toMap()

        val initialState = BlizzardState(cells, input.first().length, input.size)
        var state = initialState

        do {
            add(state)
            state = state.next()
        } while (state != initialState)
    }

    private val width = input.first().length
    private val height = input.size

    val start = Point2D(input.first().indexOfFirst { it == '.' }, 0)
    val end = Point2D(input.last().indexOfFirst { it == '.' }, height - 1)

    private fun getOrNull(point: Point2D, time: Int) = getState(time).get(point)
    private fun getState(time: Int) = states[time % states.size]

    private fun neighbours(state: State) = listOf(
        State(state.position.copy(x = state.position.x - 1), time = state.time + 1),
        State(state.position.copy(x = state.position.x + 1), time = state.time + 1),
        State(state.position.copy(y = state.position.y - 1), time = state.time + 1),
        State(state.position.copy(y = state.position.y + 1), time = state.time + 1),
        state.copy(time = state.time + 1)
    )
        .filter { it.position.x in 0 until width && it.position.y in 0 until height }
        .filter { getOrNull(it.position, it.time) == null }

    fun findShortestPath(
        from: Point2D = start,
        to: Point2D = end,
        time: Int = 0
    ): Int {
        val queue = ArrayDeque(listOf(State(from, time)))
        val seen = mutableSetOf<Pair<Point2D, Int>>()
        var minSteps = Int.MAX_VALUE

        while (queue.isNotEmpty()) {
            val state = queue.removeFirst()

            if (state.position == to) {
                minSteps = minOf(minSteps, state.time)
            }

            if (state.time >= minSteps) {
                continue
            }

            for (next in neighbours(state)) {
                val id = next.position to (next.time % states.size)

                if (!seen.contains(id)) {
                    seen.add(id)
                    queue.add(next)
                }
            }
        }

        return minSteps
    }
}

fun main() {
    fun part1(input: List<String>) = Blizzard(input).findShortestPath()
    fun part2(input: List<String>): Int {
        val blizzard = Blizzard(input)

        var time = blizzard.findShortestPath()
        time = blizzard.findShortestPath(blizzard.end, blizzard.start, time)
        time = blizzard.findShortestPath(time = time)

        return time
    }

    val testInput = readInput("day24/Day24_test")
    check(part1(testInput) == 18)
    check(part2(testInput) == 54)

    val input = readInput("day24/Day24")
    println(part1(input))
    println(part2(input))
}
