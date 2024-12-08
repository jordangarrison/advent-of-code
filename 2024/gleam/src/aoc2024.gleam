import argv
import day1
import day2
import day4
import day5
import day8
import gleam/io
import pull_day

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
    "4" -> day4.main()
    "5" -> day5.main()
    "8" -> day8.main()
    _ -> io.println("Uknown day " <> day)
  }
}
