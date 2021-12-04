pub fn run(input: &str) {
    println!("Day 1");
    println!("Part 1: {}", part1(input));
}

fn part1(input: &str) -> u64 {
    let numbers: Vec<u64> = input
        .split_whitespace()
        .map(|s| s.parse::<u64>().unwrap())
        .collect();
    let increase = numbers.windows(2).filter(|w| w[0] < w[1]).next().unwrap();
    increase[0] + increase[1]
}
