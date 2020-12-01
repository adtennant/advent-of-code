use anyhow::{Context, Result};
use std::{collections::HashSet, num::ParseIntError};

fn parse_input(data: &str) -> Result<HashSet<i32>, ParseIntError> {
    data.lines()
        .map(|line| line.trim())
        .filter(|line| !line.is_empty())
        .map(|line| line.parse::<i32>())
        .collect()
}

fn find_pair(parsed_input: &HashSet<i32>, total: i32) -> Option<i32> {
    parsed_input.iter().find_map(|x| {
        let y = total - x;
        parsed_input.get(&y).map(|_| x * y)
    })
}

fn find_triple(parsed_input: &HashSet<i32>, total: i32) -> Option<i32> {
    parsed_input
        .iter()
        .find_map(|x| find_pair(parsed_input, total - x).map(|result| x * result))
}

fn main() -> Result<()> {
    let parsed_input = parse_input(include_str!("../input"))?;

    println!(
        "Part 1: {:?}",
        find_pair(&parsed_input, 2020).with_context(|| "no matching pair found")?
    );
    println!(
        "Part 2: {:?}",
        find_triple(&parsed_input, 2020).with_context(|| "no matching triple found")?
    );

    Ok(())
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_parses_input() {
        let data = "
            1721
            979
            366
            299
            675
            1456
        ";

        let parsed_input = parse_input(data).unwrap();

        assert_eq!(
            [1721, 979, 366, 299, 675, 1456]
                .iter()
                .cloned()
                .collect::<HashSet<i32>>(),
            parsed_input
        )
    }

    #[test]
    fn it_finds_a_pair() {
        let data = "
            1721
            979
            366
            299
            675
            1456
        ";

        let parsed_input = parse_input(data).unwrap();
        let result = find_pair(&parsed_input, 2020).unwrap();

        assert_eq!(result, 514579)
    }

    #[test]
    fn it_finds_a_triple() {
        let data = "
            1721
            979
            366
            299
            675
            1456
        ";

        let parsed_input = parse_input(data).unwrap();
        let result = find_triple(&parsed_input, 2020).unwrap();

        assert_eq!(result, 241861950)
    }
}
