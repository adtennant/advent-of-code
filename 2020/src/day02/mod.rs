use anyhow::{Context, Result};
use parse_display::{Display, FromStr};
use std::str::FromStr;

#[derive(Display, Debug, FromStr, PartialEq)]
#[display("{left}-{right} {letter}")]
struct PasswordPolicy {
    left: usize,
    right: usize,
    letter: char,
}

impl PasswordPolicy {
    fn validate_corporate_policy(&self, password: &str) -> bool {
        let character_count = password.matches(self.letter).count();

        character_count >= self.left && character_count <= self.right
    }

    fn validate_authentication_system(&self, password: &str) -> bool {
        let first_index = self.left - 1;
        let second_index = self.right - 1;

        (password.chars().nth(first_index) == Some(self.letter))
            ^ (password.chars().nth(second_index) == Some(self.letter))
    }
}

#[derive(Display, Debug, FromStr, PartialEq)]
#[display("{policy}: {password}")]
struct PasswordEntry {
    policy: PasswordPolicy,
    password: String,
}

impl PasswordEntry {
    fn matches_corporate_policy(&self) -> bool {
        self.policy.validate_corporate_policy(&self.password)
    }

    fn matches_authentication_system(&self) -> bool {
        self.policy.validate_authentication_system(&self.password)
    }
}

#[aoc_generator(day2)]
fn generator(data: &str) -> Result<Vec<PasswordEntry>> {
    data.lines()
        .map(|line| PasswordEntry::from_str(line).with_context(|| "parsing input failed"))
        .collect()
}

#[aoc(day2, part1)]
fn part1(input: &[PasswordEntry]) -> usize {
    input
        .iter()
        .filter(|entry| entry.matches_corporate_policy())
        .count()
}

#[aoc(day2, part2)]
fn part2(input: &[PasswordEntry]) -> usize {
    input
        .iter()
        .filter(|entry| entry.matches_authentication_system())
        .count()
}

#[cfg(test)]
mod tests {
    use super::*;
    use indoc::indoc;

    #[test]
    fn it_parses_a_password_policy() {
        let data = "1-3 a";
        let policy = PasswordPolicy::from_str(data).unwrap();

        assert_eq!(
            PasswordPolicy {
                left: 1,
                right: 3,
                letter: 'a'
            },
            policy
        )
    }

    #[test]
    fn it_parses_a_password_entry() {
        let data = "1-3 a: abcde";
        let entry = PasswordEntry::from_str(data).unwrap();

        assert_eq!(
            PasswordEntry {
                policy: PasswordPolicy {
                    left: 1,
                    right: 3,
                    letter: 'a'
                },
                password: String::from("abcde")
            },
            entry
        )
    }

    #[test]
    fn it_checks_passwords_match_corporate_policy() {
        let data = "1-3 a: abcde";
        let entry = PasswordEntry::from_str(data).unwrap();

        assert_eq!(true, entry.matches_corporate_policy());

        let data = "1-3 b: cdefg";
        let entry = PasswordEntry::from_str(data).unwrap();

        assert_eq!(false, entry.matches_corporate_policy());

        let data = "2-9 c: ccccccccc";
        let entry = PasswordEntry::from_str(data).unwrap();

        assert_eq!(true, entry.matches_corporate_policy());
    }

    #[test]
    fn it_checks_passwords_match_authentication_system() {
        let data = "1-3 a: abcde";
        let entry = PasswordEntry::from_str(data).unwrap();

        assert_eq!(true, entry.matches_authentication_system());

        let data = "1-3 b: cdefg";
        let entry = PasswordEntry::from_str(data).unwrap();

        assert_eq!(false, entry.matches_authentication_system());

        let data = "2-9 c: ccccccccc";
        let entry = PasswordEntry::from_str(data).unwrap();

        assert_eq!(false, entry.matches_authentication_system());
    }

    #[test]
    fn it_parses_input() {
        let data = indoc! {"
            1-3 a: abcde
            1-3 b: cdefg
            2-9 c: ccccccccc
        "};

        let input = generator(data).unwrap();

        assert_eq!(
            vec![
                PasswordEntry {
                    policy: PasswordPolicy {
                        left: 1,
                        right: 3,
                        letter: 'a'
                    },
                    password: String::from("abcde")
                },
                PasswordEntry {
                    policy: PasswordPolicy {
                        left: 1,
                        right: 3,
                        letter: 'b'
                    },
                    password: String::from("cdefg")
                },
                PasswordEntry {
                    policy: PasswordPolicy {
                        left: 2,
                        right: 9,
                        letter: 'c'
                    },
                    password: String::from("ccccccccc")
                }
            ],
            input
        )
    }
}
