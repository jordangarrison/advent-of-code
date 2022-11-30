pub fn run(input: &str) {
    println!("Day 1");
    println!("Part 1: {}", part1(input));
    println!("Part 2: {}", part2(input));
}

fn parse_numbers(input: &str) -> Vec<u64> {
    // split the input string into a vector of numbers
    input
        .split_whitespace()
        .map(|s| s.parse::<u64>().unwrap())
        .collect()
}

fn part1(input: &str) -> u64 {
    // split the input string into a vector of numbers
    let numbers: Vec<u64> = parse_numbers(input);

    // calclate the numbers which are greater than the previous number
    numbers.windows(2).filter(|w| w[0] < w[1]).count() as u64
}

fn part2(input: &str) -> u64 {
    // split the input string into a vector of numbers
    let numbers: Vec<u64> = parse_numbers(input);

    // calculate the increases in sums of 3 numbers in the vector with the next 3 numbers in the vector one index ahead
    numbers
        .windows(3)
        .zip(numbers.windows(3).skip(1))
        .filter(|(a, b)| a.iter().sum::<u64>() < b.iter().sum::<u64>())
        .count() as u64
}

// test of part 1 of day 1
#[cfg(test)]
mod day1_tests {
    use super::*;

    #[test]
    fn test_part1() {
        let input = "199\n\
          200\n\
          208\n\
          210\n\
          200\n\
          207\n\
          240\n\
          269\n\
          260\n\
          263\n";
        assert_eq!(part1(input), 7);
    }

    #[test]
    fn test_part2() {
        let input = "199\n\
          200\n\
          208\n\
          210\n\
          200\n\
          207\n\
          240\n\
          269\n\
          260\n\
          263\n";
        assert_eq!(part2(input), 5);
    }
}
