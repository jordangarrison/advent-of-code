import os
import random
from functools import reduce


def __read_input_file(filename):
    my_dir = os.path.dirname(os.path.realpath(__file__))
    my_input_file = os.path.join(my_dir, filename)
    with open(my_input_file) as my_input:
        data = my_input.readlines()
    data = list(map(lambda d: int(d.strip()), data))
    return data


def get_2_2020_numbers(number_list):
    for d in number_list:
        for dd in number_list:
            if d + dd == 2020:
                print(f"First number: {d}\nSecond number: {dd}")
                return (d, dd)


def get_3_2020_numbers(number_list):
    for d in number_list:
        for dd in number_list:
            for ddd in number_list:
                if d + dd + ddd == 2020:
                    print(
                        f"First number: {d}\nSecond number: {dd}\nThird number: {ddd}"
                    )
                    return (d, dd, ddd)


def get_2020_numbers(number_list, length):
    target = 2020
    mysum = 0
    while mysum != 2020:
        result = random.sample(number_list, length)
        mysum = sum(result)
    return result


def multiply_list(num_list):
    return reduce(lambda x, y: x * y, num_list, 1)


def main():
    data = __read_input_file("input.txt")
    result = get_2_2020_numbers(data)
    answer = multiply_list(result)
    print(f"2 Numbers Answer: {answer}")
    result = get_3_2020_numbers(data)
    answer = multiply_list(result)
    print(f"3 Numbers Answer: {answer}")

    # Generic but really slow
    result = get_2020_numbers(data, 2)
    answer = multiply_list(result)
    print(f"2 Numbers Answer: {answer}")
    result = get_2020_numbers(data, 3)
    answer = multiply_list(result)
    print(f"3 Numbers Answer: {answer}")


if __name__ == "__main__":
    main()
