use std::{convert::TryFrom, str::FromStr};

use anyhow::{bail, Result};

#[derive(Clone, Copy, Debug, Eq, PartialEq)]
enum Tile {
    Floor,
    Seat(bool),
}

impl TryFrom<char> for Tile {
    type Error = anyhow::Error;

    fn try_from(value: char) -> Result<Self, Self::Error> {
        Ok(match value {
            '.' => Tile::Floor,
            'L' => Tile::Seat(false),
            '#' => Tile::Seat(true),
            c => bail!("invalid character: {}", c),
        })
    }
}

impl Tile {
    fn next(self, neighbours: &[Tile]) -> Tile {
        let occupied_neighbours = neighbours
            .iter()
            .filter(|n| matches!(n, Tile::Seat(true)))
            .count();

        match self {
            Tile::Seat(false) if occupied_neighbours == 0 => Tile::Seat(true),
            Tile::Seat(true) if occupied_neighbours >= 4 => Tile::Seat(false),
            _ => self,
        }
    }
}

#[derive(Clone, Debug, Eq, PartialEq)]
struct Grid(Vec<Vec<Tile>>);

impl FromStr for Grid {
    type Err = anyhow::Error;

    fn from_str(s: &str) -> Result<Self> {
        s.lines()
            .map(|line| line.chars().map(Tile::try_from).collect())
            .collect::<Result<_, _>>()
            .map(Grid)
    }
}

impl Grid {
    fn get(&self, x: i32, y: i32) -> Option<Tile> {
        if x >= 0 && x < self.width() && y >= 0 && y < self.height() {
            Some(self.0[y as usize][x as usize])
        } else {
            None
        }
    }

    fn height(&self) -> i32 {
        self.0.len() as i32
    }

    fn neighbours(&self, x: usize, y: usize) -> Vec<Tile> {
        [
            (-1, -1),
            (-1, 0),
            (-1, 1),
            (0, -1),
            (0, 1),
            (1, -1),
            (1, 0),
            (1, 1),
        ]
        .iter()
        .map(|(dx, dy)| (x as i32 + dx, y as i32 + dy))
        .filter_map(|(x, y)| self.get(x, y))
        .collect()
    }

    fn occupied(&self) -> usize {
        self.0
            .iter()
            .flatten()
            .filter(|&&tile| tile == Tile::Seat(true))
            .count()
    }

    fn width(&self) -> i32 {
        if !self.0.is_empty() {
            self.0[0].len() as i32
        } else {
            0
        }
    }
}

fn process_round(grid: &Grid) -> Grid {
    Grid(
        grid.0
            .iter()
            .enumerate()
            .map(|(y, row)| {
                row.iter()
                    .enumerate()
                    .map(|(x, &tile)| {
                        let neighbours = grid.neighbours(x, y);
                        tile.next(&neighbours)
                    })
                    .collect()
            })
            .collect(),
    )
}

#[aoc_generator(day11)]
fn generator(input: &str) -> Result<Grid> {
    Grid::from_str(&input)
}

#[aoc(day11, part1)]
fn part1(data: &Grid) -> usize {
    let mut old = data.clone();

    loop {
        let new = process_round(&old);

        if new == old {
            return new.occupied();
        }

        old = new;
    }
}

#[aoc(day11, part2)]
fn part2(data: &Grid) -> usize {
    unimplemented!()
}

#[cfg(test)]
mod tests {
    use super::*;
    use Tile::*;

    const DATA: &'static str = "L.LL.LL.LL\nLLLLLLL.LL\nL.L.L..L..\nLLLL.LL.LL\nL.LL.LL.LL\nL.LLLLL.LL\n..L.L.....\nLLLLLLLLLL\nL.LLLLLL.L\nL.LLLLL.LL";

