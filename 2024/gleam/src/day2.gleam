import gleam/io
import gleam/list
import gleam/string
import gleam/int
import pull_day.{read_day}

pub fn main() {
  let input = read_day("2")
  io.println(
    "Day 2 Result: "
    <> input
    |> run
    |> int.to_string,
  )
}

pub fn run(input: String) -> Int {
  input
  |> string.split("\n")
  |> list.map(is_safe)
  |> list.filter(fn(x) { x })
  |> list.length
}

pub fn is_safe(input: String) -> Bool {
  let input_nums =
    input
    |> string.split(" ")
    |> list.map(to_int)
  case
    all_increasing(input_nums)
    && diff_by_no_more_than(input_nums, 3)
    && !two_nums_equivalent(input_nums),
    all_decreasing(input_nums)
    && diff_by_no_more_than(input_nums, 2)
    && !two_nums_equivalent(input_nums)
  {
    True, _ -> True
    _, True -> True
    False, False -> False
  }
}

pub fn diff_by_no_more_than(input: List(Int), n: Int) -> Bool {
  list.window(input, 2)
  |> list.all(fn(win) {
    case win {
      [a, b] -> int.absolute_value(a - b) <= n
      _ -> False
    }
  })
}

pub fn all_increasing(input: List(Int)) -> Bool {
  input == list.sort(input, by: int.compare)
}

pub fn all_decreasing(input: List(Int)) -> Bool {
  input == list.sort(input, by: fn(x, y) { int.compare(y, x) })
}

pub fn two_nums_equivalent(input: List(Int)) -> Bool {
  list.window(input, 2)
  |> list.any(fn(win) {
    case win {
      [a, b] -> a == b
      _ -> False
    }
  })
}

fn to_int(input: String) -> Int {
  case int.parse(input) {
    Ok(i) -> i
    Error(_) -> panic as "Could not parse input as integer"
  }
}
