from aoc2023 import day7


def test_part1():
    assert (
        day7.part1(
            """32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483"""
        )
        == 6440
    )


def test_part2():
    assert day7.part2("") == 0
