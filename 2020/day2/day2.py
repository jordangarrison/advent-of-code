import os
from functools import reduce


def __get_contents(fileobject):
    for line in fileobject:
        yield [l.strip() for l in line.split()]


def __create_record(line):
    result = {
        "minimum": int(line[0].split("-")[0]),
        "maximum": int(line[0].split("-")[1]),
        "character": line[1].replace(":", ""),
        "password": line[2],
    }
    return result


def __read_input_file(filename):
    my_dir = os.path.dirname(os.path.realpath(__file__))
    my_input_file = os.path.join(my_dir, filename)
    with open(my_input_file) as my_input:
        data = __get_contents(my_input)
        data = [d for d in data]
    data = map(__create_record, data)
    return list(data)


def __validate1(meta):
    num_val = meta["password"].count(meta["character"])
    meta["count"] = num_val
    if num_val >= meta["minimum"] and num_val <= meta["maximum"]:
        meta["valid"] = True
        return meta
    meta["valid"] = False
    return meta


def validate_passwords1(data1):
    return list(map(__validate1, data1))


def __validate2(meta):
    pos1 = meta["minimum"] - 1
    pos2 = meta["maximum"] - 1
    pass_arr = list(meta["password"])
    valid_count = 0
    if pass_arr[pos1] == meta["character"]:
        valid_count = valid_count + 1
    if pass_arr[pos2] == meta["character"]:
        valid_count = valid_count + 1
    meta["valid"] = valid_count
    print(pass_arr[pos1], pass_arr[pos2], meta["character"], valid_count, meta["valid"])
    return meta


def validate_passwords2(data2):
    return list(map(__validate2, data2))


def count_valid(final_data):
    valid = [d for d in final_data if d["valid"] == 1]
    print(valid)
    return len(valid)


def main():
    data = __read_input_file("input.txt")
    part1 = validate_passwords1(data)
    part1_num_valid = count_valid(part1)
    print(f"Number of valid passwords are {part1_num_valid}")

    data = __read_input_file("input.txt")
    part2 = validate_passwords2(data)
    part2_num_valid = count_valid(part2)
    print(f"Number of valid passwords are {part2_num_valid}")


if __name__ == "__main__":
    main()
