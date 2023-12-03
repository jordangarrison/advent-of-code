from aoc2023 import day3


def test_part1():
    assert (
        day3.part1(
            """467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598.."""
        )
        == 4361
    )


def test_part2():
    assert day3.part2("") == 0
