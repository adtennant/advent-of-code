import run from "aocrunner";
import { parseInput } from "../utils/index.js";

type Entry = {
  signals: string[];
  output: string[];
};

const decode = (rawInput: string) => {
  const input = parseInput(rawInput);
  const entries: Entry[] = input.lines().map((line) => {
    const parts = line.split(" | ");
    const signals = parts[0]
      .split(" ")
      .map((signal) => signal.split("").sort().join(""));
    const output = parts[1]
      .split(" ")
      .map((signal) => signal.split("").sort().join(""));

    return { signals, output };
  });

  return entries.map((entry) => {
    const one = entry.signals.find((s) => s.length === 2) as string;
    const four = entry.signals.find((s) => s.length === 4) as string;
    const seven = entry.signals.find((s) => s.length === 3) as string;
    const eight = entry.signals.find((s) => s.length === 7) as string;

    const three = entry.signals.find(
      (s) => s.length === 5 && seven.chars().every((c) => s.includes(c))
    ) as string;

    const nine = entry.signals.find(
      (s) =>
        s.length === 6 &&
        seven.chars().every((c) => s.includes(c)) &&
        three.chars().every((c) => s.includes(c))
    ) as string;

    const six = entry.signals.find(
      (s) =>
        s.length === 6 && one.chars().filter((c) => s.includes(c)).length === 1
    ) as string;
    const zero = entry.signals.find(
      (s) => s.length === 6 && s !== six && s != nine
    ) as string;
    const five = entry.signals.find(
      (s) =>
        s.length == 5 && six.chars().filter((c) => !s.includes(c)).length === 1
    ) as string;

    const two = entry.signals.find(
      (s) =>
        ![zero, one, three, four, five, six, seven, eight, nine].includes(s)
    ) as string;

    const digits = [zero, one, two, three, four, five, six, seven, eight, nine];

    return entry.output.reduce(
      (result, o, i) => result + Math.pow(10, 3 - i) * digits.indexOf(o),
      0
    );
  });
};

const part1 = (rawInput: string) => {
  return decode(rawInput).reduce(
    (result, value) =>
      result +
      value
        .toString()
        .split("")
        .filter((c) => c === "1" || c === "4" || c === "7" || c === "8").length,
    0
  );
};

const part2 = (rawInput: string) => {
  return decode(rawInput).reduce((result, value) => result + value, 0);
};

const exampleInput = `be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe
edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc
fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg
fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb
aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea
fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb
dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe
bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef
egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb
gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce`;

run({
  part1: {
    tests: [
      {
        input: exampleInput,
        expected: 26,
      },
    ],
    solution: part1,
  },
  part2: {
    tests: [
      {
        input: exampleInput,
        expected: 61229,
      },
    ],
    solution: part2,
  },
  trimTestInputs: true,
  onlyTests: false,
});
