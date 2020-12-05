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

        let row = u8::from_str_radix(
            &parts
                .0
                .chars()
                .map(|c| match c {
                    'F' => Ok('0'),
                    'B' => Ok('1'),
                    c => bail!("invalid character: {}", c),
                })
                .collect::<Result<String, _>>()?,
            2,
        )?;

        let column = u8::from_str_radix(
            &parts
                .1
                .chars()
                .map(|c| match c {
                    'L' => Ok('0'),
                    'R' => Ok('1'),
                    c => bail!("invalid character: {}", c),
                })
                .collect::<Result<String, _>>()?,
            2,
        )?;

        Ok(Seat { row, column })
    }
}

impl Seat {
    fn id(&self) -> u16 {
        u16::from(self.row) * 8 + u16::from(self.column)
    }
}

fn parse_input(data: &str) -> Result<Vec<Seat>> {
    data.lines()
        .map(|line| line.trim())
        .filter(|line| !line.is_empty())
        .map(Seat::from_str)
        .collect()
}

fn main() -> Result<()> {
    println!("Day 5: Binary Boarding");

    let input = parse_input(include_str!("../input.txt"))?;

    println!("Part 1: {:?}", input.iter().map(|s| s.id()).max());
    println!(
        "Part 2: {:?}",
        input
            .iter()
            .map(|s| s.id())
            .sorted()
            .tuple_windows::<(_, _)>()
            .find_map(|(s1, s2)| if s1 + 2 == s2 { Some(s1 + 1) } else { None })
    );

    Ok(())
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
