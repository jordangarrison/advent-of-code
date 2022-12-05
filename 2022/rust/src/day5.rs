use regex::Regex;

pub fn run(input: &str) {
    println!("Day 5");
    println!("Part 1: {}", part1(input));
    println!("Part 2: {}", part2(input));
}

fn transpose(input: Vec<Vec<String>>) -> Vec<Vec<String>> {
    let mut transposed = vec![vec![]; input[0].len()];
    for row in input {
        for (i, col) in row.iter().enumerate() {
            transposed[i].push(col.clone());
        }
    }
    transposed
}

fn get_stacks(input: &str) -> Vec<Vec<String>> {
    // parse stack into a vector of characters by columns
    // create a vector of vectors of characters
    let mut stack = Vec::new();
    for line in input.lines() {
        let mut row = Vec::new();
        for c in line.chars() {
            row.push(c);
        }
        stack.push(row);
    }
    // start by replacing every 4th character with a | character
    for row in stack.iter_mut() {
        for (i, c) in row.iter_mut().enumerate() {
            if i % 4 == 3 {
                *c = '|';
            }
        }
    }
    // rejoin the stack into a string
    let stack_str = stack
        .iter()
        .map(|row| row.iter().collect::<String>())
        .collect::<Vec<String>>()
        .join("\n");

    // parse each column of the stackStr into a vector of characters by '|' delimited columns
    let mut stack = Vec::new();
    for line in stack_str.lines() {
        let mut row = Vec::new();
        for c in line.split('|') {
            // remove '[' and ']' characters from the string
            let c = c.replace("[", "").replace("]", "");
            row.push(c);
        }
        stack.push(row);
    }

    // transpose stack
    let mut stack = transpose(stack);
    // pop the last element from each Vec<String> in the stack
    let stack = stack
        .iter_mut()
        .map(|row| {
            row.pop();
            row
        })
        .collect::<Vec<&mut Vec<String>>>();
    // reverse each row of stacks
    stack
        .iter()
        .map(|row| {
            row.iter()
                .rev()
                .map(|c| c.clone())
                // filter whitespace entries
                .filter(|c| !Regex::new(r"\s").unwrap().is_match(c))
                .collect::<Vec<String>>()
        })
        .collect::<Vec<Vec<String>>>()
}

fn get_instructions(input: &str) -> Vec<Vec<u8>> {
    input
        .lines()
        .map(|line| {
            // parse each line into a vector of u8
            // filter out the words
            line.split_whitespace()
                .filter(|word| word.parse::<u8>().is_ok())
                .map(|word| word.parse::<u8>().unwrap())
                .collect::<Vec<u8>>()
        })
        .collect::<Vec<Vec<u8>>>()
}

fn get_top_crates(stack: Vec<Vec<String>>) -> String {
    let top_crates = stack
        .into_iter()
        .map(|row| {
            // grab the last element of each row
            row.last().unwrap().clone()
        })
        // join the elements of each row into a string
        .collect::<Vec<String>>();
    // cast the vector of &String to a vector of &str
    // join the top crates into a string
    top_crates
        .iter()
        .map(|c| c.as_str())
        .collect::<Vec<&str>>()
        .join("")
}

fn move_crate(instruction: Vec<u8>, stack: Vec<Vec<String>>) -> Vec<Vec<String>> {
    let num_crate = instruction[0];
    let from = instruction[1] - 1;
    let to = instruction[2] - 1;
    let mut stack = stack;
    // pop num_crates from from stack and push onto to stack
    for _ in 0..num_crate {
        let c = stack[from as usize].pop().unwrap();
        stack[to as usize].push(c);
    }
    stack
}

fn move_multi_crate(instruction: Vec<u8>, stack: Vec<Vec<String>>) -> Vec<Vec<String>> {
    let num_crate = instruction[0];
    let from = instruction[1] - 1;
    let to = instruction[2] - 1;
    let mut stack = stack;
    // pop num_crates from from stack and push onto to stack
    let mut intermediate: Vec<String> = Vec::new();
    for _ in 0..num_crate {
        let c = stack[from as usize].pop().unwrap();
        intermediate.push(c);
    }
    intermediate
        .into_iter()
        .rev()
        .for_each(|c| stack[to as usize].push(c));
    stack
}

fn part1(input: &str) -> String {
    // read input and split on emtpy line
    let data = input.split("\n\n").collect::<Vec<&str>>();
    let stack_data = data[0];
    let instructions_data = data[1];
    let stack = get_stacks(stack_data);

    // parse instructions into a vector of u8
    let instructions = get_instructions(instructions_data);

    // execute instructions
    let mut stack = stack;
    for instruction in instructions {
        stack = move_crate(instruction, stack);
    }
    get_top_crates(stack)
}

fn part2(input: &str) -> String {
    let data = input.split("\n\n").collect::<Vec<&str>>();
    let stack_data = data[0];
    let instructions_data = data[1];
    let stack = get_stacks(stack_data);
    let instructions = get_instructions(instructions_data);

    let mut stack = stack;
    for instruction in instructions {
        stack = move_multi_crate(instruction, stack);
    }
    get_top_crates(stack)
}

// tests
#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        let input = "    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2";
        assert_eq!(part1(input), "CMZ");
    }

    #[test]
    fn test_part2() {
        let input = "    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2";
        assert_eq!(part2(input), "MCD");
    }
}
