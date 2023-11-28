import run from "aocrunner";
import { parseInput } from "../utils/index.js";

const solve = (rawInput: string, fuel: (distance: number) => number) => {
  const input = parseInput(rawInput);
  const crabs = input.split(",").map((value) => parseInt(value, 10));

  const min = Math.min(...crabs);
  const max = Math.max(...crabs);

  return Math.min(
    ...Array(max - min)
      .fill(0)
      .map((_, i) => i + min)
      .map((i) =>
        crabs.reduce((result, value) => result + fuel(Math.abs(value - i)), 0)
      )
  );
};

const part1 = (rawInput: string) => {
  return solve(rawInput, (distance) => distance);
};

const part2 = (rawInput: string) => {
  return solve(rawInput, (distance) =>
    Array(distance)
      .fill(0)
      .reduce((result, _, i) => result + i + 1, 0)
  );
};

const exampleInput = `16,1,2,0,4,2,7,1,2,14`;

run({
  part1: {
    tests: [
      {
        input: exampleInput,
        expected: 37,
      },
    ],
    solution: part1,
  },
  part2: {
    tests: [
      {
        input: exampleInput,
        expected: 168,
      },
    ],
    solution: part2,
  },
  trimTestInputs: true,
  onlyTests: false,
});
