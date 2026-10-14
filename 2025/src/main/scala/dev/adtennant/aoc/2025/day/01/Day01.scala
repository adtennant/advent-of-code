package dev.adtennant.aoc.`2025`.day.`01`

import scala.io.Source._
import dev.adtennant.aoc.`2025`.Solution
import scala.util.matching.Regex

sealed trait Rotation:
  val amount: Int

object Rotation:
  case class Left(amount: Int) extends Rotation
  case class Right(amount: Int) extends Rotation

case class Dial(value: Int):
  def rotate(rotation: Rotation): Dial = rotation match
    case Rotation.Left(amount)  => Dial(math.floorMod(value - amount, 100))
    case Rotation.Right(amount) => Dial(math.floorMod(value + amount, 100))

def parseRotations(input: String): Seq[Rotation] =
  input.linesIterator.collect {
    case s"L$amount" => Rotation.Left(amount.toInt)
    case s"R$amount" => Rotation.Right(amount.toInt)
  }.toSeq

object Day01 extends Solution[Int]:
  def day = 1

  def solvePart1(input: String): Int =
    val rotations = parseRotations(input)
    rotations
      .foldLeft((Dial(50), 0)) { case ((dial, result), rotation) =>
        val newDial = dial.rotate(rotation)
        val newResult = result + (if newDial.value == 0 then 1 else 0)
        (newDial, newResult)
      }
      ._2

  def solvePart2(input: String): Int =
    val rotations = parseRotations(input)
    rotations
      .foldLeft((Dial(50), 0)) { case ((dial, result), rotation) =>
        val toZero = rotation match
          case Rotation.Left(_) =>
            if dial.value == 0 then 100 else dial.value
          case _ => 100 - dial.value

        val zeros = (rotation.amount - toZero) / 100;
        val atLeastOneZero = rotation.amount >= toZero;

        val newResult =
          result + (if atLeastOneZero then zeros + 1 else 0)

        val newDial = dial.rotate(rotation)
        (newDial, newResult)
      }
      ._2

@main
def main: Unit =
  Day01.run
