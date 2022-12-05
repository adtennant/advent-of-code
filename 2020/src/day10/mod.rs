use anyhow::Result;
use itertools::Itertools;
use std::collections::HashMap;

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
    use indoc::indoc;

    #[test]
    fn it_parses_input() {
        let data = indoc! {"
            16
            10
            15
            5
            1
            11
            7
            19
            6
            12
            4
        "};

        let input = generator(data).expect("input to be parsed");
        assert_eq!([16, 10, 15, 5, 1, 11, 7, 19, 6, 12, 4].to_vec(), input);
    }

    #[test]
    fn it_solves_part1() {
        let data = indoc! {"
            16
            10
            15
            5
            1
            11
            7
            19
            6
            12
            4
        "};

        let input = generator(data).expect("input to be parsed");

        let mut result = HashMap::new();
        result.insert(1, 7);
        result.insert(3, 5);

        assert_eq!(result, find_differences(&input));

        let data = indoc! {"
            28
            33
            18
            42
            31
            14
            46
            20
            48
            47
            24
            23
            49
            45
            19
            38
            39
            11
            1
            32
            25
            35
            8
            17
            7
            9
            4
            2
            34
            10
            3
        "};

        let input = generator(data).expect("input to be parsed");

        let mut result = HashMap::new();
        result.insert(1, 22);
        result.insert(3, 10);

        assert_eq!(result, find_differences(&input));
    }

    #[test]
    fn it_solves_part2() {
        let data = indoc! {"
            16
            10
            15
            5
            1
            11
            7
            19
            6
            12
            4
        "};

        let input = generator(data).expect("input to be parsed");
        assert_eq!(8, find_combinations(&input));

        let data = indoc! {"
            28
            33
            18
            42
            31
            14
            46
            20
            48
            47
            24
            23
            49
            45
            19
            38
            39
            11
            1
            32
            25
            35
            8
            17
            7
            9
            4
            2
            34
            10
            3
        "};

        let input = generator(data).expect("input to be parsed");
        assert_eq!(19208, find_combinations(&input));
    }
}
