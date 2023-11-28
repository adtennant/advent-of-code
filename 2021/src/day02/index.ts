import run from "aocrunner";
import { parseDecimal, parseInput } from "../utils/index.js";

type Submarine = {
  position: number;
  depth: number;
};

type SubmarineWithAim = Submarine & {
  aim: number;
};

type Command = {
  direction: string;
  amount: number;
};

const handleCommandPart1 = (submarine: Submarine, command: Command) => {
  switch (command.direction) {
    case "forward":
      return { ...submarine, position: submarine.position + command.amount };
    case "down":
      return { ...submarine, depth: submarine.depth + command.amount };
    case "up":
      return { ...submarine, depth: submarine.depth - command.amount };
    default:
      throw new Error("invalid command");
  }
};

const handleCommandPart2 = (submarine: SubmarineWithAim, command: Command) => {
  switch (command.direction) {
    case "forward":
      return {
        ...submarine,
        position: submarine.position + command.amount,
        depth: submarine.depth + submarine.aim * command.amount,
      };
    case "down":
      return { ...submarine, aim: submarine.aim + command.amount };
    case "up":
      return { ...submarine, aim: submarine.aim - command.amount };
    default:
      throw new Error("invalid command");
  }
};

const followCourse = <T extends Submarine>(
  input: string,
  commandFn: (submarine: T, command: Command) => T,
  initialValue: T
) =>
  input
    .lines()
    .map((line) => {
      const parts = line.split(" ");
      return { direction: parts[0], amount: parts[1].map(parseDecimal) };
    })
    .reduce(commandFn, initialValue)
    .map(({ position, depth }) => position * depth);

const part1 = (rawInput: string) => {
  const input = parseInput(rawInput);

  return followCourse(input, handleCommandPart1, {
    position: 0,
    depth: 0,
  });
};

const part2 = (rawInput: string) => {
  const input = parseInput(rawInput);

  return followCourse(input, handleCommandPart2, {
    position: 0,
    depth: 0,
    aim: 0,
  });
};

const exampleInput = `forward 5
down 5
forward 8
up 3
down 8
forward 2`;

run({
  part1: {
    tests: [
      {
        input: exampleInput,
        expected: 150,
      },
    ],
    solution: part1,
  },
  part2: {
    tests: [
      {
        input: exampleInput,
        expected: 900,
      },
    ],
    solution: part2,
  },
  trimTestInputs: true,
  onlyTests: false,
});
