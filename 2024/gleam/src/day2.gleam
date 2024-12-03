import gleam/io
import gleam/list
import gleam/string
import gleam/int
import gleam_stats/stats

pub fn main() {
  io.println("Day 2")
}

pub fn run(input: String) -> Int {
  io.println("Input:\n" <> input)
  input
  |> string.split("\n")
  |> list.map(is_increasing_or_decreasing)
  |> list.filter(fn(x) { x })
  |> list.length
}

fn is_increasing_or_decreasing(input: String) -> Bool {
  let input_nums =
    input
    |> string.split(" ")
    |> list.map(to_int)
  let all_increasing =
    input_nums == list.sort(input_nums, by: fn(x, y) { int.compare(x, y) })
  let all_decreasing =
    input_nums == list.sort(input_nums, by: fn(x, y) { int.compare(y, x) })
  let distances =
    list.reduce(input_nums, fn(x, y) { int.absolute_value(x - y) })
  case all_increasing, all_decreasing, distances {
    False, False, _ -> False
    True, _, distances -> stats.amax(distances) <= 3
    _, True, distances -> stats.amax(distances) <= 3
    _, _, _ -> False
  }
}

fn to_int(input: String) -> Int {
  case int.parse(input) {
    Ok(i) -> i
    Error(_) -> 0
  }
}
