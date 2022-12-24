package day21

import checkedAdd
import checkedSub
import floorDiv
import readInput
import kotlin.math.sign

fun List<String>.toInstructions() = fold(emptyMap<String, String>()) { acc, s ->
    acc + s.split(": ").let { (it[0] to it[1]) }
}

data class Monkeys(val jobs: Map<String, String>) {
    fun calculate(name: String): Double {
        val job = jobs[name]!!
        val parts = job.split(" ")

        if (parts.size == 1) {
            return parts[0].toDouble()
        }

        val (left, op, right) = parts

        return when (op) {
            "+" -> calculate(left) + calculate(right)
            "-" -> calculate(left) - calculate(right)
            "*" -> calculate(left) * calculate(right)
            "/" -> calculate(left) / calculate(right)
            else -> error("unknown operator")
        }
    }
}

fun main() {
    fun part1(input: List<String>) = input.toInstructions()
        .let { Monkeys(it).calculate("root") }
        .toLong()

    fun part2(input: List<String>): Long {
        val instructions = input.toInstructions().toMutableMap()

        val (left, _, right) = instructions["root"]!!.split(" ")
        instructions["root"] = "$left - $right"

        val initial = Monkeys(instructions).calculate("root")

        var min = 0L
        var max = Long.MAX_VALUE

        while (true) {
            val mid = min checkedAdd ((max checkedSub min) floorDiv 2L)
            instructions["humn"] = mid.toString()

            val root = Monkeys(instructions).calculate("root")

            if (root == 0.0) {
                return mid
            }

            if (root.sign == initial.sign) {
                min = mid checkedAdd 1
            } else {
                max = mid
            }
        }

        error("failed to find a solution")
    }

    val testInput = readInput("day21/Day21_test")
    check(part1(testInput) == 152L)
    check(part2(testInput) == 301L)

    val input = readInput("day21/Day21")
    println(part1(input))
    println(part2(input))
}
