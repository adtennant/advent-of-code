use anyhow::{bail, Result};
use std::str::FromStr;

#[derive(Debug, PartialEq)]
struct Grid(Vec<Vec<bool>>);

impl FromStr for Grid {
    type Err = anyhow::Error;

    fn from_str(s: &str) -> Result<Self> {
        Ok(Grid(
            s.lines()
                .map(|line| line.trim())
                .filter(|line| !line.is_empty())
                .map(|line| {
                    line.chars()
                        .map(|c| match c {
                            '#' => Ok(true),
                            '.' => Ok(false),
                            c => bail!("invalid character: {}", c),
                        })
                        .collect()
                })
                .collect::<Result<_, _>>()?,
        ))
    }
}

impl Grid {
    fn width(&self) -> usize {
        if !self.0.is_empty() {
            self.0[0].len()
        } else {
            0
        }
    }

    fn count_trees_for_trajectory(&self, right: usize, down: usize) -> usize {
        self.0
            .iter()
            .step_by(down)
            .enumerate()
            .filter(|(i, row)| row[(i * right) % self.width()])
            .count()
    }
}

fn main() -> Result<()> {
    println!("Day 3: Toboggan Trajectory");

    let grid = Grid::from_str(include_str!("../input.txt"))?;
    println!("Part 1: {:?}", grid.count_trees_for_trajectory(3, 1));
    println!(
        "Part 2: {:?}",
        [(1, 1), (3, 1), (5, 1), (7, 1), (1, 2)]
            .iter()
            .map(|&(right, down)| grid.count_trees_for_trajectory(right, down))
            .product::<usize>()
    );

    Ok(())
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_parses_a_grid() {
        let data = "
            ..##.......
            #...#...#..
            .#....#..#.
            ..#.#...#.#
            .#...##..#.
            ..#.##.....
            .#.#.#....#
            .#........#
            #.##...#...
            #...##....#
            .#..#...#.#
        ";
        let grid = Grid::from_str(data).unwrap();

        assert_eq!(
            Grid(vec![
                vec![false, false, true, true, false, false, false, false, false, false, false],
                vec![true, false, false, false, true, false, false, false, true, false, false],
                vec![false, true, false, false, false, false, true, false, false, true, false],
                vec![false, false, true, false, true, false, false, false, true, false, true],
                vec![false, true, false, false, false, true, true, false, false, true, false],
                vec![false, false, true, false, true, true, false, false, false, false, false],
                vec![false, true, false, true, false, true, false, false, false, false, true],
                vec![false, true, false, false, false, false, false, false, false, false, true],
                vec![true, false, true, true, false, false, false, true, false, false, false],
                vec![true, false, false, false, true, true, false, false, false, false, true],
                vec![false, true, false, false, true, false, false, false, true, false, true]
            ]),
            grid
        )
    }

    #[test]
    fn it_counts_trees_for_a_trajectory() {
        let data = "
            ..##.......
            #...#...#..
            .#....#..#.
            ..#.#...#.#
            .#...##..#.
            ..#.##.....
            .#.#.#....#
            .#........#
            #.##...#...
            #...##....#
            .#..#...#.#
        ";
        let grid = Grid::from_str(data).unwrap();

        assert_eq!(2, grid.count_trees_for_trajectory(1, 1));
        assert_eq!(7, grid.count_trees_for_trajectory(3, 1));
        assert_eq!(3, grid.count_trees_for_trajectory(5, 1));
        assert_eq!(4, grid.count_trees_for_trajectory(7, 1));
        assert_eq!(2, grid.count_trees_for_trajectory(1, 2))
    }
}
