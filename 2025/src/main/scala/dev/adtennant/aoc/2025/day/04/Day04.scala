package dev.adtennant.aoc.`2025`.day.`04`

import scala.io.Source._
import dev.adtennant.aoc.`2025`.Solution
import scala.util.matching.Regex
import scala.annotation.tailrec

type Cell = Option[Unit]

trait Grid[T]:
  var rows: Seq[Seq[T]]

  def height: Int = rows.length
  def width: Int = if rows.isEmpty then 0 else rows.head.length

  def get(x: Int, y: Int): T = rows(y)(x)

  def neighbors(x: Int, y: Int): Seq[(Int, Int)] =
    Seq(
      (x - 1, y - 1),
      (x, y - 1),
      (x + 1, y - 1),
      (x - 1, y),
      (x + 1, y),
      (x - 1, y + 1),
      (x, y + 1),
      (x + 1, y + 1)
    ).filter { case (nx, ny) =>
      nx >= 0 && nx < width && ny >= 0 && ny < height
    }

case class PaperGrid(var rows: Seq[Seq[Cell]]) extends Grid[Cell]:
  def removePaper(maxAdjacent: Int): PaperGrid =
    PaperGrid(
      rows.zipWithIndex
        .map { case (row, y) =>
          row.zipWithIndex.map { case (cell, x) =>
            cell
              .map { _ =>
                neighbors(x, y).count { case (nx, ny) =>
                  get(nx, ny).isDefined
                }
              }
              .flatMap(adjacent => Option.when(adjacent >= maxAdjacent)(()))
          }
        }
    )

  def removeAllPaper(maxAdjacent: Int): PaperGrid =
    val newGrid = removePaper(maxAdjacent)

    if newGrid == this then this
    else newGrid.removeAllPaper(maxAdjacent)

def parseGrid(input: String): PaperGrid =
  PaperGrid(
    input.linesIterator
      .map(
        _.split("")
          .map {
            case "@" => Some(())
            case "." => None
          }
          .toSeq
      )
      .toSeq
  )

object Day04 extends Solution[Int]:
  def day = 4

  def solvePart1(input: String): Int =
    val grid = parseGrid(input)
    val count = grid.rows.flatten.count(_.isDefined)

    val newGrid = grid.removePaper(4)
    val newCount = newGrid.rows.flatten.count(_.isDefined)

    count - newCount

  def solvePart2(input: String): Int =
    val grid = parseGrid(input)
    val count = grid.rows.flatten.count(_.isDefined)

    val newGrid = grid.removeAllPaper(4)
    val newCount = newGrid.rows.flatten.count(_.isDefined)

    count - newCount

@main
def main: Unit =
  Day04.run
