package day19

import readInput

data class ObsidianRobotCost(val ore: Int, val clay: Int)
data class GeodeRobotCost(val ore: Int, val obsidian: Int)

data class Blueprint(
    val id: Int,
    val oreRobotCost: Int,
    val clayRobotCost: Int,
    val obsidianRobotCost: ObsidianRobotCost,
    val geodeRobotCost: GeodeRobotCost
) {
    companion object {
        val regex =
            """Blueprint (\d+): Each ore robot costs (\d+) ore\. Each clay robot costs (\d+) ore\. Each obsidian robot costs (\d+) ore and (\d+) clay\. Each geode robot costs (\d+) ore and (\d+) obsidian\.""".toRegex()
    }
}

fun String.toBlueprint() = Blueprint.regex.find(this)!!
    .let(MatchResult::groupValues)
    .drop(1)
    .map(String::toInt)
    .let {
        Blueprint(
            it[0],
            it[1],
            it[2],
            ObsidianRobotCost(it[3], it[4]),
            GeodeRobotCost(it[5], it[6])
        )
    }

data class State(
    val oreRobots: Int,
    val clayRobots: Int,
    val obsidianRobots: Int,
    val geodeRobots: Int,
    val ore: Int,
    val clay: Int,
    val obsidian: Int,
    val geodes: Int,
    val remainingTime: Int
) {
    fun advance() = copy(
        ore = ore + oreRobots,
        clay = clay + clayRobots,
        obsidian = obsidian + obsidianRobots,
        geodes = geodes + geodeRobots,
        remainingTime = remainingTime - 1
    )

    fun buildOreRobot(blueprint: Blueprint) = copy(
        oreRobots = oreRobots + 1,
        ore = ore - blueprint.oreRobotCost,
    )

    fun buildClayRobot(blueprint: Blueprint) = copy(
        clayRobots = clayRobots + 1,
        ore = ore - blueprint.clayRobotCost,
    )

    fun buildObsidianRobot(blueprint: Blueprint) = copy(
        obsidianRobots = obsidianRobots + 1,
        ore = ore - blueprint.obsidianRobotCost.ore,
        clay = clay - blueprint.obsidianRobotCost.clay
    )

    fun buildGeodeRobot(blueprint: Blueprint) = copy(
        geodeRobots = geodeRobots + 1,
        ore = ore - blueprint.geodeRobotCost.ore,
        obsidian = obsidian - blueprint.geodeRobotCost.obsidian
    )
}

fun findMaxGeodes(blueprint: Blueprint, time: Int): Int {
    val maxOreCost =
        maxOf(
            blueprint.oreRobotCost,
            blueprint.clayRobotCost,
            blueprint.obsidianRobotCost.ore,
            blueprint.geodeRobotCost.ore
        )
    val initialState = State(1, 0, 0, 0, 0, 0, 0, 0, time)
    val queue = ArrayDeque(listOf(initialState))
    val seen = mutableSetOf<State>()
    var maxGeodes = 0

    while (queue.isNotEmpty()) {
        val state = queue.removeFirst()

        maxGeodes = maxOf(maxGeodes, state.geodes)

        if (state.remainingTime == 0) {
            continue
        }

        // Ignore excess materials, those that don't contribute to advancing the state in a meaningful way
        val ore = minOf(state.ore, maxOreCost + (maxOreCost - state.oreRobots) * (state.remainingTime - 1))
        val clay = minOf(
            state.clay,
            blueprint.obsidianRobotCost.clay + (blueprint.obsidianRobotCost.clay - state.clayRobots) * (state.remainingTime - 1)
        )
        val obsidian = minOf(
            state.obsidian,
            blueprint.geodeRobotCost.obsidian + (blueprint.geodeRobotCost.obsidian - state.obsidianRobots) * (state.remainingTime - 1)
        )

        // Don't build excess robots
        val oreRobots = minOf(state.oreRobots, maxOreCost)
        val clayRobots = minOf(state.clayRobots, blueprint.obsidianRobotCost.clay)
        val obsidianRobots = minOf(state.obsidianRobots, blueprint.geodeRobotCost.obsidian)

        val next = state.copy(
            ore = ore,
            clay = clay,
            obsidian = obsidian,
            oreRobots = oreRobots,
            clayRobots = clayRobots,
            obsidianRobots = obsidianRobots
        )

        if (seen.contains(next)) {
            continue
        }

        seen.add(next)

        if (next.ore >= blueprint.geodeRobotCost.ore && next.obsidian >= blueprint.geodeRobotCost.obsidian) {
            // Assuming that building a geode robot is always the best thing to do
            queue.add(next.advance().buildGeodeRobot(blueprint))
        } else {
            if (next.ore >= blueprint.obsidianRobotCost.ore && next.clay >= blueprint.obsidianRobotCost.clay && next.obsidianRobots < blueprint.geodeRobotCost.obsidian) {
                queue.add(next.advance().buildObsidianRobot(blueprint))
            }

            if (next.ore >= blueprint.clayRobotCost && next.clayRobots < blueprint.obsidianRobotCost.clay) {
                queue.add(next.advance().buildClayRobot(blueprint))
            }

            if (next.ore >= blueprint.oreRobotCost && next.oreRobots < maxOreCost) {
                queue.add(next.advance().buildOreRobot(blueprint))
            }
        }

        queue.add(next.advance())
    }

    return maxGeodes
}

fun main() {
    fun part1(input: List<String>) =
        input.map(String::toBlueprint)
            .sumOf { it.id * findMaxGeodes(it, 24) }

    fun part2(input: List<String>) =
        input.map(String::toBlueprint)
            .take(3)
            .map { findMaxGeodes(it, 32) }
            .fold(1) { acc, i -> acc * i }

    val testInput = readInput("day19/Day19_test")
    check(part1(testInput) == 33)
    check(part2(testInput) == 3472)

    val input = readInput("day19/Day19")
    println(part1(input))
    println(part2(input))
}
