use anyhow::Result;
use itertools::Itertools;

fn is_valid(value: u64, preamble: &[u64]) -> bool {
    preamble
        .iter()
        .tuple_combinations()
        .map(|(x, y)| x + y)
        .any(|n| n == value)
}

fn find_invalid(data: &[u64], len: usize) -> Option<u64> {
    data.windows(len + 1).find_map(|window| {
        let value = window[len];
        let is_valid = is_valid(value, &window[..len]);

        if is_valid {
            None
        } else {
            Some(value)
        }
    })
}

fn find_contiguous_range(data: &[u64], value: u64) -> Option<(usize, usize)> {
    (2..data.len())
        .into_iter()
        .flat_map(|n| {
            data.windows(n)
                .enumerate()
                .map(move |(i, s)| (i, n, s.iter().sum::<u64>()))
        })
        .find_map(|(i, n, sum)| (sum == value).then_some((i, i + n)))
}

fn find_contiguous_min_max(data: &[u64], value: u64) -> Option<(u64, u64)> {
    let (start, end) = find_contiguous_range(data, value)?;
    let min = data[start..end].iter().min()?;
    let max = data[start..end].iter().max()?;

    Some((*min, *max))
}

#[aoc_generator(day9)]
fn generator(input: &str) -> Result<Vec<u64>> {
    input
        .lines()
        .map(str::parse)
        .collect::<Result<Vec<_>, _>>()
        .map_err(Into::into)
}

#[aoc(day9, part1)]
fn part1(data: &[u64]) -> Option<u64> {
    find_invalid(data, 25)
}

#[aoc(day9, part2)]
fn part2(data: &[u64]) -> Option<u64> {
    find_contiguous_min_max(&data, 556543474).map(|(min, max)| min + max)
}

#[cfg(test)]
mod tests {
    use super::*;
    use indoc::indoc;

    #[test]
    fn it_parses_input() {
        let data = indoc! {"
            35
            20
            15
            25
            47
            40
            62
            55
            65
            95
            102
            117
            150
            182
            127
            219
            299
            277
            309
            576
        "};

        let input = generator(data).expect("input to be parsed");
        assert_eq!(
            [
                35, 20, 15, 25, 47, 40, 62, 55, 65, 95, 102, 117, 150, 182, 127, 219, 299, 277,
                309, 576
            ]
            .to_vec(),
            input
        );
    }

    #[test]
    fn it_solves_part1() {
        let data = indoc! {"
            35
            20
            15
            25
            47
            40
            62
            55
            65
            95
            102
            117
            150
            182
            127
            219
            299
            277
            309
            576
        "};

        let input = generator(data).expect("input to be parsed");
        assert_eq!(Some(127), find_invalid(&input, 5));
    }

    #[test]
    fn it_solves_part2() {
        let data = indoc! {"
            35
            20
            15
            25
            47
            40
            62
            55
            65
            95
            102
            117
            150
            182
            127
            219
            299
            277
            309
            576
        "};

        let input = generator(data).expect("input to be parsed");
        assert_eq!(Some((15, 47)), find_contiguous_min_max(&input, 127));
    }
}
