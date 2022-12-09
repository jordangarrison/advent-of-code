pub fn run(input: &str) {
    println!("Day 9");
    println!("Part 1: {}", part1(input));
    println!("Part 2: {}", part2(input));
}

fn read_input(input: &str) -> Vec<(&str, i32)> {
    input
        .lines()
        .map(|line| {
            let mut iter = line.split_whitespace();
            let direction = iter.next().unwrap();
            let distance = iter.next().unwrap();
            (direction, distance.parse::<i32>().unwrap())
        })
        .collect::<Vec<(&str, i32)>>()
}

fn calculate_tail(head: (i32, i32), tail: (i32, i32)) -> (i32, i32) {
    let x = head.0 - tail.0;
    let y = head.1 - tail.1;
    let mut tail = tail;
    println!("x: {}, y: {}", x, y);
    tail.0 = match x {
        -1 | 0 | 1 => tail.0,
        _ => if x > 0 {
            tail.0 + x -1
        } else {
            tail.0 - x + 1
        },
    };
    tail.1 = match y {
        -1 |  0 | 1 => tail.1,
        _ => if y > 0 {
            tail.1 + y - 1
        } else {
            tail.1 - y + 1
        },
    };
    tail
}

fn part1(input: &str) -> u64 {
    let instructions = read_input(input);
    println!("{:?}", instructions);
    let mut head_pos: (i32, i32) = (0, 0);
    let mut tail_pos: (i32, i32) = (0, 0);
    let mut tracked: Vec<(i32, i32)> = vec![];
    for instruction in instructions.into_iter() {
        let direction = instruction.0;
        let distance = instruction.1;
        for i in 0..distance {
            if direction == "L" {
                head_pos.0 -= 1;
            } else if direction == "R" {
                head_pos.0 += 1;
            } else if direction == "U" {
                head_pos.1 += 1;
            } else if direction == "D" {
                head_pos.1 -= 1;
            }
            tail_pos = calculate_tail(head_pos, tail_pos);
            tracked.push(tail_pos);
            println!("Instruction: {:?}, Head: {:?}, Tail: {:?}", instruction, head_pos, tail_pos);
        }
    }
    println!("Tracked: {:?}", tracked);

    13
}

fn part2(input: &str) -> u64 {
    0
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        let input = "R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2";
        assert_eq!(part1(input), 13)
    }
}