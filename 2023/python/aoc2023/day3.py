from pydoc import isdata
import symbol


def main(data: str):
    p1 = part1(data)
    p2 = part2(data)
    print(f"Day 3 Part 1: {p1}, Part 2: {p2}")


def get_symbol_coordinates(symbol_coordinate: tuple[int, int]) -> list[tuple[int, int]]:
    line, c = symbol_coordinate
    return [
        (line - 1 if line > 0 else line, c - 1 if c > 0 else c),
        (line, c - 1 if c > 0 else c),
        (line + 1 if line < 9 else line, c - 1 if c > 0 else c),
        (line - 1 if line > 0 else line, c),
        (line + 1 if line < 9 else line, c),
        (line - 1 if line > 0 else line, c + 1 if c < 9 else c),
        (line, c + 1 if c < 9 else c),
        (line + 1 if line < 9 else line, c + 1 if c < 9 else c),
    ]


def part1(data: str) -> int:
    sum = 0

    lines = data.splitlines()
    for i, line in enumerate(lines):
        print(i, line)
        num_array: list[list[tuple[int, str]]] = []
        symbol_coordinate_array: list[list]
        current_num: list[tuple[int, str]] = []
        for j, c in enumerate(line):
            if c == ".":
                if current_num != []:
                    num_array.append(current_num)
                    current_num = []
            elif c.isdigit():
                current_num.append((j, c))
            else:
                symbol_coordinates = get_symbol_coordinates((j, i))

    return sum


def part2(data: str) -> int:
    return 0
