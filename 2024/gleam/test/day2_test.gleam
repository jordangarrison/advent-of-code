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

// gleeunit test functions end in `_test`
pub fn main2_test() {
  "7 10 3 2 1
1 5 7 8 9
9 1
9 9 9 9 9
1 2 3 4 5
5 4 3 2 1"
  |> day2.run
  |> should.equal(2)
}

pub fn all_increasing_test() {
  [1, 2, 3, 4, 5]
  |> day2.all_increasing
  |> should.equal(True)
}

pub fn all_not_increasing_test() {
  [1, 2, 3, 2, 1]
  |> day2.all_increasing
  |> should.equal(False)
}

pub fn all_decreasing_test() {
  [5, 4, 3, 2, 1]
  |> day2.all_decreasing
  |> should.equal(True)
}

pub fn all_not_decreasing_test() {
  [5, 4, 3, 4, 5]
  |> day2.all_decreasing
  |> should.equal(False)
}

pub fn two_nums_equivalent_test() {
  [1, 2, 3, 3, 5]
  |> day2.two_nums_equivalent
  |> should.equal(True)
}

pub fn two_nums_not_equivalent_test() {
  [1, 2, 3, 4, 5]
  |> day2.two_nums_equivalent
  |> should.equal(False)
}

pub fn is_safe_increasing_test() {
  "1 2 3 4 5"
  |> day2.is_safe
  |> should.equal(True)
}

pub fn is_safe_increasing_by_3_test() {
  "1 2 5"
  |> day2.is_safe
  |> should.equal(True)
}

pub fn is_safe_decreasing_test() {
  "5 4 3 2 1"
  |> day2.is_safe
  |> should.equal(True)
}

pub fn is_safe_decreasing_by_2_test() {
  "5 4 2"
  |> day2.is_safe
  |> should.equal(True)
}

pub fn is_not_safe_decreasing_by_1_test() {
  "5 4 3 2 2"
  |> day2.is_safe
  |> should.equal(False)
}

pub fn increasing_by_3_test() {
  [1, 2, 5]
  |> day2.diff_by_no_more_than(3)
  |> should.equal(True)
}

pub fn increasing_by_more_than_3_test() {
  [1, 2, 6]
  |> day2.diff_by_no_more_than(3)
  |> should.equal(False)
}

pub fn decreasing_by_2_test() {
  [5, 4, 2]
  |> day2.diff_by_no_more_than(2)
  |> should.equal(True)
}

pub fn decreasing_by_more_than_2_test() {
  [5, 4, 1]
  |> day2.diff_by_no_more_than(2)
  |> should.equal(False)
}
