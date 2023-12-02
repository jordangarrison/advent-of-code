import re


def main(data: str):
    p1 = part1(data)
    p2 = part2(data)
    print(f"Day 1 Part 1: {p1}, Part 2: {p2}")


def part1(data: str) -> int:
    print("Part 1")
    sum = 0

    lines = [line for line in data.split("\n") if line != ""]
    for line in lines:
        digits: list[str] = []
        for i in range(len(line)):
            if line[i].isdigit():
                digits.append(line[i])
        result = int(f"{digits[0]}{digits[-1]}")
        sum += result
    return sum


def part2(data: str):
    print("Part 2")
    sum = 0

    lines = [line for line in data.split("\n") if line != ""]
    num_words = ["one", "two", "three", "four", "five", "six", "seven", "eight", "nine"]

    for line in lines:
        # All numbers in the line with their positions
        num_map: list[tuple[int, int]] = []

        # Digit numbers
        for i in range(len(line)):
            if line[i].isdigit():
                num_map.append((i, int(line[i])))

        # Word numbers
        for i, num in enumerate(num_words):
            positions = [m.start() for m in re.finditer(num, line)]
            for pos in positions:
                num_map.append((pos, i + 1))
        first_num = (0, 0)
        last_num = (0, 0)
        for i, tup in enumerate(num_map):
            if i == 0:
                first_num = tup
                last_num = tup
                continue
            if tup[0] < first_num[0]:
                first_num = tup
            if tup[0] > last_num[0]:
                last_num = tup
        result = int(f"{first_num[1]}{last_num[1]}")
        print(num_map, result)
        sum += result
    return sum
