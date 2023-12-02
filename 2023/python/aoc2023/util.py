import os


# read files from data folder for given day of advent of code
def read_file(day):
    with open(f"data/day{day}.txt", "r") as f:
        return f.read()


# Function which creates the code boilerplate for a given day
def create_day(day):
    curdir = os.path.dirname(os.path.realpath(__file__))
    with open(f"{curdir}/day{day}.py", "w") as f:
        f.write(
            """def main(data: str):
    p1 = part1(data)
    p2 = part2(data)
    print(f"Day {day} Part 1: {p1}, Part 2: {p2}")
    
def part1(data: str) -> int:
    return 0
    
def part2(data: str) -> int:
    return 0
""".format(
                day=day,
                p1="{p1}",
                p2="{p2}",
            )
        )

    with open(f"{curdir}/../tests/test_day{day}.py", "w") as f:
        f.write(
            """from aoc2023 import day{day}

def test_part1():
    assert day{day}.part1("") == 0
    
def test_part2():
    assert day{day}.part2("") == 0
""".format(
                day=day
            )
        )
