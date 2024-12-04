import argv
import gleam/io
import pull_day
import day1
import day2

pub fn main() {
  case argv.load().arguments {
    ["pull", day] -> pull_day.pull_day(day, "./data")
    ["day", day] -> run_day(day)
    _ -> io.println("Usage: aoc2024 pull <day> | aoc2024 day <day>")
  }
}

fn run_day(day) {
  case day {
    "1" -> day1.main()
    "2" -> day2.main()
    _ -> io.println("Uknown day " <> day)
  }
}