    #[test]
    fn it_parses_input() {
        let data = generator(DATA).expect("input to be parsed");
        assert_eq!(
            Grid(vec![
                vec![
                    Seat(false),
                    Floor,
                    Seat(false),
                    Seat(false),
                    Floor,
                    Seat(false),
                    Seat(false),
                    Floor,
                    Seat(false),
                    Seat(false)
                ],
                vec![
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Floor,
                    Seat(false),
                    Seat(false)
                ],
                vec![
                    Seat(false),
                    Floor,
                    Seat(false),
                    Floor,
                    Seat(false),
                    Floor,
                    Floor,
                    Seat(false),
                    Floor,
                    Floor
                ],
                vec![
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Floor,
                    Seat(false),
                    Seat(false),
                    Floor,
                    Seat(false),
                    Seat(false)
                ],
                vec![
                    Seat(false),
                    Floor,
                    Seat(false),
                    Seat(false),
                    Floor,
                    Seat(false),
                    Seat(false),
                    Floor,
                    Seat(false),
                    Seat(false)
                ],
                vec![
                    Seat(false),
                    Floor,
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Floor,
                    Seat(false),
                    Seat(false)
                ],
                vec![
                    Floor,
                    Floor,
                    Seat(false),
                    Floor,
                    Seat(false),
                    Floor,
                    Floor,
                    Floor,
                    Floor,
                    Floor
                ],
                vec![
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Seat(false)
                ],
                vec![
                    Seat(false),
                    Floor,
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Floor,
                    Seat(false)
                ],
                vec![
                    Seat(false),
                    Floor,
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Floor,
                    Seat(false),
                    Seat(false)
                ]
            ]),
            data
        );
    }

