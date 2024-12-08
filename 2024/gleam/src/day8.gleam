import gleam/int
import gleam/io
import pull_day

pub type Coordinate {
  Coordinate(x: Int, y: Int)
}

pub type AntiNode {
  AntiNode(coord: Coordinate, number: Int)
}

pub type Direction =
  Coordinate

pub type Grid =
  List(List(Int))

pub fn main() {
  let input = pull_day.read_day("8")
  io.println(
    "Day 2 Result:\nPart1: "
    <> input
    |> run_part1
    |> int.to_string
    <> "\nPart2: "
    <> input
    |> run_part2
    |> int.to_string,
  )
}

pub fn run_part1(input: String) -> Int {
  input |> io.debug
  0
}

pub fn run_part2(_input: String) -> Int {
  //   input |> io.debug
  0
}
