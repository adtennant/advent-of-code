package day22

import readInput
import split

data class Point2D(val x: Int, val y: Int) {
    operator fun plus(v: Vector2D): Point2D = Point2D(x + v.x, y + v.y)
    operator fun minus(v: Vector2D): Point2D = Point2D(x - v.x, y - v.y)
}

data class Vector2D(val x: Int, val y: Int)

data class Rectangle(val min: Point2D, val max: Point2D) {
    fun contains(point: Point2D) = point.x >= min.x &&
            point.y >= min.y &&
            point.x < max.x &&
            point.y < max.y
}

enum class Facing {
    RIGHT, DOWN, LEFT, UP;

    fun delta() = when (this) {
        RIGHT -> Vector2D(1, 0)
        DOWN -> Vector2D(0, 1)
        LEFT -> Vector2D(-1, 0)
        UP -> Vector2D(0, -1)
    }

    fun rotate(rotation: Rotation) = when (rotation) {
        Rotation.CLOCKWISE -> when (this) {
            RIGHT -> DOWN
            DOWN -> LEFT
            LEFT -> UP
            UP -> RIGHT
        }
        Rotation.COUNTERCLOCKWISE -> when (this) {
            RIGHT -> UP
            DOWN -> RIGHT
            LEFT -> DOWN
            UP -> LEFT
        }
    }

    fun toInt() = when (this) {
        RIGHT -> 0
        DOWN -> 1
        LEFT -> 2
        UP -> 3
    }
}

enum class Tile {
    OPEN, WALL
}

fun Char.toTileOrNull() = when (this) {
    ' ' -> null
    '.' -> Tile.OPEN
    '#' -> Tile.WALL
    else -> error("invalid tile")
}

fun List<String>.toCells() = map { row -> row.map { cell -> cell.toTileOrNull() } }

data class Cube(val cells: List<List<Tile?>>) {
    private fun getOrNull(point: Point2D) = cells.getOrNull(point.y)?.getOrNull(point.x)

    private fun wrappingGetOrNull(
        point: Point2D,
        facing: Facing,
        wrapFn: WrappingFn
    ): Triple<Point2D, Facing, Tile> {
        val cell = getOrNull(point)

        if (cell == null) {
            val (next, nextFacing) = wrapFn(cells, point, facing)
            return Triple(next, nextFacing, getOrNull(next)!!)
        }

        return Triple(point, facing, cell)
    }

    fun walk(
        start: Point2D,
        steps: Int,
        facing: Facing,
        wrapFn: WrappingFn
    ): Pair<Point2D, Facing> {
        var currentPosition = start
        var currentFacing = facing

        repeat(steps) {
            val delta = currentFacing.delta()
            val (nextPosition, nextFacing, tile) = wrappingGetOrNull(currentPosition + delta, currentFacing, wrapFn)

            if (tile != Tile.WALL) {
                currentPosition = nextPosition
                currentFacing = nextFacing
            }
        }

        return currentPosition to currentFacing
    }
}

enum class Rotation {
    CLOCKWISE, COUNTERCLOCKWISE
}

fun String.toRotation() = when (this) {
    "R" -> Rotation.CLOCKWISE
    "L" -> Rotation.COUNTERCLOCKWISE
    else -> error("invalid direction")
}

sealed class Move {
    data class Forward(val distance: Int) : Move()
    data class Turn(val rotation: Rotation) : Move()
}

fun String.toMove() = when {
    this.first().isDigit() -> Move.Forward(this.toInt())
    this.first().isLetter() -> Move.Turn(this.toRotation())
    else -> error("invalid move")
}

fun String.toMoves(): List<Move> {
    val result = mutableListOf<Move>()
    var remainder = this

    while (remainder.isNotEmpty()) {
        val next = if (remainder.first().isDigit()) {
            remainder.takeWhile(Char::isDigit)
        } else {
            remainder.takeWhile(Char::isLetter)
        }

        result.add(next.toMove())
        remainder = remainder.drop(next.length)
    }

    return result
}

fun List<String>.toCubeAndMoves() = this.split(String::isEmpty)
    .let { Cube(it[0].toCells()) to it[1].first().toMoves() }

data class Actor(val position: Point2D, val facing: Facing)

typealias Square = Pair<Rectangle, Map<Facing, (Point2D) -> Pair<Point2D, Facing>>>

val testSquares = listOf<Square>(
    Rectangle(Point2D(8, 0), Point2D(12, 4)) to mapOf(
        Facing.UP to { point -> Point2D(3 - point.x, 4) to Facing.DOWN },
        Facing.LEFT to { point -> Point2D(4 + point.y, 4) to Facing.DOWN },
        Facing.RIGHT to { point -> Point2D(15, 11 - point.y) to Facing.LEFT }),
    Rectangle(Point2D(0, 4), Point2D(4, 8)) to mapOf(
        Facing.UP to { point -> Point2D(11 - point.x, 0) to Facing.DOWN },
        Facing.LEFT to { point -> Point2D(20 - point.y, 11) to Facing.UP },
        Facing.DOWN to { point -> Point2D(12 - point.x, 11) to Facing.UP }),
    Rectangle(Point2D(4, 4), Point2D(8, 8)) to mapOf(
        Facing.UP to { point -> Point2D(8, point.x - 4) to Facing.RIGHT },
        Facing.DOWN to { point -> Point2D(8, 15 - point.x) to Facing.RIGHT }),
    Rectangle(Point2D(8, 4), Point2D(12, 8)) to mapOf(
        Facing.RIGHT to { point -> Point2D(19 - point.y, 9) to Facing.DOWN }),
    Rectangle(Point2D(8, 8), Point2D(12, 12)) to mapOf(
        Facing.LEFT to { point -> Point2D(15 - point.y, 7) to Facing.UP },
        Facing.DOWN to { point -> Point2D(11 - point.x, 7) to Facing.UP }),
    Rectangle(Point2D(12, 8), Point2D(16, 12)) to mapOf(
        Facing.UP to { point -> Point2D(11, 19 - point.x) to Facing.LEFT },
        Facing.RIGHT to { point -> Point2D(11, 11 - point.y) to Facing.LEFT },
        Facing.DOWN to { point -> Point2D(0, 19 - point.y) to Facing.RIGHT }),
)

