use std::collections::HashMap;

pub fn run(input: &str) {
    println!("Day 7");
    let (part1, part2) = problem(input);
    println!("Part 1: {}", part1);
    println!("Part 2: {}", part2);
}

fn current_dir_formatted(current_dir: &str) -> String {
    if current_dir == "/" {
        "".to_string()
    } else {
        format!("{}", current_dir)
    }
}

fn problem(input: &str) -> (u64, u64) {
    // find all directories with a size of at least 100000
    // set up recursive hashmap
    let mut dir_map: HashMap<String, u64> = HashMap::new();
    dir_map.insert("/".to_string(), 0);
    let mut file_map: HashMap<String, u64> = HashMap::new();
    let mut current_dir = String::from("/");
    input.lines().into_iter().for_each(|line| {
        // if the line starts with a $ then its a command
        // the available commands are ls and cd
        if line.starts_with("$") {
            let command = line.split_whitespace().nth(1).unwrap();
            match command {
                "cd" => {
                    let dir = line.split_whitespace().nth(2).unwrap();

                    if dir == "/" {
                        current_dir = String::from("/");
                    } else if dir == "." {
                        // do nothing
                    } else if dir == ".." {
                        // go up a directory
                        let mut split = current_dir.split("/").collect::<Vec<&str>>();
                        split.pop();
                        current_dir = split.join("/");
                    } else {
                        current_dir = format!("{}/{}", current_dir_formatted(&current_dir), dir);
                    }
                }
                "ls" => {
                    // do nothing
                }
                _ => {
                    println!("Unknown command: {}", command);
                }
            }
            return;
        }
        // if the line starts with a number then its a file
        // if the line starts with a dir then its a directory
        if line.starts_with("dir") {
            let dir = line.split_whitespace().nth(1).unwrap();
            dir_map.insert(
                format!("{}/{}", current_dir_formatted(&current_dir), dir),
                0,
            );
        } else {
            let size = line
                .split_whitespace()
                .nth(0)
                .unwrap()
                .parse::<u64>()
                .unwrap();
            let file = line.split_whitespace().nth(1).unwrap();
            file_map.insert(
                format!("{}/{}", current_dir_formatted(&current_dir), file),
                size,
            );
        }
    });
    // if the file in the file map starts with a dir in the dir map then add the size to the dir map
    file_map.iter().for_each(|(file, size)| {
        dir_map.iter_mut().for_each(|(dir, dir_size)| {
            if file.starts_with(dir) {
                *dir_size += size;
            }
        });
    });

    // find all directories smaller than 100000 and sum their sizes
    let part1 = dir_map.iter().fold(
        0,
        |acc, (_, size)| {
            if *size <= 100000 {
                acc + size
            } else {
                acc
            }
        },
    );
    // sum all the files
    let total_size = file_map.iter().fold(0, |acc, (_, size)| acc + size);

    // Find the available space on the device
    let fs_size = 70000000;
    let required_space = 30000000;
    let available_space = fs_size - total_size;
    let needed_space = required_space - available_space;

    // find the directories larger than the needed space
    let part2 = dir_map
        .into_iter()
        .filter(|(_, size)| *size >= needed_space)
        .map(|(_, size)| size)
        .fold(0, |acc, size| if acc == 0 { size } else { acc.min(size) });
    (part1, part2)
}

// tests
#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_all() {
        let input = "$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k";
        let solution = problem(input);
        assert_eq!(solution, (95437, 24933642));
    }
}