    #[test]
    fn it_solves_part1() {
        let data = generator(DATA).expect("input to be parsed");
        let result = process_round(&data);

        assert_eq!(
            Grid(vec![
                vec![
                    Seat(true),
                    Floor,
                    Seat(true),
                    Seat(true),
                    Floor,
                    Seat(true),
                    Seat(true),
                    Floor,
                    Seat(true),
                    Seat(true)
                ],
                vec![
                    Seat(true),
                    Seat(true),
                    Seat(true),
                    Seat(true),
                    Seat(true),
                    Seat(true),
                    Seat(true),
                    Floor,
                    Seat(true),
                    Seat(true)
                ],
                vec![
                    Seat(true),
                    Floor,
                    Seat(true),
                    Floor,
                    Seat(true),
                    Floor,
                    Floor,
                    Seat(true),
                    Floor,
                    Floor
                ],
                vec![
                    Seat(true),
                    Seat(true),
                    Seat(true),
                    Seat(true),
                    Floor,
                    Seat(true),
                    Seat(true),
                    Floor,
                    Seat(true),
                    Seat(true)
                ],
                vec![
                    Seat(true),
                    Floor,
                    Seat(true),
                    Seat(true),
                    Floor,
                    Seat(true),
                    Seat(true),
                    Floor,
                    Seat(true),
                    Seat(true)
                ],
                vec![
                    Seat(true),
                    Floor,
                    Seat(true),
                    Seat(true),
                    Seat(true),
                    Seat(true),
                    Seat(true),
                    Floor,
                    Seat(true),
                    Seat(true)
                ],
                vec![
                    Floor,
                    Floor,
                    Seat(true),
                    Floor,
                    Seat(true),
                    Floor,
                    Floor,
                    Floor,
                    Floor,
                    Floor
                ],
                vec![
                    Seat(true),
                    Seat(true),
                    Seat(true),
                    Seat(true),
                    Seat(true),
                    Seat(true),
                    Seat(true),
                    Seat(true),
                    Seat(true),
                    Seat(true)
                ],
                vec![
                    Seat(true),
                    Floor,
                    Seat(true),
                    Seat(true),
                    Seat(true),
                    Seat(true),
                    Seat(true),
                    Seat(true),
                    Floor,
                    Seat(true)
                ],
                vec![
                    Seat(true),
                    Floor,
                    Seat(true),
                    Seat(true),
                    Seat(true),
                    Seat(true),
                    Seat(true),
                    Floor,
                    Seat(true),
                    Seat(true)
                ]
            ]),
            result
        );

        let result = process_round(&result);

        assert_eq!(
            Grid(vec![
                vec![
                    Seat(true),
                    Floor,
                    Seat(false),
                    Seat(false),
                    Floor,
                    Seat(false),
                    Seat(true),
                    Floor,
                    Seat(true),
                    Seat(true)
                ],
                vec![
                    Seat(true),
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Floor,
                    Seat(false),
                    Seat(true)
                ],
                vec![
                    Seat(false),
                    Floor,
                    Seat(false),
                    Floor,
                    Seat(false),
                    Floor,
                    Floor,
                    Seat(false),
                    Floor,
                    Floor
                ],
                vec![
                    Seat(true),
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Floor,
                    Seat(false),
                    Seat(false),
                    Floor,
                    Seat(false),
                    Seat(true)
                ],
                vec![
                    Seat(true),
                    Floor,
                    Seat(false),
                    Seat(false),
                    Floor,
                    Seat(false),
                    Seat(false),
                    Floor,
                    Seat(false),
                    Seat(false)
                ],
                vec![
                    Seat(true),
                    Floor,
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Seat(true),
                    Floor,
                    Seat(true),
                    Seat(true)
                ],
                vec![
                    Floor,
                    Floor,
                    Seat(false),
                    Floor,
                    Seat(false),
                    Floor,
                    Floor,
                    Floor,
                    Floor,
                    Floor
                ],
                vec![
                    Seat(true),
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Seat(true)
                ],
                vec![
                    Seat(true),
                    Floor,
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Floor,
                    Seat(false)
                ],
                vec![
                    Seat(true),
                    Floor,
                    Seat(true),
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Floor,
                    Seat(true),
                    Seat(true)
                ]
            ]),
            result
        );

        let result = process_round(&result);

        assert_eq!(
            Grid(vec![
                vec![
                    Seat(true),
                    Floor,
                    Seat(true),
                    Seat(true),
                    Floor,
                    Seat(false),
                    Seat(true),
                    Floor,
                    Seat(true),
                    Seat(true)
                ],
                vec![
                    Seat(true),
                    Seat(false),
                    Seat(true),
                    Seat(true),
                    Seat(true),
                    Seat(false),
                    Seat(false),
                    Floor,
                    Seat(false),
                    Seat(true)
                ],
                vec![
                    Seat(false),
                    Floor,
                    Seat(true),
                    Floor,
                    Seat(true),
                    Floor,
                    Floor,
                    Seat(true),
                    Floor,
                    Floor
                ],
                vec![
                    Seat(true),
                    Seat(false),
                    Seat(true),
                    Seat(true),
                    Floor,
                    Seat(true),
                    Seat(true),
                    Floor,
                    Seat(false),
                    Seat(true)
                ],
                vec![
                    Seat(true),
                    Floor,
                    Seat(true),
                    Seat(true),
                    Floor,
                    Seat(false),
                    Seat(false),
                    Floor,
                    Seat(false),
                    Seat(false)
                ],
                vec![
                    Seat(true),
                    Floor,
                    Seat(true),
                    Seat(true),
                    Seat(true),
                    Seat(false),
                    Seat(true),
                    Floor,
                    Seat(true),
                    Seat(true)
                ],
                vec![
                    Floor,
                    Floor,
                    Seat(true),
                    Floor,
                    Seat(true),
                    Floor,
                    Floor,
                    Floor,
                    Floor,
                    Floor
                ],
                vec![
                    Seat(true),
                    Seat(false),
                    Seat(true),
                    Seat(true),
                    Seat(true),
                    Seat(true),
                    Seat(true),
                    Seat(true),
                    Seat(false),
                    Seat(true)
                ],
                vec![
                    Seat(true),
                    Floor,
                    Seat(false),
                    Seat(false),
                    Seat(true),
                    Seat(true),
                    Seat(true),
                    Seat(false),
                    Floor,
                    Seat(false)
                ],
                vec![
                    Seat(true),
                    Floor,
                    Seat(true),
                    Seat(false),
                    Seat(true),
                    Seat(true),
                    Seat(true),
                    Floor,
                    Seat(true),
                    Seat(true)
                ],
            ]),
            result
        );

        let result = process_round(&result);

        assert_eq!(
            Grid(vec![
                vec![
                    Seat(true),
                    Floor,
                    Seat(true),
                    Seat(false),
                    Floor,
                    Seat(false),
                    Seat(true),
                    Floor,
                    Seat(true),
                    Seat(true)
                ],
                vec![
                    Seat(true),
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Seat(true),
                    Seat(false),
                    Seat(false),
                    Floor,
                    Seat(false),
                    Seat(true)
                ],
                vec![
                    Seat(false),
                    Floor,
                    Seat(false),
                    Floor,
                    Seat(false),
                    Floor,
                    Floor,
                    Seat(true),
                    Floor,
                    Floor
                ],
                vec![
                    Seat(true),
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Floor,
                    Seat(true),
                    Seat(true),
                    Floor,
                    Seat(false),
                    Seat(true)
                ],
                vec![
                    Seat(true),
                    Floor,
                    Seat(false),
                    Seat(false),
                    Floor,
                    Seat(false),
                    Seat(false),
                    Floor,
                    Seat(false),
                    Seat(false)
                ],
                vec![
                    Seat(true),
                    Floor,
                    Seat(false),
                    Seat(false),
                    Seat(true),
                    Seat(false),
                    Seat(true),
                    Floor,
                    Seat(true),
                    Seat(true)
                ],
                vec![
                    Floor,
                    Floor,
                    Seat(false),
                    Floor,
                    Seat(false),
                    Floor,
                    Floor,
                    Floor,
                    Floor,
                    Floor
                ],
                vec![
                    Seat(true),
                    Seat(false),
                    Seat(true),
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Seat(true),
                    Seat(false),
                    Seat(true)
                ],
                vec![
                    Seat(true),
                    Floor,
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Floor,
                    Seat(false)
                ],
                vec![
                    Seat(true),
                    Floor,
                    Seat(true),
                    Seat(false),
                    Seat(true),
                    Seat(false),
                    Seat(true),
                    Floor,
                    Seat(true),
                    Seat(true)
                ],
            ]),
            result
        );

        let result = process_round(&result);

        assert_eq!(
            Grid(vec![
                vec![
                    Seat(true),
                    Floor,
                    Seat(true),
                    Seat(false),
                    Floor,
                    Seat(false),
                    Seat(true),
                    Floor,
                    Seat(true),
                    Seat(true)
                ],
                vec![
                    Seat(true),
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Seat(true),
                    Seat(false),
                    Seat(false),
                    Floor,
                    Seat(false),
                    Seat(true)
                ],
                vec![
                    Seat(false),
                    Floor,
                    Seat(true),
                    Floor,
                    Seat(false),
                    Floor,
                    Floor,
                    Seat(true),
                    Floor,
                    Floor
                ],
                vec![
                    Seat(true),
                    Seat(false),
                    Seat(true),
                    Seat(true),
                    Floor,
                    Seat(true),
                    Seat(true),
                    Floor,
                    Seat(false),
                    Seat(true)
                ],
                vec![
                    Seat(true),
                    Floor,
                    Seat(true),
                    Seat(false),
                    Floor,
                    Seat(false),
                    Seat(false),
                    Floor,
                    Seat(false),
                    Seat(false)
                ],
                vec![
                    Seat(true),
                    Floor,
                    Seat(true),
                    Seat(false),
                    Seat(true),
                    Seat(false),
                    Seat(true),
                    Floor,
                    Seat(true),
                    Seat(true)
                ],
                vec![
                    Floor,
                    Floor,
                    Seat(false),
                    Floor,
                    Seat(false),
                    Floor,
                    Floor,
                    Floor,
                    Floor,
                    Floor
                ],
                vec![
                    Seat(true),
                    Seat(false),
                    Seat(true),
                    Seat(false),
                    Seat(true),
                    Seat(true),
                    Seat(false),
                    Seat(true),
                    Seat(false),
                    Seat(true)
                ],
                vec![
                    Seat(true),
                    Floor,
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Floor,
                    Seat(false)
                ],
                vec![
                    Seat(true),
                    Floor,
                    Seat(true),
                    Seat(false),
                    Seat(true),
                    Seat(false),
                    Seat(true),
                    Floor,
                    Seat(true),
                    Seat(true)
                ],
            ]),
            result
        );

        assert_eq!(result.occupied(), 37);
    }

    #[test]
    fn it_solves_part2() {
        let data = generator(DATA).expect("input to be parsed");
    }
}
