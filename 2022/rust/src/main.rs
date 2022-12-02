mod day1;
mod day2;
mod util;
use std::env;

// Timing function for execution of another function
fn time_execution<F>(f: F)
where
    F: FnOnce(),
{
    let start = std::time::Instant::now();
    f();
    // pretty print the execution time with appropriate units to 3 decimal places
    println!(
        "Execution time: {}",
        util::pretty_print_time(start.elapsed().as_secs_f64())
    );
}

fn main() {
    let args: Vec<String> = env::args().collect();
    let days = &args[1..];
    days.iter().for_each(|day| {
        let day = day.parse::<u8>().unwrap();
        let input = util::get_day_input(day);
        match day {
            1 => time_execution(|| day1::run(&input)),
            2 => time_execution(|| day2::run(&input)),
            _ => println!("Day {} not implemented", day),
        };
    });
}
