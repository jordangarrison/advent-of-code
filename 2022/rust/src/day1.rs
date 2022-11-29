pub fn run(input: &str) {
    println!("Day 1");
    println!("Part 1: {}", part1(input));
}

fn part1(input: &str) -> u64 {
    input.lines().map(|line| line.parse::<u64>().unwrap()).sum()
}
