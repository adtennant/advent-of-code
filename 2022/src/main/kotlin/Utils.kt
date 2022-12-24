import java.io.File
import java.math.BigInteger
import java.security.MessageDigest

/**
 * Reads lines from the given input txt file.
 */
fun readInput(name: String) = File("src/main/kotlin", "$name.txt")
    .readLines()

/**
 * Reads lines from the given input txt file as text.
 */
fun readInputAsText(name: String) = File("src/main/kotlin", "$name.txt")
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
 * Splits this iterable into parts according to a predicate. Removes the entry matching the predicate.
 */
fun <T> Iterable<T>.split(predicate: (T) -> Boolean): List<List<T>> {
    val result = mutableListOf<List<T>>()
    var remainder = this

    while (remainder.toList().isNotEmpty()) {
        val (start, end) = remainder.span { i -> !predicate(i) }

        if (start.isNotEmpty()) {
            result.add(start)
        }

        remainder = end.drop(1)
    }

    return result
}

fun <T> Iterable<T>.takeWhileInclusive(predicate: (T) -> Boolean): List<T> {
    var shouldContinue = true

    return takeWhile {
        val result = shouldContinue
        shouldContinue = predicate(it)
        result
    }
}

infix fun Int.checkedAdd(other: Int): Int {
    return StrictMath.addExact(this, other)
}

infix fun Int.checkedSub(other: Int): Int {
    return StrictMath.subtractExact(this, other)
}

infix fun Int.checkedMul(other: Int): Int {
    return StrictMath.multiplyExact(this, other)
}

infix fun Int.floorDiv(other: Int): Int {
    return StrictMath.floorDiv(this, other)
}

infix fun Long.checkedAdd(other: Long): Long {
    return StrictMath.addExact(this, other)
}

infix fun Long.checkedSub(other: Long): Long {
    return StrictMath.subtractExact(this, other)
}

infix fun Long.checkedMul(other: Long): Long {
    return StrictMath.multiplyExact(this, other)
}

infix fun Long.floorDiv(other: Long): Long {
    return StrictMath.floorDiv(this, other)
}
