use anyhow::{bail, Result};
use itertools::Itertools;
use regex::Regex;
use std::{collections::HashMap, convert::TryFrom};

#[derive(Debug, PartialEq)]
struct Passport<'a>(HashMap<&'a str, &'a str>);

impl<'a> TryFrom<&'a str> for Passport<'a> {
    type Error = anyhow::Error;

    fn try_from(s: &'a str) -> Result<Self, Self::Error> {
        let passport: HashMap<_, _> = s
            .split_whitespace()
            .flat_map(|p| p.split(':'))
            .tuples()
            .collect();

        Ok(Passport(passport))
    }
}

impl<'a> Passport<'a> {
    fn has_required_fields(&self) -> bool {
        ["byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"]
            .iter()
            .all(|&k| self.0.contains_key(k))
    }

    fn is_valid(&self) -> bool {
        if !self.has_required_fields() {
            return false;
        }

        self.0.iter().all(|(&k, &v)| match k {
            "byr" => v
                .parse::<i32>()
                .map(|v| (1920..=2002).contains(&v))
                .unwrap_or_default(),
            "iyr" => v
                .parse::<i32>()
                .map(|v| (2010..=2020).contains(&v))
                .unwrap_or_default(),
            "eyr" => v
                .parse::<i32>()
                .map(|v| (2020..=2030).contains(&v))
                .unwrap_or_default(),
            "hgt" => match v.split_at(v.len() - 2) {
                (v, "cm") => v
                    .parse::<i32>()
                    .map(|v| (150..=193).contains(&v))
                    .unwrap_or_default(),
                (v, "in") => v
                    .parse::<i32>()
                    .map(|v| (59..=76).contains(&v))
                    .unwrap_or_default(),
                _ => false,
            },
            "hcl" => {
                v.starts_with("#")
                    && v.len() == 7
                    && v.chars().skip(1).all(|c| c.to_digit(16).is_some())
            }
            "ecl" => ["amb", "blu", "brn", "gry", "grn", "hzl", "oth"].contains(&v),
            "pid" => v.len() == 9 && v.chars().all(|c| c.is_numeric()),
            _ => true,
        })
    }
}

fn parse_input(input: &str) -> Result<Vec<Passport>> {
    let regex = Regex::new("\n *\n")?;

    regex.split(input).map(Passport::try_from).collect()
}

fn main() -> Result<()> {
    println!("Day 4: Passport Processing");

    let input = parse_input(include_str!("../input.txt"))?;

    println!(
        "Part 1: {}",
        input
            .iter()
            .filter(|passport| passport.has_required_fields())
            .count()
    );

    println!(
        "Part 2: {}",
        input.iter().filter(|passport| passport.is_valid()).count()
    );

    Ok(())
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_parses_a_passport() {
        let data = "
            ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
            byr:1937 iyr:2017 cid:147 hgt:183cm
        ";
        let passport = Passport::try_from(data).unwrap();

        assert_eq!(
            Passport(
                [
                    ("ecl", "gry"),
                    ("pid", "860033327"),
                    ("eyr", "2020"),
                    ("hcl", "#fffffd"),
                    ("byr", "1937"),
                    ("iyr", "2017"),
                    ("cid", "147"),
                    ("hgt", "183cm")
                ]
                .iter()
                .cloned()
                .collect()
            ),
            passport
        )
    }

    #[test]
    fn it_parses_input() {
        let data = "
            pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980
            hcl:#623a2f
            
            eyr:2029 ecl:blu cid:129 byr:1989
            iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm
            
            hcl:#888785
            hgt:164cm byr:2001 iyr:2015 cid:88
            pid:545766238 ecl:hzl
            eyr:2022
            
            iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719
        ";

        let input = parse_input(data).unwrap();

        assert_eq!(4, input.len())
    }

    #[test]
    fn it_validates_a_passport() {
        let data = "
            eyr:1972 cid:100
            hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926

            iyr:2019
            hcl:#602927 eyr:1967 hgt:170cm
            ecl:grn pid:012533040 byr:1946

            hcl:dab227 iyr:2012
            ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277

            hgt:59cm ecl:zzz
            eyr:2038 hcl:74454a iyr:2023
            pid:3556412378 byr:2007
        ";

        let input = parse_input(data).unwrap();

        assert_eq!(
            0,
            input.iter().filter(|passport| passport.is_valid()).count()
        );

        let data = "
            pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980
            hcl:#623a2f
            
            eyr:2029 ecl:blu cid:129 byr:1989
            iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm
            
            hcl:#888785
            hgt:164cm byr:2001 iyr:2015 cid:88
            pid:545766238 ecl:hzl
            eyr:2022
            
            iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719
        ";

        let input = parse_input(data).unwrap();

        assert_eq!(
            4,
            input.iter().filter(|passport| passport.is_valid()).count()
        )
    }
}
