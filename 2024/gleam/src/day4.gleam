import gleam/io
import gleam/bool
import gleam/string
import gleam/list
import gleam/int
import gleam/dict.{type Dict}
import pull_day.{read_day}

pub const directions = [
  Coordinate(0, 1),
  Coordinate(1, 0),
  Coordinate(0, -1),
  Coordinate(-1, 0),
  Coordinate(1, 1),
  Coordinate(1, -1),
  Coordinate(-1, 1),
  Coordinate(-1, -1),
]

pub const origin = Coordinate(0, 0)

pub const word = "XMAS"

pub const max = 140

pub type Grid(data) =
  Dict(Coordinate, data)

pub type Coordinate {
  Coordinate(x: Int, y: Int)
}

pub type Word =
  List(Coordinate)

pub fn main() {
  let result =
    read_day("4")
    |> run(word)
    |> int.to_string
  io.println("Day 4 Result: " <> result)
}

pub fn run(input: String, word: String) -> Int {
  let grid =
    get_grid(input, fn(x) { x })
    |> io.debug
  count_words(grid, word, origin, 0)
  0
}

pub fn get_grid(input: String, parser: fn(String) -> data) -> Grid(data) {
  {
    use row, r <- list.index_map(string.split("\n", input))
    use col, c <- list.index_map(string.to_graphemes(row))
    #(Coordinate(r, c), parser(col))
  }
  |> list.flatten
  |> dict.from_list
}

pub fn count_words(
  grid: Grid(String),
  word: String,
  coord: Coordinate,
  acc: Int,
) -> Int {
  use <- bool.guard(coord.x == max, acc)
  use <- bool.lazy_guard(coord.y == max, fn() {
    count_words(grid, word, next_row(coord), acc)
  })

  let found =
    directions
    |> list.filter_map(scan_for_word(grid, word, coord, _))
    |> list.length
  count_words(grid, word, next_col(coord), acc + found)
}

pub fn scan_for_word(
  grid: Grid(String),
  word: String,
  coord: Coordinate,
  direction: Coordinate,
) {
  case string.pop_grapheme(word), dict.get(grid, coord) {
    Ok(#(first, "")), Ok(val) if first == val -> Ok(Nil)
    Ok(#(first, rest)), Ok(val) if first == val ->
      scan_for_word(grid, rest, move_direction(coord, direction), direction)
    _, _ -> Error(Nil)
  }
}

pub fn move_direction(coord: Coordinate, direction: Coordinate) -> Coordinate {
  Coordinate(coord.x + direction.x, coord.y + direction.y)
}

pub fn next_row(coord: Coordinate) -> Coordinate {
  Coordinate(coord.x + 1, 0)
}

pub fn next_col(coord: Coordinate) -> Coordinate {
  Coordinate(0, coord.y + 1)
}
