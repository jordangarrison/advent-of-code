import day1
import gleeunit
import gleeunit/should

const input = "L68
L30
R48
L5
R60
L55
L1
L99
R14
L82"

pub fn main() {
  gleeunit.main()
}

pub fn run_part_1_test() {
  input
  |> day1.run_part1
  |> should.equal(3)
}

pub fn run_part_2_test() {
  input
  |> day1.run_part2
  |> should.equal(6)
}
