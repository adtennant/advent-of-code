use anyhow::{bail, Context, Result};
use std::str::FromStr;

#[derive(Debug, PartialEq)]
struct PasswordPolicy(usize, usize, char);

impl FromStr for PasswordPolicy {
    type Err = anyhow::Error;

    fn from_str(s: &str) -> Result<Self> {
        let policy: Vec<&str> = s
            .split(|c| c == '-' || c == ' ')
            .map(|l| l.trim())
            .collect();

        if policy.len() != 3 {
            bail!("policy has an invalid number of parts: {}", policy.len());
        }

        let min_occurences = policy[0]
            .parse::<usize>()
            .with_context(|| "failed to parse password policy: invalid first digit")?;
        let max_occurences = policy[1]
            .parse::<usize>()
            .with_context(|| "failed to parse password policy: invalid second digit")?;
        let character = policy[2]
            .chars()
            .next()
            .with_context(|| "failed to parse password policy: invalid character")?;

        Ok(PasswordPolicy(min_occurences, max_occurences, character))
    }
}

impl PasswordPolicy {
    fn matches_corporate_policy(&self, password: &str) -> bool {
        let character_count = password.matches(self.2).count();

        character_count >= self.0 && character_count <= self.1
    }

    fn matches_authentication_system(&self, password: &str) -> bool {
        let first_index = self.0 - 1;
        let second_index = self.1 - 1;

        let match_count = password
            .chars()
            .enumerate()
            .filter_map(|(i, c)| {
                if i == first_index || i == second_index {
                    Some(c)
                } else {
                    None
                }
            })
            .filter(|&c| c == self.2)
            .count();

        match_count == 1
    }
}

#[derive(Debug, PartialEq)]
struct PasswordEntry {
    policy: PasswordPolicy,
    password: String,
}

impl FromStr for PasswordEntry {
    type Err = anyhow::Error;

    fn from_str(s: &str) -> Result<Self> {
        let entry: Vec<&str> = s.split(": ").map(|l| l.trim()).collect();

        if entry.len() != 2 {
            bail!("entry has an invalid number of parts: {}", entry.len());
        }

        let policy =
            PasswordPolicy::from_str(entry[0]).with_context(|| "failed to parse password entry")?;
        let password = entry[1].to_owned();

        Ok(PasswordEntry { policy, password })
    }
}

impl PasswordEntry {
    fn matches_corporate_policy(&self) -> bool {
        self.policy.matches_corporate_policy(&self.password)
    }

    fn matches_authentication_system(&self) -> bool {
        self.policy.matches_authentication_system(&self.password)
    }
}

fn parse_input(data: &str) -> Result<Vec<PasswordEntry>> {
    data.lines()
        .map(|line| line.trim())
        .filter(|line| !line.is_empty())
        .map(|line| PasswordEntry::from_str(line))
        .collect()
}

fn main() -> Result<()> {
    println!("Day 2: Password Philosophy");

    let parsed_input = parse_input(include_str!("../input.txt"))?;

    println!(
        "Part 1: {:?}",
        parsed_input
            .iter()
            .filter(|entry| entry.matches_corporate_policy())
            .count()
    );

    println!(
        "Part 2: {:?}",
        parsed_input
            .iter()
            .filter(|entry| entry.matches_authentication_system())
            .count()
    );

    Ok(())
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_parses_a_password_policy() {
        let data = "1-3 a";
        let policy = PasswordPolicy::from_str(data).unwrap();

        assert_eq!(PasswordPolicy(1, 3, 'a'), policy)
    }

    #[test]
    fn it_parses_a_password_entry() {
        let data = "1-3 a: abcde";
        let entry = PasswordEntry::from_str(data).unwrap();

        assert_eq!(
            PasswordEntry {
                policy: PasswordPolicy(1, 3, 'a'),
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
        let data = "
            1-3 a: abcde
            1-3 b: cdefg
            2-9 c: ccccccccc
        ";

        let parsed_input = parse_input(data).unwrap();

        assert_eq!(
            vec![
                PasswordEntry {
                    policy: PasswordPolicy(1, 3, 'a'),
                    password: String::from("abcde")
                },
                PasswordEntry {
                    policy: PasswordPolicy(1, 3, 'b'),
                    password: String::from("cdefg")
                },
                PasswordEntry {
                    policy: PasswordPolicy(2, 9, 'c'),
                    password: String::from("ccccccccc")
                }
            ],
            parsed_input
        )
    }
}
