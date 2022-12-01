pub fn run(input: &str) {
    println!("Day 1");
    println!("Part 1: {}", part1(input));
    println!("Part 2: {}", part2(input));
}

fn parse_input(input: &str) -> Vec<Vec<u64>> {
    // Create the groups of snacks by splitting on the empty line
    let elfs = input.split("\n\n");

    // for each elf in elfs create a vector of u64 to represent the snacks
    elfs.map(|elf| {
        elf.lines()
            .map(|line| line.parse::<u64>().unwrap())
            .collect()
    })
    .collect()
}

fn part1(input: &str) -> u64 {
    let snacks: Vec<Vec<u64>> = parse_input(input);

    // calculate the largeest summed vector of u64 in snacks
    snacks
        .iter()
        .map(|snack| snack.iter().sum::<u64>())
        .max()
        .unwrap()
}

fn part2(input: &str) -> u64 {
    let snacks: Vec<Vec<u64>> = parse_input(input);

    // sum values in each vector of u64 in snacks and create a new vector of u64
    // this will allow us to operate on each set of snacks as summed calories instead of
    // individual snacks
    let mut sums: Vec<u64> = snacks
        .iter()
        .map(|snack| snack.iter().sum::<u64>())
        .collect();

    // return the 3 largest values in the new vector of u64 to get the solution of the largest
    // sum of 3 snacks
    sums.sort();
    sums[sums.len() - 3..].iter().sum()
}

// tests
#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        let input = "1000
2000
3000

4000

5000
6000

7000
8000
9000

10000";
        assert_eq!(part1(input), 24000);
    }

    #[test]
    fn test_part2() {
        let input = "1000
2000
3000

4000

5000
6000

7000
8000
9000

10000";
        assert_eq!(part2(input), 45000);
    }
}
