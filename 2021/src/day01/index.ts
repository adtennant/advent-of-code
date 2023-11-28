import run from "aocrunner";
import { parseDecimal, parseInput } from "../utils/index.js";

const part1 = (rawInput: string) => {
  const input = parseInput(rawInput);

  return input
    .lines()
    .map(parseDecimal)
    .filter((value, i, values) => i > 0 && value > values[i - 1]).length;
};

const part2 = (rawInput: string) => {
  const input = parseInput(rawInput);

  return input
    .lines()
    .map(parseDecimal)
    .tupleWindows(3)
    .map((window) => window.reduce((result, value) => result + value, 0))
    .filter((value, i, values) => i > 0 && value > values[i - 1]).length;
};

const exampleInput = `199
200
208
210
200
207
240
269
260
263`;

run({
  part1: {
    tests: [
      {
        input: exampleInput,
        expected: 7,
      },
    ],
    solution: part1,
  },
  part2: {
    tests: [
      {
        input: exampleInput,
        expected: 5,
      },
    ],
    solution: part2,
  },
  trimTestInputs: true,
  onlyTests: false,
});
