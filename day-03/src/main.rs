use anyhow::{bail, Result};
use std::{convert::TryFrom, str::FromStr};

#[derive(Debug, PartialEq)]
enum Tile {
    Empty,
    Tree,
}

impl TryFrom<char> for Tile {
    type Error = anyhow::Error;

    fn try_from(value: char) -> Result<Self, Self::Error> {
        match value {
            '.' => Ok(Tile::Empty),
            '#' => Ok(Tile::Tree),
            c => bail!("invalid character: {}", c),
        }
    }
}

#[derive(Debug, PartialEq)]
struct Grid(Vec<Vec<Tile>>);

impl FromStr for Grid {
    type Err = anyhow::Error;

    fn from_str(s: &str) -> Result<Self> {
        Ok(Grid(
            s.lines()
                .map(|line| line.trim())
                .filter(|line| !line.is_empty())
                .map(|line| line.chars().map(Tile::try_from).collect())
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
            .filter(|(i, row)| row[(i * right) % self.width()] == Tile::Tree)
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
                vec![
                    Tile::Empty,
                    Tile::Empty,
                    Tile::Tree,
                    Tile::Tree,
                    Tile::Empty,
                    Tile::Empty,
                    Tile::Empty,
                    Tile::Empty,
                    Tile::Empty,
                    Tile::Empty,
                    Tile::Empty
                ],
                vec![
                    Tile::Tree,
                    Tile::Empty,
                    Tile::Empty,
                    Tile::Empty,
                    Tile::Tree,
                    Tile::Empty,
                    Tile::Empty,
                    Tile::Empty,
                    Tile::Tree,
                    Tile::Empty,
                    Tile::Empty
                ],
                vec![
                    Tile::Empty,
                    Tile::Tree,
                    Tile::Empty,
                    Tile::Empty,
                    Tile::Empty,
                    Tile::Empty,
                    Tile::Tree,
                    Tile::Empty,
                    Tile::Empty,
                    Tile::Tree,
                    Tile::Empty
                ],
                vec![
                    Tile::Empty,
                    Tile::Empty,
                    Tile::Tree,
                    Tile::Empty,
                    Tile::Tree,
                    Tile::Empty,
                    Tile::Empty,
                    Tile::Empty,
                    Tile::Tree,
                    Tile::Empty,
                    Tile::Tree
                ],
                vec![
                    Tile::Empty,
                    Tile::Tree,
                    Tile::Empty,
                    Tile::Empty,
                    Tile::Empty,
                    Tile::Tree,
                    Tile::Tree,
                    Tile::Empty,
                    Tile::Empty,
                    Tile::Tree,
                    Tile::Empty
                ],
                vec![
                    Tile::Empty,
                    Tile::Empty,
                    Tile::Tree,
                    Tile::Empty,
                    Tile::Tree,
                    Tile::Tree,
                    Tile::Empty,
                    Tile::Empty,
                    Tile::Empty,
                    Tile::Empty,
                    Tile::Empty
                ],
                vec![
                    Tile::Empty,
                    Tile::Tree,
                    Tile::Empty,
                    Tile::Tree,
                    Tile::Empty,
                    Tile::Tree,
                    Tile::Empty,
                    Tile::Empty,
                    Tile::Empty,
                    Tile::Empty,
                    Tile::Tree
                ],
                vec![
                    Tile::Empty,
                    Tile::Tree,
                    Tile::Empty,
                    Tile::Empty,
                    Tile::Empty,
                    Tile::Empty,
                    Tile::Empty,
                    Tile::Empty,
                    Tile::Empty,
                    Tile::Empty,
                    Tile::Tree
                ],
                vec![
                    Tile::Tree,
                    Tile::Empty,
                    Tile::Tree,
                    Tile::Tree,
                    Tile::Empty,
                    Tile::Empty,
                    Tile::Empty,
                    Tile::Tree,
                    Tile::Empty,
                    Tile::Empty,
                    Tile::Empty
                ],
                vec![
                    Tile::Tree,
                    Tile::Empty,
                    Tile::Empty,
                    Tile::Empty,
                    Tile::Tree,
                    Tile::Tree,
                    Tile::Empty,
                    Tile::Empty,
                    Tile::Empty,
                    Tile::Empty,
                    Tile::Tree
                ],
                vec![
                    Tile::Empty,
                    Tile::Tree,
                    Tile::Empty,
                    Tile::Empty,
                    Tile::Tree,
                    Tile::Empty,
                    Tile::Empty,
                    Tile::Empty,
                    Tile::Tree,
                    Tile::Empty,
                    Tile::Tree
                ]
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
