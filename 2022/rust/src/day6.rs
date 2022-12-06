use std::collections::HashSet;

pub fn run(input: &str) {
    println!("Day 4");
    println!("Part 1: {}", part1(input));
    println!("Part 2: {}", part2(input));
}

fn get_distinct_chars_index(input: &str, window_size: usize) -> usize {
    // input is a string of lowercase letters
    // we need to window every 4 letters and check if they are different
    // if they are different we will return the index of the 4th letter
    // start by splitting the string into a vector of chars
    let chars = input.chars().collect::<Vec<char>>();
    let mut index = window_size;
    for window in chars.windows(window_size) {
        // validate non of the letters are the same in the window
        let window_len = window.len();
        let window_set = window.iter().collect::<HashSet<&char>>();
        let window_set_len = window_set.len();
        if window_len == window_set_len {
            break;
        }
        index += 1;
    }
    index
}

fn part1(input: &str) -> usize {
    get_distinct_chars_index(input, 4)
}

fn part2(input: &str) -> usize {
    get_distinct_chars_index(input, 14)
}

//test
#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        let inputs = vec![
            ("mjqjpqmgbljsphdztnvjfqwrcgsmlb", 7),
            ("bvwbjplbgvbhsrlpgdmjqwftvncz", 5),
            ("nppdvjthqldpwncqszvftbrmjlhg", 6),
            ("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10),
            ("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11),
        ];

        for input in inputs.into_iter() {
            println!("input: {:?}", input);
            assert_eq!(part1(input.0), input.1);
        }
    }

    #[test]
    fn test_part2() {
        let inputs = vec![
            ("mjqjpqmgbljsphdztnvjfqwrcgsmlb", 19),
            ("bvwbjplbgvbhsrlpgdmjqwftvncz", 23),
            ("nppdvjthqldpwncqszvftbrmjlhg", 23),
            ("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 29),
            ("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 26),
        ];

        for input in inputs.into_iter() {
            println!("input: {:?}", input);
            assert_eq!(part2(input.0), input.1);
        }
    }
}
