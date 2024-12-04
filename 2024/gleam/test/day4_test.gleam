import gleeunit
import gleeunit/should
import day4

pub fn main() {
  gleeunit.main()
}

// main tests for day4
/// Test the example input from day 4
pub fn run_test() {
  "MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX"
  |> day4.run("XMAS")
  |> should.equal(18)
}
