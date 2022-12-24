package day11

import checkedAdd
import checkedMul
import checkedSub
import floorDiv
import readInput
import split

data class Monkey(
    val items: List<Long>,
    val operation: (Long) -> Long,
    val test: Long,
    val success: Int,
    val failure: Int,
    val business: Long = 0L
) {
    fun takeTurn(monkeys: List<Monkey>, worryTransform: (Long) -> Long): List<Monkey> {
        val targets = items.map {
            val worry = worryTransform(operation(it))
            val target = if (worry % test == 0L) {
                success
            } else {
                failure
            }
            IndexedValue(target, worry)
        }

        return monkeys.withIndex().map { (index, monkey) ->
            if (monkey == this) {
                monkey.copy(items = emptyList(), business = business + items.size)
            } else {
                val received = targets.filter { it.index == index }.map { it.value }
                monkey.copy(items = monkey.items + received)
            }
        }
    }
}

fun List<Monkey>.doRound(worryTransform: (Long) -> Long) =
    foldIndexed(this) { i, acc, _ -> acc[i].takeTurn(acc, worryTransform) }

fun List<Monkey>.doRounds(rounds: Int, worryTransform: (Long) -> Long) = (0 until rounds)
    .fold(this) { acc, _ -> acc.doRound(worryTransform) }

fun List<Monkey>.business() = sortedByDescending { it.business }
    .take(2)
    .let { it[0].business checkedMul it[1].business }

fun String.toOperation(): (Long) -> Long {
    val parts = split(" ")
    val operator = parts[3]

    return when (val value = parts[4]) {
        "old" -> when (operator) {
            "*" -> { old -> old checkedMul old }
            else -> error("invalid operation")
        }
        else -> when (operator) {
            "+" -> { old -> old checkedAdd value.toLong() }
            "-" -> { old -> old checkedSub value.toLong() }
            "*" -> { old -> old checkedMul value.toLong() }
            "/" -> { old -> old floorDiv value.toLong() }
            else -> error("invalid operation")
        }
    }
}

fun List<String>.toMonkey(): Monkey {
    val items = this[1].substringAfter("Starting items: ").split(", ").map(String::toLong)
    val operation = this[2].substringAfter("Operation: ").toOperation()
    val test = this[3].substringAfter("Test: divisible by ").toLong()
    val success = this[4].substringAfter("If true: throw to monkey ").toInt()
    val failure = this[5].substringAfter("If false: throw to monkey ").toInt()

    return Monkey(items, operation, test, success, failure)
}

fun main() {
    fun parseMonkeys(input: List<String>) = input.split(String::isEmpty)
        .map(List<String>::toMonkey)

    fun part1(input: List<String>) = parseMonkeys(input)
        .doRounds(20) { worry -> worry / 3 }
        .business()

    fun part2(input: List<String>): Long {
        val monkeys = parseMonkeys(input)
        val divisor = monkeys.fold(1L) { acc, monkey -> acc checkedMul monkey.test }

        return monkeys.doRounds(10000) { worry -> worry % divisor }
            .business()
    }

    val testInput = readInput("day11/Day11_test")
    check(part1(testInput) == 10605L)
    check(part2(testInput) == 2713310158)

    val input = readInput("day11/Day11")
    println(part1(input))
    println(part2(input))
}
