import argv
import day1
import gleam/io
import pull_day

pub fn main() -> Nil {
  case argv.load().arguments {
    ["pull", day] -> pull_day.pull_day(day, "./data")
    ["day", day] -> run_day(day)
    _ -> io.println("Usage: aoc2025 pull <day> | aoc2025 day <day>")
  }
}

fn run_day(day) {
  case day {
    "1" -> day1.main()
    _ -> io.println("Unknown day " <> day)
  }
}
