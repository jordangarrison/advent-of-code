import os
import re
from pathlib import Path


def __read_input_file(filename):
    my_dir = os.path.dirname(os.path.realpath(__file__))
    my_input_file = os.path.join(my_dir, filename)
    data = Path(my_input_file).read_text()
    return data


def parse_data(data):
    group_responses = data.split("\n\n")
    trim_whitespace = [i.split("\n") for i in group_responses]
    clean_data = trim_whitespace
    return clean_data


def __build_dict(response):
    data = {}
    # print(f"Required number {num_required}")
    for individual in response:
        for r in individual:
            # print(r)
            if r in data.keys():
                data[r] += 1
            else:
                data[r] = 1
    return data


def get_answers(responses):
    answers = []
    for response in responses:
        data = __build_dict(response)
        answers.append(len(data))
    return answers


def get_all_answers(responses):
    answers = []
    for response in responses:
        num_required = len(response)
        data = __build_dict(response)
        for k, v in data.items():
            if num_required == v:
                answers.append(1)
            else:
                answers.append(0)
    return answers


def evaluate_answers(answers):
    return sum(answers)


def main():
    # data = __read_input_file("test_input.txt")
    data = __read_input_file("input.txt")
    clean_data = parse_data(data)
    # print(clean_data)
    answers = get_answers(clean_data)
    result = evaluate_answers(answers)
    print(f"Passenger answers: {result}")

    answers = get_all_answers(clean_data)
    result = evaluate_answers(answers)
    print(f"Refined passenger answers: {result}")


if __name__ == "__main__":
    main()
