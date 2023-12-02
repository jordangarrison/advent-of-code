import math


def main(data: str):
    p1 = part1(data)
    p2 = part2(data)
    print(f"Day 2 Part 1: {p1}, Part 2: {p2}")


def part1(data: str) -> int:
    configuration = {
        "red": 12,
        "green": 13,
        "blue": 14,
    }

    lines = [line for line in data.split("\n") if line != ""]

    possible_games: set[int] = set()
    for i, line in enumerate(lines):
        game = i + 1
        possible_games.add(game)
        line = line.replace(f"Game {game}: ", "")
        for draw in line.split("; "):
            for boxes in draw.split(", "):
                box_set = boxes.split(" ")
                if int(box_set[0]) > configuration[box_set[1]]:
                    possible_games.discard(game)
                    break

    return sum(set(possible_games))


def part2(data: str) -> int:
    lines = [line for line in data.split("\n") if line != ""]

    game_powers: list[int] = []
    for i, line in enumerate(lines):
        minimum_configuration = {
            "red": 0,
            "green": 0,
            "blue": 0,
        }
        game = i + 1
        line = line.replace(f"Game {game}: ", "")
        print(line)
        for j, draw in enumerate(line.split("; ")):
            for boxes in draw.split(", "):
                box_set = boxes.split(" ")
                if j == 0:
                    minimum_configuration[box_set[1]] = int(box_set[0])
                elif minimum_configuration[box_set[1]] < int(box_set[0]):
                    minimum_configuration[box_set[1]] = int(box_set[0])
        print(minimum_configuration)
        counts = [count for _, count in minimum_configuration.items()]
        power = math.prod(counts)
        print(power)
        game_powers.append(power)
    print(game_powers)
    return sum(game_powers)
