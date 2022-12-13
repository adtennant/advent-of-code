package dev.adtennant.adventofcode.day13

import kotlinx.serialization.decodeFromString
import kotlinx.serialization.json.Json
import kotlinx.serialization.json.JsonArray
import kotlinx.serialization.json.JsonElement
import kotlinx.serialization.json.JsonPrimitive
import readInput

sealed interface Packet : Comparable<Packet> {
    data class IntValue(val value: Int) : Packet {
        override fun compareTo(other: Packet) = when (other) {
            is IntValue -> value compareTo other.value
            is ListValue -> ListValue(listOf(this)) compareTo other
        }
    }

    data class ListValue(val value: List<Packet>) : Packet {
        override fun compareTo(other: Packet) = when (other) {
            is IntValue -> this compareTo ListValue(listOf(other))
            is ListValue -> {
                repeat(minOf(value.size, other.value.size)) { i ->
                    val result = value[i] compareTo other.value[i]

                    if (result != 0) {
                        // Fall through to list size if all values the same
                        return result
                    }
                }

                value.size compareTo other.value.size
            }
        }
    }
}

fun JsonElement.toPackets(): Packet = when (this) {
    is JsonPrimitive -> content.toInt().let { Packet.IntValue(it) }
    is JsonArray -> map { it.toPackets() }.let { Packet.ListValue(it) }
    else -> error("invalid json")
}

fun String.toPacket() = Json.decodeFromString<JsonArray>(this).toPackets()
fun List<String>.toPackets() = filter(String::isNotEmpty).map(String::toPacket)

fun main() {
    fun part1(input: List<String>) = input.toPackets()
        .chunked(2) { (left, right) -> left compareTo right }
        .withIndex()
        .filter { it.value < 0 }
        .sumOf { it.index + 1 }

    fun part2(input: List<String>): Int {
        val firstDivider = "[[2]]".toPacket()
        val secondDivider = "[[6]]".toPacket()

        val packets = listOf(firstDivider, secondDivider) + input.toPackets()
        val organized = packets.sorted()

        val firstIndex = organized.indexOf(firstDivider)
        val secondIndex = organized.indexOf(secondDivider)
        return (firstIndex + 1) * (secondIndex + 1)
    }

    val testInput = readInput("day13/Day13_test")
    check(part1(testInput) == 13)
    check(part2(testInput) == 140)

    val input = readInput("day13/Day13")
    println(part1(input))
    println(part2(input))
}
