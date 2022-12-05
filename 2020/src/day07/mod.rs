use anyhow::Result;
use std::collections::HashMap;

fn get_contents(
    input: &HashMap<String, Vec<(usize, String)>>,
    color: &str,
) -> Vec<(usize, String)> {
    input
        .get(color)
        .unwrap()
        .iter()
        .flat_map(|content| std::iter::once(content.clone()).chain(get_contents(input, &content.1)))
        .collect()
}

fn count_contents(input: &HashMap<String, Vec<(usize, String)>>, color: &str) -> usize {
    input
        .get(color)
        .unwrap()
        .iter()
        .map(|(number, color)| number * count_contents(input, &color))
        .sum::<usize>()
        + 1
}

fn parse_bag(input: &str) -> (String, Vec<(usize, String)>) {
    let parts: Vec<_> = input.split(" bags contain ").collect();

    let color = parts[0].trim().to_owned();
    let contents = if parts[1] != "no other bags." {
        parts[1]
            .split(", ")
            .map(|s| s.splitn(2, " "))
            .map(|mut parts| {
                (
                    parts.next().unwrap().parse::<usize>().unwrap(),
                    parts
                        .next()
                        .unwrap()
                        .trim_end_matches(" bag")
                        .trim_end_matches(" bags")
                        .trim_end_matches(" bag.")
                        .trim_end_matches(" bags.")
                        .to_owned(),
                )
            })
            .collect()
    } else {
        Vec::new()
    };

    (color, contents)
}

#[aoc_generator(day7)]
fn generator(input: &str) -> HashMap<String, Vec<(usize, String)>> {
    input.lines().map(parse_bag).collect()
}

#[aoc(day7, part1)]
fn part1(input: &HashMap<String, Vec<(usize, String)>>) -> usize {
    input
        .keys()
        .filter(|&k| {
            get_contents(input, k)
                .iter()
                .any(|(_, color)| color == "shiny gold")
        })
        .count()
}

#[aoc(day7, part2)]
fn part2(input: &HashMap<String, Vec<(usize, String)>>) -> usize {
    count_contents(input, "shiny gold") - 1
}

#[cfg(test)]
mod tests {
    use super::*;
    use indoc::indoc;

    #[test]
    fn it_solves_part1() {
        let data = indoc! {"
            light red bags contain 1 bright white bag, 2 muted yellow bags.
            dark orange bags contain 3 bright white bags, 4 muted yellow bags.
            bright white bags contain 1 shiny gold bag.
            muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
            shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
            dark olive bags contain 3 faded blue bags, 4 dotted black bags.
            vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
            faded blue bags contain no other bags.
            dotted black bags contain no other bags.
        "};

        let input = generator(data);
        assert_eq!(4, part1(&input));
    }

    #[test]
    fn it_solves_part2() {
        let data = indoc! {"
            shiny gold bags contain 2 dark red bags.
            dark red bags contain 2 dark orange bags.
            dark orange bags contain 2 dark yellow bags.
            dark yellow bags contain 2 dark green bags.
            dark green bags contain 2 dark blue bags.
            dark blue bags contain 2 dark violet bags.
            dark violet bags contain no other bags.
        "};

        let input = generator(data);
        assert_eq!(126, part2(&input))
    }
}
