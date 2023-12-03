from pydoc import isdata


def main(data: str):
    p1 = part1(data)
    p2 = part2(data)
    print(f"Day 3 Part 1: {p1}, Part 2: {p2}")


def get_symbol_coordinates(symbol_coordinate: tuple[int, int]) -> list[tuple[int, int]]:
    x, y = symbol_coordinate
    return [
        (x - 1 if x > 0 else x, y - 1 if y > 0 else y),
        (x, y - 1 if y > 0 else y),
        (x + 1 if x < 9 else x, y - 1 if y > 0 else y),
        (x - 1 if x > 0 else x, y),
        (x + 1 if x < 9 else x, y),
        (x - 1 if x > 0 else x, y + 1 if y < 9 else y),
        (x, y + 1 if y < 9 else y),
        (x + 1 if x < 9 else x, y + 1 if y < 9 else y),
    ]


def part1(data: str) -> int:
    sum = 0

    lines = data.splitlines()
    for i, line in enumerate(lines):
        print(i, line)
        num_array: list[int] = []
        current_num: list[tuple[int, str]] = []
        for j, c in enumerate(line):
            if c == ".":
                if current_num != []:
                    current_num.append((j, c))
            elif c.isdigit():
                current_num.append((j, c))

    return sum


def part2(data: str) -> int:
    return 0
