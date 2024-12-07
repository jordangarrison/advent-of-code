import day1
import gleeunit
import gleeunit/should

pub fn main() {
  gleeunit.main()
}

// gleeunit test functions end in `_test`
pub fn run_part1_test() {
  "3   4
4   3
2   5
1   3
3   9
3   3"
  |> day1.run_part1
  |> should.equal(11)
}
