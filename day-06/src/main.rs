use anyhow::Result;
use itertools::Itertools;
use regex::Regex;
use std::collections::HashSet;

fn parse_input(input: &str) -> Result<Vec<&str>> {
    Ok(Regex::new("\n *\n")?.split(input).collect())
}

fn count_anyone(input: &[&str]) -> usize {
    input
        .iter()
        .map(|s| s.chars().filter(|c| !c.is_whitespace()).unique().count())
        .sum::<usize>()
}

fn count_everyone(input: &[&str]) -> usize {
    input
        .iter()
        .map(|s| {
            s.lines()
                .map(|l| l.chars().collect::<HashSet<_>>())
                .fold1(|result, chars| result.intersection(&chars).cloned().collect())
        })
        .map(|s| s.unwrap_or_default().len())
        .sum::<usize>()
}

fn main() -> Result<()> {
    println!("Day 6: Custom Customs");

    let input = parse_input(include_str!("../input.txt"))?;

    println!("Part 1: {:?}", count_anyone(&input));

    println!("Part 2: {:?}", count_everyone(&input));

    Ok(())
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_parses_input() {
        let data = "
            abc

            a
            b
            c

            ab
            ac

            a
            a
            a
            a

            b
        ";
        let input = parse_input(data).unwrap();
        assert_eq!(5, input.len())
    }

    #[test]
    fn it_counts_where_anyone_answered_yes() {
        let data = "
            abc

            a
            b
            c

            ab
            ac

            a
            a
            a
            a

            b
        ";
        let input = parse_input(data).unwrap();
        assert_eq!(11, count_anyone(&input))
    }

    #[test]
    fn it_counts_where_everyone_answered_yes() {
        let data = "
            abc

            a
            b
            c

            ab
            ac

            a
            a
            a
            a

            b
        ";
        let input = parse_input(data).unwrap();
        assert_eq!(6, count_everyone(&input))
    }
}
