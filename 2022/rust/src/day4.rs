pub fn run(input: &str) {
    println!("Day 4");
    println!("Part 1: {}", part1(input));
    println!("Part 2: {}", part2(input));
}

fn read_input(input: &str) -> Vec<(Vec<u64>, Vec<u64>)> {
    input
        .lines()
        .map(|line| {
            let mut iter = line.split_terminator(',');
            let mut range1 = iter.next().unwrap().split_terminator('-');
            let mut range2 = iter.next().unwrap().split_terminator('-');
            // see if the larger range contains the smaller range
            let range1_start = range1.next().unwrap().parse::<u64>().unwrap();
            let range1_end = range1.next().unwrap().parse::<u64>().unwrap();
            let range2_start = range2.next().unwrap().parse::<u64>().unwrap();
            let range2_end = range2.next().unwrap().parse::<u64>().unwrap();
            (
                vec![range1_start, range1_end],
                vec![range2_start, range2_end],
            )
        })
        .collect::<Vec<(Vec<u64>, Vec<u64>)>>()
}

fn part1(input: &str) -> u64 {
    let ranges = read_input(input);
    let mut overlaps = 0;
    for range in ranges.iter() {
        let range1 = range.0.clone();
        let range2 = range.1.clone();
        let (larger_range, smaller_range) = if range1[1] - range1[0] > range2[1] - range2[0] {
            (range1, range2)
        } else {
            (range2, range1)
        };
        // if the larger range contains all of the smaller range +1, then there is an overlap
        if larger_range[0] <= smaller_range[0] && larger_range[1] >= smaller_range[1] {
            overlaps += 1;
        }
    }
    overlaps
}

fn part2(input: &str) -> u64 {
    let ranges = read_input(input);
    let mut overlaps = 0;
    for range in ranges.iter() {
        let range1 = range.0.clone();
        let range2 = range.1.clone();
        // if range 1 contains any numbers in range 2, or vice versa, then there is an overlap
        if range1[0] <= range2[0] && range1[1] >= range2[0] {
            overlaps += 1;
        } else if range2[0] <= range1[0] && range2[1] >= range1[0] {
            overlaps += 1;
        } else if range1[1] >= range2[1] && range1[0] <= range2[1] {
            overlaps += 1;
        } else if range2[1] >= range1[1] && range2[0] <= range1[1] {
            overlaps += 1;
        }
    }
    overlaps
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        let input = "2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8";
        assert_eq!(part1(input), 2);
    }

    #[test]
    fn test_part2() {
        let input = "2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8";
        assert_eq!(part2(input), 4);
    }
}
