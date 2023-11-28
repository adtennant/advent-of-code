import run from "aocrunner";
import { parseInput } from "../utils/index.js";

const solve = (rawInput: string, days: number) => {
  const input = parseInput(rawInput);
  const fishByAge = input
    .split(",")
    .map((value) => parseInt(value, 10))
    .reduce(
      (result, value) => {
        result[value]++;
        return result;
      },
      [0, 0, 0, 0, 0, 0, 0, 0]
    );

  return Array<number>(days)
    .fill(0)
    .reduce((result) => {
      const toReset = result[0];

      for (let j = 0; j < 8; j++) {
        result[j] = result[j + 1] || 0;
      }

      result[8] = toReset;
      result[6] = result[6] + toReset;

      return result;
    }, fishByAge)
    .reduce((result, value) => result + value, 0);
};

const part1 = (rawInput: string) => {
  return solve(rawInput, 80);
};

const part2 = (rawInput: string) => {
  return solve(rawInput, 256);
};

const exampleInput = `3,4,3,1,2`;

run({
  part1: {
    tests: [
      {
        input: exampleInput,
        expected: 5934,
      },
    ],
    solution: part1,
  },
  part2: {
    tests: [
      {
        input: exampleInput,
        expected: 26984457539,
      },
    ],
    solution: part2,
  },
  trimTestInputs: true,
  onlyTests: false,
});
