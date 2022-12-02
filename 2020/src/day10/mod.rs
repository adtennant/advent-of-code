use std::collections::HashMap;

use anyhow::Result;
use itertools::Itertools;

fn find_differences(data: &[u64]) -> HashMap<u64, usize> {
    let max = data.iter().max().unwrap();

    [0, max + 3]
        .iter()
        .chain(data.into_iter())
        .sorted()
        .tuple_windows()
        .counts_by(|(x, y)| y - x)
}

fn find_combinations(data: &[u64]) -> usize {
    let max = data.iter().max().unwrap();

    let mut paths = HashMap::new();
    paths.insert(max + 3, 1);

    let paths = [0]
        .iter()
        .chain(data.into_iter())
        .sorted()
        .rev()
        .fold(paths, |mut result, &n| {
            let count: usize = (1..=3)
                .into_iter()
                .filter_map(|i| result.get(&(n + i)))
                .sum();

            result.insert(n, count);
            result
        });

    paths[&0]
}

#[aoc_generator(day10)]
fn generator(input: &str) -> Result<Vec<u64>> {
    input
        .lines()
        .map(str::parse)
        .collect::<Result<Vec<_>, _>>()
        .map_err(Into::into)
}

#[aoc(day10, part1)]
fn part1(data: &[u64]) -> usize {
    let differences = find_differences(data);
    differences[&1] * differences[&3]
}

#[aoc(day10, part2)]
fn part2(data: &[u64]) -> usize {
    find_combinations(data)
}

#[cfg(test)]
mod tests {
    use super::*;

    const DATA: &'static str = "16\n10\n15\n5\n1\n11\n7\n19\n6\n12\n4";
    const DATA_2: &'static str = "28\n33\n18\n42\n31\n14\n46\n20\n48\n47\n24\n23\n49\n45\n19\n38\n39\n11\n1\n32\n25\n35\n8\n17\n7\n9\n4\n2\n34\n10\n3";

    #[test]
    fn it_parses_input() {
        let data = generator(DATA).expect("input to be parsed");
        assert_eq!([16, 10, 15, 5, 1, 11, 7, 19, 6, 12, 4].to_vec(), data);
    }

    #[test]
    fn it_solves_part1() {
        let data = generator(DATA).expect("input to be parsed");

        let mut result = HashMap::new();
        result.insert(1, 7);
        result.insert(3, 5);

        assert_eq!(result, find_differences(&data));

        let data = generator(DATA_2).expect("input to be parsed");

        let mut result = HashMap::new();
        result.insert(1, 22);
        result.insert(3, 10);

        assert_eq!(result, find_differences(&data));
    }

    #[test]
    fn it_solves_part2() {
        let data = generator(DATA).expect("input to be parsed");
        assert_eq!(8, find_combinations(&data));

        let data = generator(DATA_2).expect("input to be parsed");
        assert_eq!(19208, find_combinations(&data));
    }
}
