mod day1;
mod day2;
mod day3;
mod util;
use std::env;

fn main() {
    let args: Vec<String> = env::args().collect();
    let days = &args[1..];
    days.iter().for_each(|day| {
        let day = day.parse::<u8>().unwrap();
        let input = util::get_day_input(day);
        match day {
            1 => util::time_execution(|| day1::run(&input)),
            2 => util::time_execution(|| day2::run(&input)),
            3 => util::time_execution(|| day3::run(&input)),
            _ => println!("Day {} not implemented", day),
        };
    });
}
