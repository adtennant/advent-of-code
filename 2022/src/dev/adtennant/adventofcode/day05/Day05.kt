package dev.adtennant.adventofcode.day05

import readInput
import split
import takeExact
import java.util.*

typealias Stacks = Array<Stack<Char>>

val Stacks.message
    get() = map(Stack<Char>::peek)
        .joinToString("")

class Move(val count: Int, val from: Int, val to: Int)

interface CrateMover {
    fun apply(stacks: Stacks, move: Move): Stacks
    fun applyAll(stacks: Stacks, moves: List<Move>) = moves.fold(stacks) { acc, move -> apply(acc, move) }
}

class CrateMover9000 : CrateMover {
    override fun apply(stacks: Stacks, move: Move) =
        (0 until move.count).fold(stacks) { acc, _ ->
            val crate = acc[move.from - 1].pop()
            acc[move.to - 1].push(crate)

            acc
        }
}

class CrateMover9001 : CrateMover {
    override fun apply(stacks: Stacks, move: Move) =
        (0 until move.count).map { stacks[move.from - 1].pop() }
            .reversed()
            .fold(stacks) { acc, crate ->
                acc[move.to - 1].push(crate)
                acc
            }
}

fun List<String>.toStacks(): Stacks {
    val count = last()
        .trim()
        .takeLastWhile { it.isDigit() }
        .toInt()

    return reversed()
        .drop(1)
        .fold(Array(count) { Stack<Char>() }) { stacks, line ->
            for (i in 0 until count) {
                val c = line[1 + i * 4]

                if (c.isLetter()) {
                    stacks[i].push(c)
                }
            }

            stacks
        }
}

val moveRegex = "move (\\d*) from (\\d*) to (\\d*)".toRegex()

fun String.toMove() =
    moveRegex.find(this)!!
        .let(MatchResult::groupValues)
        .let { Move(it[1].toInt(), it[2].toInt(), it[3].toInt()) }

fun main() {
    fun getStacksAndMoves(input: List<String>) = input.split(String::isNotEmpty)
        .takeExact(2)
        .let {
            Pair(
                it[0].toStacks(),
                it[1].map(String::toMove)
            )
        }

    fun getMessage(input: List<String>, mover: CrateMover) = getStacksAndMoves(input)
        .let { mover.applyAll(it.first, it.second) }
        .let(Stacks::message)

    fun part1(input: List<String>) = getMessage(input, CrateMover9000())

    fun part2(input: List<String>) = getMessage(input, CrateMover9001())

    val testInput = readInput("day05/Day05_test")
    check(part1(testInput) == "CMZ")
    check(part2(testInput) == "MCD")

    val input = readInput("day05/Day05")
    println(part1(input))
    println(part2(input))
}
