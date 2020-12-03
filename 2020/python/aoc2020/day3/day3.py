from aoc2020.day1.day1 import multiply_list
import os


def __read_input_file(filename):
    my_dir = os.path.dirname(os.path.realpath(__file__))
    my_input_file = os.path.join(my_dir, filename)
    with open(my_input_file) as my_input:
        data = [list(line.strip()) for line in my_input.readlines()]
        # for line in my_input:
        #     line_arr = list(line.strip())
    return data


def path_finder(data, right, down):
    i = 0
    j = 0
    tree = 0
    for n in range(len(data) - 1):
        i = (i + right) % (len(data[n]))
        j = j + down
        if j % down != 0 or j > len(data):
            continue
        print(f"i: {i}, j: {j}")
        if data[j][i] == "#":
            tree = tree + 1
            data[j][i] = "X"
        else:
            data[j][i] = "O"
    return tree, data


def __print_path(data):
    for line in data:
        for c in line:
            print(c, end="")
        print()


def solve_day3(right, down):
    data = __read_input_file("input.txt")
    trees, data = path_finder(data, right, down)
    __print_path(data)
    del data
    return trees


def main():
    trees = []
    trees.append(solve_day3(1, 1))
    trees.append(solve_day3(3, 1))
    trees.append(solve_day3(5, 1))
    trees.append(solve_day3(7, 1))
    trees.append(solve_day3(1, 2))
    print(f"You would encounter {trees} trees")

    result = multiply_list(trees)
    print(f"Product of trees: {result}")


if __name__ == "__main__":
    main()
