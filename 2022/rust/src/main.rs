mod day1;
mod util;
use std::env;

fn main() {
    let args: Vec<String> = env::args().collect();
    let days = &args[1..];
    days.iter().for_each(|day| {
        let day = day.parse::<u8>().unwrap();
        let input = util::get_day_input(day);
        match day {
            1 => day1::run(&input),
            _ => println!("Day {} not implemented", day),
        };
    });
}
