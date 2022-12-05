use anyhow::{bail, Context, Result};
use std::{convert::TryFrom, str::FromStr};

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

#[derive(Clone, Debug, Eq, PartialEq)]
struct Grid {
    width: usize,
    height: usize,
    tiles: Vec<Tile>,
}

impl FromStr for Grid {
    type Err = anyhow::Error;

    fn from_str(s: &str) -> Result<Self> {
        let width = s
            .lines()
            .next()
            .context("couldn't determine grid width")?
            .len();
        let height = s.lines().count();

        s.lines()
            .flat_map(str::chars)
            .map(Tile::try_from)
            .collect::<Result<_, _>>()
            .map(|tiles| Grid {
                width,
                height,
                tiles,
            })
    }
}

impl Grid {
    fn get(&self, x: i32, y: i32) -> Option<Tile> {
        let x = usize::try_from(x).ok()?;
        let y = usize::try_from(y).ok()?;

        (x < self.width && y < self.height).then(|| self.tiles[y * self.width + x])
    }

    fn neighbouring_seats(&self, index: usize) -> impl Iterator<Item = Tile> + '_ {
        let y = (index / self.width) as i32;
        let x = (index % self.width) as i32;

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
        .map(move |(dx, dy)| (x + dx, y + dy))
        .filter_map(move |(x, y)| self.get(x, y))
    }

    fn occupied_seats(&self) -> usize {
        self.tiles
            .iter()
            .filter(|tile| matches!(tile, Tile::Seat(true)))
            .count()
    }

    fn visible_seats(&self, index: usize) -> impl Iterator<Item = Tile> + '_ {
        let y = (index / self.width) as i32;
        let x = (index % self.width) as i32;

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
        .map(move |(dx, dy)| {
            itertools::iterate((x + dx, y + dy), move |(x, y)| (x + dx, y + dy))
                .map_while(move |(x, y)| self.get(x, y))
                .skip_while(|t| matches!(t, Tile::Floor))
                .take(1)
        })
        .flatten()
    }

    fn into_iter(self, f: fn(&Grid, usize, Tile) -> Tile) -> GridIterator {
        GridIterator { current: self, f }
    }
}

struct GridIterator {
    current: Grid,
    f: fn(&Grid, usize, Tile) -> Tile,
}

impl Iterator for GridIterator {
    type Item = Grid;

    fn next(&mut self) -> Option<Self::Item> {
        let next_tiles = self
            .current
            .tiles
            .iter()
            .enumerate()
            .map(|(index, &tile)| (self.f)(&self.current, index, tile))
            .collect();

        if self.current.tiles != next_tiles {
            self.current.tiles = next_tiles;
            Some(self.current.clone())
        } else {
            None
        }
    }
}

fn part1_visibility(grid: &Grid, index: usize, tile: Tile) -> Tile {
    let occupied = grid
        .neighbouring_seats(index)
        .filter(|n| matches!(n, Tile::Seat(true)))
        .count();

    match tile {
        Tile::Seat(false) if occupied == 0 => Tile::Seat(true),
        Tile::Seat(true) if occupied >= 4 => Tile::Seat(false),
        _ => tile,
    }
}

fn part2_visibility(grid: &Grid, index: usize, tile: Tile) -> Tile {
    let occupied = grid
        .visible_seats(index)
        .filter(|n| matches!(n, Tile::Seat(true)))
        .count();

    match tile {
        Tile::Seat(false) if occupied == 0 => Tile::Seat(true),
        Tile::Seat(true) if occupied >= 5 => Tile::Seat(false),
        _ => tile,
    }
}

#[aoc_generator(day11)]
fn generator(input: &str) -> Result<Grid> {
    Grid::from_str(&input)
}

#[aoc(day11, part1)]
fn part1(data: &Grid) -> Option<usize> {
    data.clone()
        .into_iter(part1_visibility)
        .last()
        .map(|grid| grid.occupied_seats())
}

#[aoc(day11, part2)]
fn part2(data: &Grid) -> Option<usize> {
    data.clone()
        .into_iter(part2_visibility)
        .last()
        .map(|grid| grid.occupied_seats())
}

#[cfg(test)]
mod tests {
    use super::*;
    use indoc::indoc;
    use Tile::*;

