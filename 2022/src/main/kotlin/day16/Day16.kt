package day16

import readInput
import java.util.PriorityQueue

data class Node(val name: String, val rate: Int)

typealias Graph = Map<Node, Map<Node, Int>>

fun Graph.start(): Node = keys.find { it.name == "AA" }!!

typealias Route = List<Node>

val nodeRegex = """Valve (\S+) has flow rate=(\d+); tunnels? leads? to valves? (.+)$""".toRegex()

fun graph(input: List<String>): Graph {
    val nodes = input.fold(mutableMapOf<String, Node>()) { nodes, line ->
        val (name, rate, _) = nodeRegex.find(line)!!.destructured
        nodes[name] = Node(name, rate.toInt())
        nodes
    }

    val graph = input.fold(mutableMapOf<Node, Map<Node, Int>>()) { graph, line ->
        val (name, _, valves) = nodeRegex.find(line)!!.destructured
        val node = nodes[name]!!
        val edges = valves.split(", ")
            .fold(mutableMapOf<Node, Int>()) { edges, valve ->
                edges[nodes[valve]!!] = 1
                edges
            }

        graph[node] = edges
        graph
    }

    return graph
}

fun findPaths(graph: Graph, goal: Node): Map<Node, Int> {
    val q = PriorityQueue<Pair<Int, Node>>(compareBy { it.first })
    q.add(Pair(0, goal))

    val paths = mutableMapOf(goal to 0)

    while (q.isNotEmpty()) {
        val (cost, current) = q.remove()

        for ((next, nextCost) in graph[current]!!) {
            if (!paths.contains(next) || cost + nextCost < paths[next]!!) {
                paths[next] = cost + nextCost
                q.add(Pair(cost + nextCost, next))
            }
        }
    }

    return paths
}

fun findAllPaths(graph: Graph, start: Node): Map<Node, Map<Node, Int>> {
    return graph.keys
        .filter { node -> node == start || node.rate > 0 }
        .fold(mutableMapOf()) { acc, node ->
            acc[node] = findPaths(graph, node)
            acc
        }
}

fun allRoutes(
    from: Node,
    paths: Map<Node, Map<Node, Int>>,
    time: Int,
    remaining: Map<Node, Map<Node, Int>> = paths - from,
    searched: List<Node> = emptyList(),
): List<Route> {
    return buildList {
        for (next in remaining.keys) {
            val cost = paths[from]!![next]!! + 1

            if (cost < time) {
                addAll(allRoutes(next, paths, time - cost, remaining - next, (searched + next)))
            }
        }

        add(searched)
    }
}

fun totalPressure(route: Route, start: Node, time: Int, paths: Map<Node, Map<Node, Int>>): Int {
    var result = 0
    var current = start
    var t = time

    for (node in route) {
        val cost = paths[current]!![node]!! + 1
        t -= cost

        result += t * node.rate
        current = node
    }

    return result
}

fun main() {
    fun part1(input: List<String>): Int {
        val graph = graph(input)
        val start = graph.start()

        val paths = findAllPaths(graph, start)
        val allRoutes = allRoutes(start, paths, 30)

        return allRoutes.maxOf { totalPressure(it, start, 30, paths) }
    }

    fun part2(input: List<String>): Int {
        val graph = graph(input)
        val start = graph.start()

        val paths = findAllPaths(graph, start)
        val allRoutes = allRoutes(start, paths, 26)

        val allResults = allRoutes
            .map { Pair(totalPressure(it, start, 26, paths), it.toSet()) }
            .sortedByDescending { it.first } // Sort to speed things up

        var highestResult = 0

        for ((i, value) in allResults.withIndex()) {
            val (result, route) = value

            if (result * 2 < highestResult) {
                // Early exit as highest cannot be beaten
                break
            }

            for ((elephantResult, elephantRoute) in allResults.drop(i + 1)) {
                // Only care about routes where the elephant turns different valves
                if (route.intersect(elephantRoute).isEmpty()) {
                    val totalScore = result + elephantResult

                    if (totalScore > highestResult) {
                        highestResult = result
                    }
                }
            }
        }

        return highestResult
    }

    val testInput = readInput("day16/Day16_test")
    check(part1(testInput) == 1651)
    check(part2(testInput) == 1707)

    val input = readInput("day16/Day16")
    println(part1(input))
    println(part2(input))
}
