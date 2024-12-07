import gleam/int
import gleam/io
import gleam/list
import gleam/string
import pull_day

pub type Rule {
  Rule(before: String, after: String)
}

pub fn main() {
  let result =
    pull_day.read_day("5")
    |> io.debug
    |> run
  io.println("Day 5 Result: " <> result |> int.to_string)
}

pub fn run(input: String) -> Int {
  let parsed_input =
    input
    |> string.split("\n\n")
  let rules_input = case parsed_input |> list.first {
    Ok(rules) -> rules
    Error(_) -> ""
  }
  let data_input = case parsed_input |> list.last {
    Ok(data) -> data
    Error(_) -> ""
  }
  // rules_input
  // |> string.split("\n")
  // |> list.map(fn(rule) {
  //   case rule |> string.split("|") {
  //     _ -> todo
  //   }
  // })
  // io.debug(rules)
  // io.debug(data)
  0
}
