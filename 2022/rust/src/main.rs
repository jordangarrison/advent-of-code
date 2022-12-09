mod day1;
mod day2;
mod day3;
mod day4;
mod day5;
mod day6;
mod day7;
mod day8;
mod day9;
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
            4 => util::time_execution(|| day4::run(&input)),
            5 => util::time_execution(|| day5::run(&input)),
            6 => util::time_execution(|| day6::run(&input)),
            7 => util::time_execution(|| day7::run(&input)),
            8 => util::time_execution(|| day8::run(&input)),
            9 => util:: time_execution(|| day9::run(&input)),
            _ => println!("Day {} not implemented", day),
        };
    });
}
