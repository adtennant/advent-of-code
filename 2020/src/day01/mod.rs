use anyhow::Result;
use std::{collections::HashSet, num::ParseIntError};

fn find_pair(input: &HashSet<i32>, total: i32) -> Option<i32> {
    input.iter().find_map(|x| {
        let y = total - x;
        input.get(&y).map(|_| x * y)
    })
}

fn find_triple(input: &HashSet<i32>, total: i32) -> Option<i32> {
    input
        .iter()
        .find_map(|x| find_pair(input, total - x).map(|result| x * result))
}

#[aoc_generator(day1)]
pub fn generator(data: &str) -> Result<HashSet<i32>, ParseIntError> {
    data.lines().map(|line| line.parse::<i32>()).collect()
}

#[aoc(day1, part1)]
pub fn part1(input: &HashSet<i32>) -> Option<i32> {
    find_pair(input, 2020)
}

#[aoc(day1, part2)]
pub fn part2(input: &HashSet<i32>) -> Option<i32> {
    find_triple(&input, 2020)
}

#[cfg(test)]
mod tests {
    use super::*;
    use indoc::indoc;

    #[test]
    fn it_parses_input() {
        let data = indoc! {"
            1721
            979
            366
            299
            675
            1456
        "};

        let input = generator(data).unwrap();

        assert_eq!(
            [1721, 979, 366, 299, 675, 1456]
                .iter()
                .cloned()
                .collect::<HashSet<i32>>(),
            input
        )
    }

    #[test]
    fn it_finds_a_pair() {
        let data = indoc! {"
            1721
            979
            366
            299
            675
            1456
        "};

        let input = generator(data).unwrap();
        let result = find_pair(&input, 2020).unwrap();

        assert_eq!(result, 514579)
    }

    #[test]
    fn it_finds_a_triple() {
        let data = indoc! {"
            1721
            979
            366
            299
            675
            1456
        "};

        let input = generator(data).unwrap();
        let result = find_triple(&input, 2020).unwrap();

        assert_eq!(result, 241861950)
    }
}
