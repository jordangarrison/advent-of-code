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
    let path = format!("./data/day{}.txt", day);
    read_file(&path)
}

// Pretty print a time in seconds with appropriate units to 3 decimal places
pub fn pretty_print_time(time: f64) -> String {
    if time < 1e-6 {
        format!("{:.3} ns", time * 1e9)
    } else if time < 1e-3 {
        format!("{:.3} μs", time * 1e6)
    } else if time < 1.0 {
        format!("{:.3} ms", time * 1e3)
    } else {
        format!("{:.3} s", time)
    }
}

// Timing function for execution of another function
pub fn time_execution<F>(f: F)
where
    F: FnOnce(),
{
    let start = std::time::Instant::now();
    f();
    // pretty print the execution time with appropriate units to 3 decimal places
    println!(
        "Execution time: {}",
        pretty_print_time(start.elapsed().as_secs_f64())
    );
}
