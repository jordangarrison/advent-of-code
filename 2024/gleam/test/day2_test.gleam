import gleeunit
import gleeunit/should
import day2

pub fn main() {
  gleeunit.main()
}

// gleeunit test functions end in `_test`
pub fn main_test() {
  "7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9"
  |> day2.run
  |> should.equal(2)
}
