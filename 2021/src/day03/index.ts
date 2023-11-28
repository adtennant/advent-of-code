import run from "aocrunner";
import { parseBinary, parseInput } from "../utils/index.js";

const findMostCommonBit = (input: string[], i: number) => {
  const onesCount = input.filter((value) => value[i] === "1").length;
  return onesCount >= input.length / 2 ? "1" : "0";
};

const part1 = (rawInput: string) => {
  const input = parseInput(rawInput);

  const lines = input.lines();
  const bitCount = lines.first().length;

  const gamma = Array.from({ length: bitCount })
    .map((_, i) => i)
    .reduce((result, i) => result + findMostCommonBit(lines, i), "")
    .map(parseBinary);

  const mask = "1".repeat(bitCount).map(parseBinary);
  const epsilon = gamma ^ mask;

  return gamma * epsilon;
};

const part2 = (rawInput: string) => {
  const input = parseInput(rawInput);

  const lines = input.lines();
  const bitCount = lines[0].length;

  const oxygen = Array.from({ length: bitCount })
    .map((_, i) => i)
    .reduceWhile(
      (result, i) =>
        result.filter((value) => value[i] === findMostCommonBit(result, i)),
      (result) => result.length == 1,
      lines
    )
    .first()
    .map(parseBinary);

  const co2 = Array.from({ length: bitCount })
    .map((_, i) => i)
    .reduceWhile(
      (result, i) =>
        result.filter((value) => value[i] !== findMostCommonBit(result, i)),
      (result) => result.length == 1,
      lines
    )
    .first()
    .map(parseBinary);

  return oxygen * co2;
};

const exampleInput = `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`;

run({
  part1: {
    tests: [
      {
        input: exampleInput,
        expected: 198,
      },
    ],
    solution: part1,
  },
  part2: {
    tests: [
      {
        input: exampleInput,
        expected: 230,
      },
    ],
    solution: part2,
  },
  trimTestInputs: true,
  onlyTests: false,
});
