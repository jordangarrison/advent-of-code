import day2
import gleam/int
import gleam/io
import gleam/list
import gleam/regexp
import gleam/string
import pull_day

pub fn main() {
  let result_part1 =
    pull_day.read_day("1")
    |> run_part1

  let result_part2 = 0

  io.println(
    "Day 1 Results\nPart1: "
    <> result_part1 |> int.to_string
    <> "\nPart2: "
    <> result_part2 |> int.to_string,
  )
}

pub fn run_part1(input: String) -> Int {
  input
  |> parse
  |> rotate_list([], [])
  |> sort_lists
  |> io.debug
  0
}

/// convert string to 2 lists of ints
// sameple input 
// "3   4
//  4   3
//  2   5
//  1   3
//  3   9
//  3   3"
// output
// #([3, 4, 2, 1, 3, 3], [4, 3, 5, 3, 9, 3])
pub fn parse(input: String) {
  let assert Ok(re) = regexp.from_string("\\s+")
  input
  |> string.split("\n")
  |> list.map(fn(x) {
    x
    |> regexp.split(with: re, content: _)
    |> list.map(day2.to_int)
  })
}

fn rotate_list(
  input: List(List(Int)),
  list1: List(Int),
  list2: List(Int),
) -> #(List(Int), List(Int)) {
  case input {
    [pair, ..rest] ->
      case pair {
        [a, b] ->
          rotate_list(rest, list.append(list1, [a]), list.append(list2, [b]))
        _ -> #(list1, list2)
      }
    [] -> #(list1, list2)
  }
}

fn sort_lists(lists: #(List(Int), List(Int))) -> #(List(Int), List(Int)) {
  case lists {
    #(list1, list2) -> {
      let sorted_list1 = list.sort(list1, by: int.compare)
      let sorted_list2 = list.sort(list2, by: int.compare)
      #(sorted_list1, sorted_list2)
    }
  }
}
