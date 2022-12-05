use anyhow::{bail, Result};
use parse_display::{Display, FromStr};
use std::collections::HashSet;

#[derive(Clone, Copy, Debug, Display, FromStr, PartialEq)]
enum Sign {
    #[display("+")]
    Positive,
    #[display("-")]
    Negative,
}

#[derive(Clone, Copy, Debug, Display, FromStr, PartialEq)]
#[display(style = "lowercase")]
enum Op {
    NOP,
    ACC,
    JMP,
}

#[derive(Clone, Copy, Debug, FromStr, PartialEq)]
#[from_str(regex = "(?P<op>[a-z]+) (?P<sign>[+-])(?P<value>[0-9]+)")]
struct Instruction {
    op: Op,
    sign: Sign,
    value: usize,
}

struct Console {
    acc: isize,
    pc: usize,
    rom: Vec<Instruction>,
}

impl Console {
    fn step(&mut self) {
        match self.rom[self.pc] {
            Instruction { op: Op::NOP, .. } => self.pc = self.pc.wrapping_add(1),
            Instruction {
                op: Op::ACC,
                sign,
                value,
            } => {
                self.acc = match sign {
                    Sign::Positive => self.acc.wrapping_add(value as isize),
                    Sign::Negative => self.acc.wrapping_sub(value as isize),
                };

                self.pc = self.pc.wrapping_add(1)
            }
            Instruction {
                op: Op::JMP,
                sign,
                value,
            } => match sign {
                Sign::Positive => self.pc = self.pc.wrapping_add(value),
                Sign::Negative => self.pc = self.pc.wrapping_sub(value),
            },
        }
    }
}

fn run_until_loop_or_end(rom: &[Instruction]) -> (isize, bool) {
    let mut console = Console {
        acc: 0,
        pc: 0,
        rom: rom.to_vec(),
    };
    let mut executed_ops = HashSet::new();

    while !executed_ops.contains(&console.pc) && console.pc < rom.len() {
        executed_ops.insert(console.pc);
        console.step();
    }

    (console.acc, console.pc >= rom.len())
}

#[aoc_generator(day8)]
fn generator(input: &str) -> Vec<Instruction> {
    input.lines().map(|line| line.parse().unwrap()).collect()
}

#[aoc(day8, part1)]
fn part1(rom: &[Instruction]) -> isize {
    let (acc, _) = run_until_loop_or_end(rom);
    acc
}

#[aoc(day8, part2)]
fn part2(rom: &[Instruction]) -> Result<isize> {
    for i in 0..rom.len() {
        let mut modified_rom: Vec<Instruction> = rom.to_vec();
        modified_rom[i].op = match modified_rom[i].op {
            Op::NOP => Op::JMP,
            Op::JMP => Op::NOP,
            _ => continue,
        };

        if let (acc, true) = run_until_loop_or_end(&modified_rom) {
            return Ok(acc);
        }
    }

    bail!("no instruction found")
}

#[cfg(test)]
mod tests {
    use super::*;
    use indoc::indoc;

    #[test]
    fn it_parses_input() {
        let data = indoc! {"
            nop +0
            acc +1
            jmp +4
            acc +3
            jmp -3
            acc -99
            acc +1
            jmp -4
            acc +6
        "};

        let input = generator(data);
        assert_eq!(
            [
                Instruction {
                    op: Op::NOP,
                    sign: Sign::Positive,
                    value: 0
                },
                Instruction {
                    op: Op::ACC,
                    sign: Sign::Positive,
                    value: 1
                },
                Instruction {
                    op: Op::JMP,
                    sign: Sign::Positive,
                    value: 4
                },
                Instruction {
                    op: Op::ACC,
                    sign: Sign::Positive,
                    value: 3
                },
                Instruction {
                    op: Op::JMP,
                    sign: Sign::Negative,
                    value: 3
                },
                Instruction {
                    op: Op::ACC,
                    sign: Sign::Negative,
                    value: 99
                },
                Instruction {
                    op: Op::ACC,
                    sign: Sign::Positive,
                    value: 1
                },
                Instruction {
                    op: Op::JMP,
                    sign: Sign::Negative,
                    value: 4
                },
                Instruction {
                    op: Op::ACC,
                    sign: Sign::Positive,
                    value: 6
                },
            ]
            .to_vec(),
            input
        );
    }

    #[test]
    fn it_solves_part1() {
        let data = indoc! {"
            nop +0
            acc +1
            jmp +4
            acc +3
            jmp -3
            acc -99
            acc +1
            jmp -4
            acc +6
        "};

        let input = generator(data);
        assert_eq!(5, part1(&input));
    }

    #[test]
    fn it_solves_part2() {
        let data = indoc! {"
            nop +0
            acc +1
            jmp +4
            acc +3
            jmp -3
            acc -99
            acc +1
            jmp -4
            acc +6
        "};

        let input = generator(data);
        assert_eq!(8, part2(&input).unwrap());
    }
}
