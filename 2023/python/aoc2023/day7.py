VALUE_MAP = {
    "A": 14,
    "K": 13,
    "Q": 12,
    "J": 11,
    "T": 10,
    "9": 9,
    "8": 8,
    "7": 7,
    "6": 6,
    "5": 5,
    "4": 4,
    "3": 3,
    "2": 2,
}


def is_five_of_a_kind(hand: dict[str, int]) -> int:
    for k, v in hand.items():
        if v == 5:
            return VALUE_MAP[k] * 10000000
    return 0


def is_four_of_a_kind(hand: dict[str, int]) -> int:
    for k, v in hand.items():
        if v == 4:
            return VALUE_MAP[k] * 1000000
    return 0


def is_full_house(hand: dict[str, int]) -> int:
    value = 0
    for k, v in hand.items():
        if v == 3:
            value += 100000 * VALUE_MAP[k]
        elif v == 2:
            value += 10000 * VALUE_MAP[k]
    return value


def is_three_of_a_kind(hand: dict[str, int]) -> int:
    for k, v in hand.items():
        if v == 3:
            return VALUE_MAP[k] * 10000
    return 0


def is_two_pairs(hand: dict[str, int]) -> int:
    value = 0
    for k, v in hand.items():
        if v == 2:
            value += 1000 * VALUE_MAP[k]
    return value


def is_one_pair(hand: dict[str, int]) -> int:
    for k, v in hand.items():
        if v == 2:
            return 100 * VALUE_MAP[k]
    return 0


def is_high_card(hand: dict[str, int]) -> int:
    value = 0
    for k, _ in hand.items():
        value += VALUE_MAP[k]
    return value


def get_hand_strength(hand: dict[str, int]) -> dict[str, dict[str, int] | str | int]:
    value: int = 0
    if is_five_of_a_kind(hand) > 0:
        value = is_five_of_a_kind(hand)
    elif is_four_of_a_kind(hand) > 0:
        value = is_four_of_a_kind(hand)
    elif is_full_house(hand) > 0:
        value = is_full_house(hand)
    elif is_three_of_a_kind(hand) > 0:
        value = is_three_of_a_kind(hand)
    elif is_two_pairs(hand) > 0:
        value = is_two_pairs(hand)
    elif is_one_pair(hand) > 0:
        value = is_one_pair(hand)
    else:
        value = is_high_card(hand)
    return {"hand": hand, "value": value}


def part1(data: str) -> int:
    hand_strength_array = []
    for line in data.splitlines():
        hand, bid = line.split(" ")
        hand_map = {}
        for card in hand:
            if card in hand_map:
                hand_map[card] += 1
            else:
                hand_map[card] = 1
        print(hand_map)
        print(bid)
        hand_strength_array.append({"hand": get_hand_strength(hand_map), "bid": bid})

    sorted_hand_strength_array = sorted(
        hand_strength_array,
        key=lambda x: (x["hand"]["value"]),
    )

    sum = 0
    for i, hand in enumerate(sorted_hand_strength_array):
        bid_calc = int(hand["bid"]) * (i + 1)
        print(
            f"Rank: {i+1}, Hand: {hand['hand']}, Bid: {hand['bid']}, Bid Calc: {bid_calc}"
        )
        sum += bid_calc

    return sum


def part2(data: str) -> int:
    return 0


def main(data: str):
    p1 = part1(data)
    p2 = part2(data)
    print(f"Day 7 Part 1: {p1}, Part 2: {p2}")
