import java.io.File
import java.math.BigInteger
import java.security.MessageDigest

/**
 * Reads lines from the given input txt file.
 */
fun readInput(name: String) = File("src/dev/adtennant/adventofcode", "$name.txt")
    .readLines()

/**
 * Reads text from the given input txt file.
 */
fun readInputAsText(name: String) = File("src/dev/adtennant/adventofcode", "$name.txt")
    .readText()

/**
 * Converts string to md5 hash.
 */
fun String.md5() = BigInteger(1, MessageDigest.getInstance("MD5").digest(toByteArray()))
    .toString(16)
    .padStart(32, '0')

/**
 * Takes exactly n elements from a Collection. Throws if n != size.
 */
fun <T> Iterable<T>.takeExact(n: Int): List<T> = if (count() == n) {
    toList()
} else {
    error("collection is not $n elements in size")
}

/**
 * Splits this iterable into a prefix/suffix pair according to a predicate.
 */
fun <T> Iterable<T>.span(predicate: (T) -> Boolean): Pair<List<T>, List<T>> =
    Pair(takeWhile(predicate), dropWhile(predicate))

/**
 * Splits this iterable into a prefix/suffix pair according to a predicate. Removes the entry matching the predicate.
 */
fun <T> Iterable<T>.split(predicate: (T) -> Boolean): List<List<T>> {
    val result = mutableListOf<List<T>>()
    var remainder = this

    while (remainder.toList().isNotEmpty()) {
        val (start, end) = remainder.span(predicate)
        result.add(start)

        remainder = end.drop(1)
    }

    return result
}