val squares = listOf<Square>(
    Rectangle(Point2D(50, 0), Point2D(100, 50)) to mapOf(
        Facing.UP to { point -> Point2D(0, 100 + point.x) to Facing.RIGHT }, // 0 -> 5
        Facing.LEFT to { point -> Point2D(0, 149 - point.y) to Facing.RIGHT }), // 0 -> 3
    Rectangle(Point2D(100, 0), Point2D(150, 50)) to mapOf(
        Facing.UP to { point -> Point2D(point.x - 100, 199) to Facing.UP }, // 1 -> 5
        Facing.RIGHT to { point -> Point2D(99, 149 - point.y) to Facing.LEFT }, // 1 -> 4
        Facing.DOWN to { point -> Point2D(99, point.x - 50) to Facing.LEFT }), // 1 -> 2
    Rectangle(Point2D(50, 50), Point2D(100, 100)) to mapOf(
        Facing.LEFT to { point -> Point2D(point.y - 50, 100) to Facing.DOWN }, // 2 -> 3
        Facing.RIGHT to { point -> Point2D(point.y + 50, 49) to Facing.UP }), // 2 -> 1
    Rectangle(Point2D(0, 100), Point2D(50, 150)) to mapOf(
        Facing.UP to { point -> Point2D(50, point.x + 50) to Facing.RIGHT }, // 3 -> 2
        Facing.LEFT to { point -> Point2D(50, 149 - point.y) to Facing.RIGHT }), // 3 -> 0
    Rectangle(Point2D(50, 100), Point2D(100, 150)) to mapOf(
        Facing.RIGHT to { point -> Point2D(149, 149 - point.y) to Facing.LEFT }, // 4 -> 1
        Facing.DOWN to { point -> Point2D(49, 100 + point.x) to Facing.LEFT }), // 4 -> 5
    Rectangle(Point2D(0, 150), Point2D(50, 200)) to mapOf(
        Facing.LEFT to { point -> Point2D(point.y - 100, 0) to Facing.DOWN }, // 5 -> 0
        Facing.RIGHT to { point -> Point2D(point.y - 100, 149) to Facing.UP }, // 5 -> 4
        Facing.DOWN to { point -> Point2D(point.x + 100, 0) to Facing.DOWN }), // 5 -> 1
)

typealias WrappingFn = (List<List<Tile?>>, Point2D, Facing) -> Pair<Point2D, Facing>

fun main() {
    fun findPassword(input: List<String>, wrapFn: WrappingFn): Int {
        val (cube, moves) = input.toCubeAndMoves()

        val initialPosition = Point2D(cube.cells.first().indexOfFirst { it != null }, 0)
        val initialFacing = Facing.RIGHT

        val result = moves.fold(Actor(initialPosition, initialFacing)) { acc, move ->
            when (move) {
                is Move.Forward -> {
                    val (position, facing) = cube.walk(acc.position, move.distance, acc.facing, wrapFn)
                    Actor(position, facing)
                }
                is Move.Turn -> {
                    val facing = acc.facing.rotate(move.rotation)
                    Actor(acc.position, facing)
                }
            }
        }

        return (result.position.y + 1) * 1000 + (result.position.x + 1) * 4 + result.facing.toInt()
    }

    fun part1(input: List<String>) = findPassword(input) { cells, point, facing ->
        when (facing) {
            Facing.RIGHT -> cells[point.y].withIndex().first { (_, tile) -> tile != null }
                .let { point.copy(x = it.index) }
            Facing.DOWN -> cells.map { row -> row.getOrNull(point.x) }.withIndex()
                .first { (_, tile) -> tile != null }
                .let { point.copy(y = it.index) }
            Facing.LEFT -> cells[point.y].withIndex().last { (_, tile) -> tile != null }
                .let { point.copy(x = it.index) }
            Facing.UP -> cells.map { row -> row.getOrNull(point.x) }.withIndex()
                .last { (_, tile) -> tile != null }
                .let { point.copy(y = it.index) }
        } to facing
    }

    fun cubeWrappingFn(squares: List<Square>) = { _: List<List<Tile?>>, point: Point2D, facing: Facing ->
        val start = point - facing.delta()

        val currentSquare = squares.first { it.first.contains(start) }
        currentSquare.second[facing]!!(start)
    }

    fun part2(input: List<String>, wrapFn: WrappingFn) =
        findPassword(input, wrapFn)

    val testInput = readInput("day22/Day22_test")
    check(part1(testInput) == 6032)
    check(part2(testInput, cubeWrappingFn(testSquares)) == 5031)

    val input = readInput("day22/Day22")
    println(part1(input))
    println(part2(input, cubeWrappingFn(squares)))
}
