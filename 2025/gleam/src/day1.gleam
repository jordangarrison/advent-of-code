import gleam/int
import gleam/io
import gleam/list
import gleam/string
import pull_day

pub fn main() {
  let result_part1 = pull_day.read_day("1") |> run_part1
  io.println("Day 1 Results\nPart1: " <> 0 |> int.to_string)
}

pub fn run_part1(input: String) -> Int {
  let parsed_input =
    input
    |> parse
    |> echo
  0
}

fn parse(input: String) {
  input
  |> string.split("\n")
  |> list.map(fn(line) { line })
}

fn update_position(change: Int, current: Int) -> Int {
  current + change % 100
}

fn count_zeros(positions: List(Int)) {
  positions
  |> list.filter(fn(x) { x == 0 })
  |> echo
  |> list.count
}
