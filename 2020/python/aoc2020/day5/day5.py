import os


rules = {"F": -1, "B": 1, "L": -1, "R": 1}
horizontal = list(range(8))
vertical = list(range(128))


def __read_input_file(filename):
    my_dir = os.path.dirname(os.path.realpath(__file__))
    my_input_file = os.path.join(my_dir, filename)
    with open(my_input_file) as my_input:
        lines = my_input.readlines()
    # data = {
    #     "vertical": [list(line.strip()[:7]) for line in lines],
    #     "horizontal": [list(line.strip()[7:]) for line in lines],
    # }
    data = [
        {"vertical": list(line.strip()[:7]), "horizontal": list(line.strip()[7:])}
        for line in lines
    ]
    return data


def get_half(code, whole):
    # print(code)
    # print(whole)
    side = rules[code]
    # print(side)
    length = int(len(whole) / 2)
    # print(length)
    if side > 0:
        return whole[length:]
    return whole[:length]


def binary_search(ticket):
    # print(ticket)
    values = vertical
    for v in ticket["vertical"]:
        values = get_half(v, values)
        # print(values)
    vresult = values[0]
    values = horizontal
    for h in ticket["horizontal"]:
        values = get_half(h, values)
        # print(values)
    hresult = values[0]
    return {"vertical": vresult, "horizontal": hresult}


def get_numbers(tickets):
    ticket_numbers = [binary_search(ticket) for ticket in tickets]
    return ticket_numbers


def evaluate_numbers(seat):
    return (seat["vertical"] * 8) + seat["horizontal"]


def find_ticket(ticket_numbers):
    results = []
    previous_ticket = ticket_numbers[0]
    for ticket in ticket_numbers:
        difference = ticket - previous_ticket
        if difference == 2:
            return ticket - 1
        previous_ticket = ticket
    return None


def main():
    # tickets = __read_input_file("test_input.txt")
    tickets = __read_input_file("input.txt")
    # [print(t) for t in tickets]
    seats = get_numbers(tickets)
    print(f"Seats: {seats}")
    ticket_numbers = [evaluate_numbers(seat) for seat in seats]
    print(f"Ticket numbers {ticket_numbers}")
    max_ticket = max(ticket_numbers)
    print(f"Max Ticket: {max_ticket}")
    sorted_tickets = sorted(ticket_numbers)
    print(f"Sorted tickets: {sorted_tickets}")
    result = find_ticket(sorted_tickets)
    print(f"My ticket: {result}")


if __name__ == "__main__":
    main()
