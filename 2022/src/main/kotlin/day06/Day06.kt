package day06

import readInputAsText

fun main() {
    fun findMarker(input: String, length: Int) =
        input.windowed(length, 1)
            .takeWhile { it.length != it.toSet().size }
            .count() + length

    fun part1(input: String) = findMarker(input, 4)

    fun part2(input: String) = findMarker(input, 14)

    check(part1("bvwbjplbgvbhsrlpgdmjqwftvncz") == 5)
    check(part1("nppdvjthqldpwncqszvftbrmjlhg") == 6)
    check(part1("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg") == 10)
    check(part1("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw") == 11)

    check(part2("mjqjpqmgbljsphdztnvjfqwrcgsmlb") == 19)
    check(part2("bvwbjplbgvbhsrlpgdmjqwftvncz") == 23)
    check(part2("nppdvjthqldpwncqszvftbrmjlhg") == 23)
    check(part2("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg") == 29)
    check(part2("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw") == 26)

    val input = readInputAsText("day06/Day06")
    println(part1(input))
    println(part2(input))
}
