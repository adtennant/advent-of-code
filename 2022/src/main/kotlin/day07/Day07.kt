package day07

import readInput

fun String.isChangeDirectory() = startsWith("$ cd")
fun String.isFile() = first().isDigit()

fun allSubDirs(path: String): List<String> {
    val parts = path.count { it == '/' };
    return (1..parts).map {
        path.split("/")
            .drop(1) // Skip the initial /
            .take(it)
            .joinToString("/", prefix = "/")
    }
}

fun pathForIndex(input: List<String>, index: Int) =
    input.take(index + 1)
        .filter(String::isChangeDirectory)
        .fold(emptyList<String>()) { acc, command ->
            when (val dir = command.split(" ")[2]) {
                "/" -> emptyList()
                ".." -> acc.dropLast(1)
                else -> acc + dir
            }
        }.joinToString("/", prefix = "/")

data class File(val path: String, val name: String, val size: Int)

fun files(input: List<String>) =
    input.mapIndexed { i, value ->
        when {
            value.isFile() -> {
                val path = pathForIndex(input, i)
                val name = value.split(" ")[1];
                val size = value.split(" ")[0].toInt()
                File(path, name, size)
            }
            else -> null
        }
    }.filterNotNull()

data class Directory(val path: String, val size: Int)

fun directories(files: List<File>) = files
    .flatMap { allSubDirs(it.path) }
    .distinct()
    .map { path ->
        val size = files
            .filter { it.path.startsWith(path) }
            .sumOf(File::size)
        Directory(path, size)
    }

fun main() {
    fun part1(input: List<String>) = directories(files(input))
        .filter { it.size <= 100000 }
        .sumOf(Directory::size)

    fun part2(input: List<String>): Int {
        val directories = directories(files(input))

        val free = 70000000 - directories.maxOf(Directory::size)
        val required = 30000000 - free

        return directories
            .filter { it.size >= required }
            .minOf(Directory::size)
    }

    val testInput = readInput("day07/Day07_test")
    check(part1(testInput) == 95437)
    check(part2(testInput) == 24933642)

    val input = readInput("day07/Day07")
    println(part1(input))
    println(part2(input))
}
