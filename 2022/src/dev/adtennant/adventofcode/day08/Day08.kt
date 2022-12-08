package dev.adtennant.adventofcode.day08

import readInput
import takeWhileInclusive

data class Tree(val height: Int, val row: Int, val column: Int)

typealias Grid = List<Tree>

fun Grid.column(column: Int) = this.filter { it.column == column }
fun Grid.row(row: Int) = this.filter { it.row == row }

fun Grid.toLeft(tree: Tree) = this.row(tree.row)
    .filter { it.column < tree.column }
    .sortedBy { it.column }
    .reversed()

fun Grid.toRight(tree: Tree) = this.row(tree.row)
    .filter { it.column > tree.column }
    .sortedBy { it.column }

fun Grid.toTop(tree: Tree) = this.column(tree.column)
    .filter { it.row < tree.row }
    .sortedBy { it.row }
    .reversed()

fun Grid.toBottom(tree: Tree) = this.column(tree.column)
    .filter { it.row > tree.row }
    .sortedBy { it.row }

fun Grid.onEdge(tree: Tree): Boolean {
    val maxRow = this.maxOf { it.row }
    val maxColumn = this.maxOf { it.column }

    return tree.row == 0 ||
        tree.row == maxRow ||
        tree.column == 0 ||
        tree.column == maxColumn
}

fun Grid.visibleFromLeft(tree: Tree) = this.toLeft(tree)
    .all { it.height < tree.height }

fun Grid.visibleFromRight(tree: Tree) = this.toRight(tree)
    .all { it.height < tree.height }

fun Grid.visibleFromTop(tree: Tree) = this.toTop(tree)
    .all { it.height < tree.height }

fun Grid.visibleFromBottom(tree: Tree) = this.toBottom(tree)
    .all { it.height < tree.height }

fun Grid.isVisible(tree: Tree) = onEdge(tree) ||
    visibleFromLeft(tree) ||
    visibleFromRight(tree) ||
    visibleFromTop(tree) ||
    visibleFromBottom(tree)

fun Grid.countVisible() = count { tree ->
    isVisible(tree)
}

fun Grid.treesVisibleToLeft(tree: Tree) = this.toLeft(tree)
    .takeWhileInclusive { it.height < tree.height }
    .count()

fun Grid.treesVisibleToRight(tree: Tree) = this.toRight(tree)
    .takeWhileInclusive { it.height < tree.height }
    .count()

fun Grid.treesVisibleToTop(tree: Tree) = this.toTop(tree)
    .takeWhileInclusive { it.height < tree.height }
    .count()

fun Grid.treesVisibleToBottom(tree: Tree) = this.toBottom(tree)
    .takeWhileInclusive { it.height < tree.height }
    .count()

fun Grid.scenicScore(tree: Tree) =
    treesVisibleToLeft(tree) * treesVisibleToRight(tree) * treesVisibleToTop(tree) * treesVisibleToBottom(tree)

fun Grid.maxScenicScore() = maxOf { tree -> scenicScore(tree) }

fun List<String>.toGrid() = flatMapIndexed { row, line ->
    line.mapIndexed { column, height ->
        Tree(
            height.toString().toInt(),
            row,
            column
        )
    }
}

fun main() {
    fun part1(input: List<String>) = input.toGrid().countVisible()

    fun part2(input: List<String>) = input.toGrid().maxScenicScore()

    val testInput = readInput("day08/Day08_test")
    check(part1(testInput) == 21)
    check(part2(testInput) == 8)

    val input = readInput("day08/Day08")
    println(part1(input))
    println(part2(input))
}
