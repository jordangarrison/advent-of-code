import gleam/int
import gleam/io
import gleam/list
import gleam/result
import gleam/string
import pull_day

pub fn main() {
  let result_part1 = pull_day.read_day("1") |> run_part1
  io.println("Day 1 Results\nPart1: " <> result_part1 |> int.to_string)
}

pub fn run_part1(input: String) -> Int {
  let starting_position = 50
  let instructions =
    input
    |> parse

  let #(_, positions) =
    list.map_fold(instructions, starting_position, fn(position, instruction) {
      let new_position = update_position(instruction, position)
      #(new_position, new_position)
    })
  positions
  |> echo
  |> list.filter(fn(pos) { pos == 0 })
  |> list.length
}

fn parse(input: String) {
  input
  |> string.split("\n")
  |> list.map(parse_instruction)
}

fn parse_instruction(instruction: String) {
  case instruction {
    "L" <> num_str ->
      int.parse(num_str)
      |> result.map(fn(number) { -number })
      |> result.unwrap(0)
    "R" <> num_str -> int.parse(num_str) |> result.unwrap(0)
    _ -> panic(Nil)
  }
}

fn update_position(change: Int, current: Int) -> Int {
  let valid_change = change % 100
  let position_raw = valid_change + current
  case position_raw {
    position if position < 0 -> position + 100
    position if position > 99 -> position - 100
    _ -> position_raw
  }
}

fn count_zeros(positions: List(Int)) {
  positions
  |> list.filter(fn(x) { x == 0 })
  |> echo
}