    #[test]
    fn it_parses_input() {
        let data = generator(indoc!(
            "
            L.LL.LL.LL
            LLLLLLL.LL
            L.L.L..L..
            LLLL.LL.LL
            L.LL.LL.LL
            L.LLLLL.LL
            ..L.L.....
            LLLLLLLLLL
            L.LLLLLL.L
            L.LLLLL.LL
            "
        ))
        .expect("input to be parsed");

        assert_eq!(
            Grid {
                width: 10,
                height: 10,
                tiles: vec![
                    Seat(false),
                    Floor,
                    Seat(false),
                    Seat(false),
                    Floor,
                    Seat(false),
                    Seat(false),
                    Floor,
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Floor,
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Floor,
                    Seat(false),
                    Floor,
                    Seat(false),
                    Floor,
                    Floor,
                    Seat(false),
                    Floor,
                    Floor,
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Floor,
                    Seat(false),
                    Seat(false),
                    Floor,
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Floor,
                    Seat(false),
                    Seat(false),
                    Floor,
                    Seat(false),
                    Seat(false),
                    Floor,
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Floor,
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Floor,
                    Seat(false),
                    Seat(false),
                    Floor,
                    Floor,
                    Seat(false),
                    Floor,
                    Seat(false),
                    Floor,
                    Floor,
                    Floor,
                    Floor,
                    Floor,
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Floor,
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Seat(false),
                    Floor,
                    Seat(false),
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
            },
            data
        );
    }

    #[test]
    fn it_solves_part1() {
        let data = generator(indoc!(
            "
            L.LL.LL.LL
            LLLLLLL.LL
            L.L.L..L..
            LLLL.LL.LL
            L.LL.LL.LL
            L.LLLLL.LL
            ..L.L.....
            LLLLLLLLLL
            L.LLLLLL.L
            L.LLLLL.LL
            "
        ))
        .expect("input to be parsed");

        let mut grid = data.into_iter(part1_visibility);

        assert_eq!(
            generator(indoc!(
                "
                #.##.##.##
                #######.##
                #.#.#..#..
                ####.##.##
                #.##.##.##
                #.#####.##
                ..#.#.....
                ##########
                #.######.#
                #.#####.##
                "
            ))
            .ok(),
            grid.next()
        );

        assert_eq!(
            generator(indoc!(
                "
                #.LL.L#.##
                #LLLLLL.L#
                L.L.L..L..
                #LLL.LL.L#
                #.LL.LL.LL
                #.LLLL#.##
                ..L.L.....
                #LLLLLLLL#
                #.LLLLLL.L
                #.#LLLL.##
                "
            ))
            .ok(),
            grid.next()
        );

        assert_eq!(
            generator(indoc!(
                "
                #.##.L#.##
                #L###LL.L#
                L.#.#..#..
                #L##.##.L#
                #.##.LL.LL
                #.###L#.##
                ..#.#.....
                #L######L#
                #.LL###L.L
                #.#L###.##
                "
            ))
            .ok(),
            grid.next()
        );

        assert_eq!(
            generator(indoc!(
                "
                #.#L.L#.##
                #LLL#LL.L#
                L.L.L..#..
                #LLL.##.L#
                #.LL.LL.LL
                #.LL#L#.##
                ..L.L.....
                #L#LLLL#L#
                #.LLLLLL.L
                #.#L#L#.##
                "
            ))
            .ok(),
            grid.next()
        );

        assert_eq!(
            generator(indoc!(
                "
                #.#L.L#.##
                #LLL#LL.L#
                L.#.L..#..
                #L##.##.L#
                #.#L.LL.LL
                #.#L#L#.##
                ..L.L.....
                #L#L##L#L#
                #.LLLLLL.L
                #.#L#L#.##
                "
            ))
            .ok(),
            grid.next()
        );
    }

    #[test]
    fn it_solves_part2() {
        let data = generator(indoc!(
            "
            L.LL.LL.LL
            LLLLLLL.LL
            L.L.L..L..
            LLLL.LL.LL
            L.LL.LL.LL
            L.LLLLL.LL
            ..L.L.....
            LLLLLLLLLL
            L.LLLLLL.L
            L.LLLLL.LL
            "
        ))
        .expect("input to be parsed");

        let mut grid = data.into_iter(part2_visibility);

        assert_eq!(
            generator(indoc!(
                "
                #.##.##.##
                #######.##
                #.#.#..#..
                ####.##.##
                #.##.##.##
                #.#####.##
                ..#.#.....
                ##########
                #.######.#
                #.#####.##
                "
            ))
            .ok(),
            grid.next()
        );

        assert_eq!(
            generator(indoc!(
                "
                #.LL.LL.L#
                #LLLLLL.LL
                L.L.L..L..
                LLLL.LL.LL
                L.LL.LL.LL
                L.LLLLL.LL
                ..L.L.....
                LLLLLLLLL#
                #.LLLLLL.L
                #.LLLLL.L#
                "
            ))
            .ok(),
            grid.next()
        );

        assert_eq!(
            generator(indoc!(
                "
                #.L#.##.L#
                #L#####.LL
                L.#.#..#..
                ##L#.##.##
                #.##.#L.##
                #.#####.#L
                ..#.#.....
                LLL####LL#
                #.L#####.L
                #.L####.L#
                "
            ))
            .ok(),
            grid.next()
        );

        assert_eq!(
            generator(indoc!(
                "
                #.L#.L#.L#
                #LLLLLL.LL
                L.L.L..#..
                ##LL.LL.L#
                L.LL.LL.L#
                #.LLLLL.LL
                ..L.L.....
                LLLLLLLLL#
                #.LLLLL#.L
                #.L#LL#.L#
                "
            ))
            .ok(),
            grid.next()
        );

        assert_eq!(
            generator(indoc!(
                "
                #.L#.L#.L#
                #LLLLLL.LL
                L.L.L..#..
                ##L#.#L.L#
                L.L#.#L.L#
                #.L####.LL
                ..#.#.....
                LLL###LLL#
                #.LLLLL#.L
                #.L#LL#.L#
                "
            ))
            .ok(),
            grid.next()
        );

        assert_eq!(
            generator(indoc!(
                "
                #.L#.L#.L#
                #LLLLLL.LL
                L.L.L..#..
                ##L#.#L.L#
                L.L#.LL.L#
                #.LLLL#.LL
                ..#.L.....
                LLL###LLL#
                #.LLLLL#.L
                #.L#LL#.L#
                "
            ))
            .ok(),
            grid.next()
        );
    }
}
