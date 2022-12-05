use std::collections::{HashMap, HashSet};

pub fn run(input: &str) {
    println!("Day 3");
    println!("Part 1: {}", part1(input));
    println!("Part 2: {}", part2(input));
}

fn get_priorities() -> HashMap<char, u64> {
    // To help prioritize item rearrangement, every item type can be converted to a priority:
    // Lowercase item types a through z have priorities 1 through 26.
    // Uppercase item types A through Z have priorities 27 through 52.
    let mut priorities = HashMap::new();
    for i in 1..27 {
        priorities.insert((i + 96) as u8 as char, i);
    }
    for i in 1..27 {
        priorities.insert((i + 64) as u8 as char, i + 26);
    }
    // make priorities immutable
    let priorities = priorities;
    priorities
}

fn part1(input: &str) -> u64 {
    // read input in line by line to a vector of strings
    let lines = input.lines().collect::<Vec<&str>>();
    // parse each line of the lines into 2 equal character vectors
    let bags = lines
        .iter()
        .map(|line| {
            // split the line into 2 equal strings
            let split_line = line.split_at(line.len() / 2);
            // convert the 2 strings into 2 character vectors
            let vec1 = split_line.0.chars().collect::<Vec<char>>();
            let vec2 = split_line.1.chars().collect::<Vec<char>>();
            (vec1, vec2)
        })
        .collect::<Vec<(Vec<char>, Vec<char>)>>();

    let priorities = get_priorities();

    bags.iter()
        .map(|(c1, c2)| {
            // find the matching character between the 2 vectors and return the character or a space
            // convert character to a number
            let findings = c1
                .iter()
                .map(|c1| {
                    // iterate over c2 and find the matching character and return the character or a space
                    c2.iter().find(|c2| c1 == c2.clone()).unwrap_or(&' ')
                })
                .map(|c| c)
                .filter(|c| c.clone() != &' ');
            // dedupe the findings
            let dedupe_findings: HashSet<&char> = findings.collect();
            // convert the findings to a number
            dedupe_findings
                .iter()
                .map(|c| priorities.get(c.clone()).unwrap())
                .sum::<u64>()
        })
        .sum()
}

fn part2(input: &str) -> u64 {
    let priorities = get_priorities();

    // Read input in 3 lines at a time to a vector of strings
    let groups = input
        .lines()
        .collect::<Vec<&str>>()
        .chunks(3)
        .map(|group| group.join("\n"))
        .collect::<Vec<String>>();
    // within each group create a vector of strings line by line
    let groups_of_elves = groups
        .iter()
        .map(|group| group.lines().collect::<Vec<&str>>())
        .collect::<Vec<Vec<&str>>>();
    // for each group of lines find the matching character between all 3 lines
    groups_of_elves
        .iter()
        .map(|group| {
            // convert the 3 lines into 3 character vectors
            let vec1 = group[0].chars().collect::<Vec<char>>();
            let vec2 = group[1].chars().collect::<Vec<char>>();
            let vec3 = group[2].chars().collect::<Vec<char>>();
            // find the matching character between the 3 vectors and return the character or a space
            let findings = vec1
                .iter()
                .map(|c1| {
                    // iterate over c2 and find the matching character and return the character or a space
                    vec2.iter().find(|c2| c1 == c2.clone()).unwrap_or(&' ')
                })
                .map(|c| {
                    // iterate over c3 and find the matching character and return the character or a space
                    vec3.iter().find(|c3| c == c3.clone()).unwrap_or(&' ')
                })
                .filter(|c| c.clone() != &' ');
            // dedupe the findings
            let dedupe_findings: HashSet<&char> = findings.collect();
            // convert the findings to a number
            dedupe_findings
                .iter()
                .map(|c| priorities.get(c.clone()).unwrap())
                .sum::<u64>()
        })
        .into_iter()
        .sum()
}

// tests
#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        let input = "vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw";
        assert_eq!(part1(input), 157);
    }

    #[test]
    fn test_part2() {
        let input = "vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw";
        assert_eq!(part2(input), 70);
    }
}
