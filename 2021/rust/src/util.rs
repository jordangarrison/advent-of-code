use std::{fs::File, io::Read};

// Function which reads a file and returns its contents as a string
fn read_file(path: &str) -> String {
    let mut file = File::open(path).expect("File not found");
    let mut contents = String::new();
    file.read_to_string(&mut contents)
        .expect("Failed to read file");
    contents
}

// Get the contents of a day's input file
pub fn get_day_input(day: u8) -> String {
    let path = format!("../go/data/day{}/part1.txt", day);
    read_file(&path)
}
