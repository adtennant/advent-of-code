import run from "aocrunner";
import { parseInput } from "../utils/index.js";

type Point = {
  x: number;
  y: number;
};

type Line = {
  start: Point;
  end: Point;
};

const solve = (rawInput: string, lineFilter: (line: Line) => boolean) => {
  const input = parseInput(rawInput);
  const lines: Line[] = input
    .lines()
    .map((line) => /(\d*),(\d*) -> (\d*),(\d*)/.exec(line))
    .filter((match): match is RegExpExecArray => match !== null)
    .map(([, startX, startY, endX, endY]) => ({
      start: { x: parseInt(startX, 10), y: parseInt(startY, 10) },
      end: { x: parseInt(endX, 10), y: parseInt(endY, 10) },
    }));

  const grid: Record<string, number> = lines
    .filter(lineFilter)
    .reduce((grid, line) => {
      const { start, end } = line;

      const minX = Math.min(start.x, end.x);
      const maxX = Math.max(start.x, end.x);

      const minY = Math.min(start.y, end.y);
      const maxY = Math.max(start.y, end.y);

      const length = Math.max(maxX - minX, maxY - minY);

      for (let i = 0; i <= length; i++) {
        const x = start.x + (end.x - start.x) * (i / length);
        const y = start.y + (end.y - start.y) * (i / length);

        grid[`${Math.round(x)},${Math.round(y)}`] =
          (grid[`${Math.round(x)},${Math.round(y)}`] || 0) + 1;
      }

      return grid;
    }, {} as Record<string, number>);

  return Object.values(grid).filter((value) => value > 1).length;
};

const part1 = (rawInput: string) => {
  return solve(
    rawInput,
    (line) => line.start.x == line.end.x || line.start.y == line.end.y
  );
};

const part2 = (rawInput: string) => {
  return solve(rawInput, (line) => true);
};

const exampleInput = `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`;

run({
  part1: {
    tests: [
      {
        input: exampleInput,
        expected: 5,
      },
    ],
    solution: part1,
  },
  part2: {
    tests: [
      {
        input: exampleInput,
        expected: 12,
      },
    ],
    solution: part2,
  },
  trimTestInputs: true,
  onlyTests: false,
});
