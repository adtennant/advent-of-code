use anyhow::{bail, Result};
use itertools::Itertools;
use std::str::FromStr;

#[derive(Debug, PartialEq)]
struct Seat {
    row: u8,
    column: u8,
}

impl FromStr for Seat {
    type Err = anyhow::Error;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let parts = s.split_at(7);

        let row = parts.0.chars().try_fold(0u8, |result, c| match c {
            'F' => Ok(result.wrapping_shl(1) | 0),
            'B' => Ok(result.wrapping_shl(1) | 1),
            _ => bail!("invalid character: {}", c),
        })?;

        let column = parts.1.chars().try_fold(0u8, |result, c| match c {
            'L' => Ok(result.wrapping_shl(1) | 0),
            'R' => Ok(result.wrapping_shl(1) | 1),
            _ => bail!("invalid character in column: {}", c),
        })?;

        Ok(Seat { row, column })
    }
}

impl Seat {
    fn id(&self) -> u16 {
        u16::from(self.row) * 8 + u16::from(self.column)
    }
}

#[aoc_generator(day5)]
fn generator(data: &str) -> Result<Vec<Seat>> {
    data.lines()
        .map(|line| line.trim())
        .filter(|line| !line.is_empty())
        .map(Seat::from_str)
        .collect()
}

#[aoc(day5, part1)]
fn part1(input: &[Seat]) -> Option<u16> {
    input.iter().map(|s| s.id()).max()
}

#[aoc(day5, part2)]
fn part2(input: &[Seat]) -> Option<u16> {
    input
        .iter()
        .map(|s| s.id())
        .sorted()
        .tuple_windows::<(_, _)>()
        .find_map(|(s1, s2)| if s1 + 2 == s2 { Some(s1 + 1) } else { None })
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_parses_a_seat() {
        let data = "BFFFBBFRRR";
        let seat = Seat::from_str(data).unwrap();

        assert_eq!((70, 7, 567), (seat.row, seat.column, seat.id()));

        let data = "FFFBBBFRRR";
        let seat = Seat::from_str(data).unwrap();

        assert_eq!((14, 7, 119), (seat.row, seat.column, seat.id()));

        let data = "BBFFBBFRLL";
        let seat = Seat::from_str(data).unwrap();

        assert_eq!((102, 4, 820), (seat.row, seat.column, seat.id()))
    }
}
