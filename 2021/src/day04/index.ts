import run from "aocrunner";
import { parseDecimal, parseInput } from "../utils/index.js";

type Cell = {
  value: number;
  marked: boolean;
};

class Board {
  constructor(public cells: ReadonlyArray<ReadonlyArray<Cell>>) {}

  hasWon = () => {
    const rowWin = this.cells.some((row) => row.every((cell) => cell.marked));
    const colWin = this.cells[0].some((_, col) =>
      this.cells.every((row) => row[col].marked)
    );

    return rowWin || colWin;
  };

  mark = (value: number) =>
    new Board(
      this.cells.map((row) =>
        row.map((cell) => ({
          ...cell,
          marked: cell.marked || cell.value == value,
        }))
      )
    );
}

type Result = {
  currentNumber: number;
  board: Board;
};

const score = ({ currentNumber, board }: Result) =>
  currentNumber *
  board.cells
    .flatMap((row) => row.filter((cell) => !cell.marked))
    .reduce((result, cell) => result + cell.value, 0);

const parseNumbers = (input: string) => input.split(",").map(parseDecimal);
const parseBoard = (input: string): Board =>
  new Board(
    input.lines().map((line) =>
      line
        .trim()
        .split(/\s+/)
        .map((cell) => cell.trim())
        .map((cell) => ({ value: parseDecimal(cell), marked: false }))
    )
  );

const getScore = (
  input: string,
  predicate: (result: { boards: Board[] }) => boolean
) => {
  const [rawNumbers, ...rawBoards] = input.split("\n\n");
  const { numbers, boards } = {
    numbers: rawNumbers.map(parseNumbers),
    boards: rawBoards.map(parseBoard),
  };

  const result = numbers
    .reduceWhile(
      ({ boards }, number) => ({
        currentNumber: number,
        boards: boards
          .filter((board) => !board.hasWon())
          .map((board) => board.mark(number)),
      }),
      predicate,
      { currentNumber: -1, boards }
    )
    .map(({ currentNumber, boards }) => ({
      currentNumber,
      board: boards.filter((board) => board.hasWon()).first(),
    }));

  return score(result);
};

const part1 = (rawInput: string) => {
  const input = parseInput(rawInput);
  return getScore(input, ({ boards }) =>
    boards.some((board) => board.hasWon())
  );
};

const part2 = (rawInput: string) => {
  const input = parseInput(rawInput);

  return getScore(
    input,
    ({ boards }) => boards.length == 1 && boards.first().hasWon()
  );
};

const exampleInput = `7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7`;

run({
  part1: {
    tests: [
      {
        input: exampleInput,
        expected: 4512,
      },
    ],
    solution: part1,
  },
  part2: {
    tests: [
      {
        input: exampleInput,
        expected: 1924,
      },
    ],
    solution: part2,
  },
  trimTestInputs: true,
  onlyTests: false,
});
