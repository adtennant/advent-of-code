package day15

import readInput
import kotlin.math.abs

typealias Point = Pair<Int, Int>

val Point.x get() = first
val Point.y get() = second

val sensorRegex = """Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)$""".toRegex()

data class Sensor(val position: Point, val closestBeacon: Point)

fun String.toSensor() = sensorRegex.find(this)!!
    .destructured
    .let { (sensorX, sensorY, beaconX, beaconY) ->
        Sensor(
            Point(sensorX.toInt(), sensorY.toInt()),
            Point(beaconX.toInt(), beaconY.toInt())
        )
    }

fun main() {
    fun part1(input: List<String>, row: Int) = input.map(String::toSensor)
        .fold(emptySet<Int>()) { acc, (sensor, beacon) ->
            val distance = abs(sensor.x - beacon.x) + abs(sensor.y - beacon.y) - abs(sensor.y - row)
            val filled = (sensor.x - distance..sensor.x + distance)
                .filter { x ->
                    // Exclude beacons
                    x != beacon.x || row != beacon.y
                }

            acc + filled
        }.size

    fun part2(input: List<String>, size: Int): Long {
        val sensors = input.map(String::toSensor)

        val rows = List(size) { row ->
            sensors.fold(emptyList<IntRange>()) { acc, (sensor, beacon) ->
                val distance = abs(sensor.x - beacon.x) + abs(sensor.y - beacon.y) - abs(sensor.y - row)
                val newRanges = buildList {
                    if (distance >= 0) {
                        add((sensor.x - distance..sensor.x + distance))
                    }
                }

                acc + newRanges
            }
        }

        val (row, col) = rows
            .map { ranges ->
                // Find the first y position on each row that could contain a beacon
                ranges
                    .sortedBy { range -> range.first }
                    .fold(0) { acc, intRange ->
                        if (intRange.contains(acc)) {
                            intRange.last + 1
                        } else {
                            acc
                        }
                    }
            }
            .withIndex()
            .first { it.value <= size }

        return col.toLong() * 4000000L + row.toLong()
    }

    val testInput = readInput("day15/Day15_test")
    check(part1(testInput, 10) == 26)
    check(part2(testInput, 20) == 56000011L)

    val input = readInput("day15/Day15")
    println(part1(input, 2000000))
    println(part2(input, 4000000))
}
