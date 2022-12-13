package day10

import readInput
import kotlin.math.abs

sealed class Instruction {
    object Noop : Instruction()
    data class AddX(val value: Int) : Instruction()
}

fun String.toInstruction() = when (substringBefore(" ")) {
    "noop" -> Instruction.Noop
    "addx" -> Instruction.AddX(split(" ")[1].toInt())
    else -> error("invalid instruction")
}

data class Cpu(val x: Int = 1, val cycle: Int = 1) {
    val signalStrength get() = x * cycle

    private fun execute(instruction: Instruction) = when (instruction) {
        is Instruction.Noop -> listOf(Cpu(x, cycle + 1))
        is Instruction.AddX -> listOf(
            Cpu(x, cycle + 1),
            Cpu(x + instruction.value, cycle + 2)
        )
    }

    fun run(instructions: List<Instruction>): List<Cpu> {
        return instructions
            .fold(listOf(this)) { acc, instruction ->
                val prev = acc.last()
                acc + prev.execute(instruction)
            }
    }
}

fun main() {
    fun part1(input: List<String>) =
        input.map(String::toInstruction).let { Cpu().run(it) }
            .filter { cpu -> (cpu.cycle + 20) % 40 == 0 }
            .sumOf { it.signalStrength }

    fun part2(input: List<String>) = input.map(String::toInstruction).let { Cpu().run(it) }
        .fold(emptyList<Char>()) { screen, cpu ->
            val pixel = if (abs(((cpu.cycle - 1) % 40) - cpu.x) <= 1) {
                '#'
            } else {
                '.'
            }
            screen + pixel
        }
        .windowed(40, 40)
        .joinToString("\n") {
            it.joinToString("")
        }

    val testInput = readInput("day10/Day10_test")
    check(part1(testInput) == 13140)
    check(
        part2(testInput) ==
                "##..##..##..##..##..##..##..##..##..##..\n" +
                "###...###...###...###...###...###...###.\n" +
                "####....####....####....####....####....\n" +
                "#####.....#####.....#####.....#####.....\n" +
                "######......######......######......####\n" +
                "#######.......#######.......#######....."
    )

    val input = readInput("day10/Day10")
    println(part1(input))
    println(part2(input))
}
