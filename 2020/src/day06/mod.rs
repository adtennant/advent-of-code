use anyhow::Result;
use itertools::Itertools;
use std::collections::HashSet;

#[aoc_generator(day6)]
fn generator(input: &str) -> Vec<String> {
    input.split("\n\n").map(String::from).collect()
}

#[aoc(day6, part1)]
fn part1(input: &[String]) -> usize {
    input
        .iter()
        .map(|s| s.chars().filter(|c| !c.is_whitespace()).unique().count())
        .sum::<usize>()
}

#[aoc(day6, part2)]
fn part2(input: &[String]) -> usize {
    input
        .iter()
        .map(|s| {
            s.lines()
                .map(|l| l.chars().collect::<HashSet<_>>())
                .reduce(|result, chars| result.intersection(&chars).cloned().collect())
        })
        .map(|s| s.unwrap_or_default().len())
        .sum::<usize>()
}

#[cfg(test)]
mod tests {
    use super::*;
    use indoc::indoc;

    #[test]
    fn it_parses_input() {
        let data = indoc! {"
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
        "};
        let input = generator(data);
        assert_eq!(5, input.len())
    }

    #[test]
    fn it_counts_where_anyone_answered_yes() {
        let data = indoc! {"
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
        "};
        let input = generator(data);
        assert_eq!(11, part1(&input))
    }

    #[test]
    fn it_counts_where_everyone_answered_yes() {
        let data = indoc! {"
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
        "};
        let input = generator(data);
        assert_eq!(6, part2(&input))
    }
}
