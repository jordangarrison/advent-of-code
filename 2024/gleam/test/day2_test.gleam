import day2
import gleeunit
import gleeunit/should

pub fn main() {
  gleeunit.main()
}

// gleeunit test functions end in `_test`
pub fn run_part1_test() {
  "7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9"
  |> day2.run_part1
  |> should.equal(2)
}

// gleeunit test functions end in `_test`
pub fn run2_part1_test() {
  "7 10 3 2 1
1 5 7 8 9
9 1
9 9 9 9 9
1 2 3 4 5
5 4 3 2 1"
  |> day2.run_part1
  |> should.equal(2)
}

pub fn run3_part1_test() {
  "79 78 76 74 72
4 7 9 12 15 17 18
70 67 65 63 62 60 58 56
42 44 46 47 48 51 53 54
64 63 62 59 58"
  |> day2.run_part1
  |> should.equal(5)
}

pub fn run4_part1_test() {
  "64 63 62 59 58"
  |> day2.run_part1
  |> should.equal(1)
}

pub fn run5_part1_test() {
  "64 63 62 60 58"
  |> day2.run_part1
  |> should.equal(1)
}

pub fn run_part1_all_the_same_test() {
  "1 1 1 1 1"
  |> day2.run_part1
  |> should.equal(0)
}

pub fn run_part1_all_decreasing_test() {
  "5 4 3 2 1"
  |> day2.run_part1
  |> should.equal(1)
}

pub fn run_part1_all_increasing_test() {
  "1 2 3 4 5"
  |> day2.run_part1
  |> should.equal(1)
}

pub fn run_part1_all_increasing_by_3_test() {
  "1 2 5"
  |> day2.run_part1
  |> should.equal(1)
}

pub fn run_part1_all_decreasing_by_2_test() {
  "5 4 2"
  |> day2.run_part1
  |> should.equal(1)
}

pub fn run_part1_all_increasing_with_4_test() {
  "1 2 3 4 8"
  |> day2.run_part1
  |> should.equal(0)
}

pub fn run_part1_all_decreasing_with_3_test() {
  "5 4 1"
  |> day2.run_part1
  |> should.equal(1)
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

pub fn decreasing_by_more_than_3_test() {
  [8, 5, 1]
  |> day2.diff_by_no_more_than(3)
  |> should.equal(False)
}

pub fn run_part2_test() {
  "7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9"
  |> day2.run_part2
  |> should.equal(4)
}
