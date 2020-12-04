# from pprint import pprint
import re
import os
from pathlib import Path


def __read_input_file(filename):
    my_dir = os.path.dirname(os.path.realpath(__file__))
    my_input_file = os.path.join(my_dir, filename)
    data = Path(my_input_file).read_text()
    return data


def clean_input(data):
    result = []
    data = data.split("\n\n")
    for d in data:
        entry = {}
        line = re.split("\s+", d)
        for l in line:
            k = l.split(":")[0]
            v = l.split(":")[1]
            entry[k] = v
        result.append(entry)
    return result


def __validate_byr(pp):
    if "byr" not in pp.keys():
        return False
    byr = int(pp["byr"])
    if byr < 1920 or byr > 2002 or len(str(byr)) != 4:
        return False
    return True


def __validate_iyr(pp):
    if "iyr" not in pp.keys():
        return False
    iyr = int(pp["iyr"])
    if iyr < 2010 or iyr > 2020 or len(str(iyr)) != 4:
        return False
    return True


def __validate_eyr(pp):
    if "eyr" not in pp.keys():
        return False
    eyr = int(pp["eyr"])
    if eyr < 2020 or eyr > 2030 or len(str(eyr)) != 4:
        return False
    return True


def __validate_hgt(pp):
    if "hgt" not in pp.keys():
        return False
    hgt = pp["hgt"]
    if "in" in hgt:
        height_inches = int(hgt.split("in")[0])
        # print(height_inches)
        if height_inches < 59 or height_inches > 76:
            # print("invalid height")
            return False
    elif "cm" in hgt:
        height_cm = int(hgt.split("cm")[0])
        if height_cm < 150 or height_cm > 193:
            return False
    else:
        # print("unknown false return")
        return False
    return True


def __validate_hcl(pp):
    if "hcl" not in pp.keys():
        return False
    hcl = pp["hcl"]
    if hcl[0] != "#":
        return False
    if len(hcl) != 7:
        return False
    for c in list(hcl[1:]):
        if re.match("[0-9]|[a-f]", c) is None:
            return False
    return True


def __validate_ecl(pp):
    if "ecl" not in pp.keys():
        return False
    ecl = pp["ecl"]
    colors = ["amb", "blu", "brn", "gry", "grn", "hzl", "oth"]
    if ecl not in colors:
        return False
    return True


def __validate_pid(pp):
    if "pid" not in pp.keys():
        return False
    pid = pp["pid"]
    match = re.match(r"0+[0-9]+|[0-9]+", pid)
    if match is None or len(match.group()) != 9:
        return False
    return True


def validate_passport(data):
    for i in range(len(data)):
        data[i]["valid"] = True
        if len(data[i].items()) <= 8:
            data[i]["valid"] = False
            if len(data[i].items()) == 8 and "cid" not in data[i].keys():
                data[i]["valid"] = True
        results = []
        results.append(__validate_byr(data[i]))
        results.append(__validate_iyr(data[i]))
        results.append(__validate_eyr(data[i]))
        results.append(__validate_hgt(data[i]))
        results.append(__validate_hcl(data[i]))
        results.append(__validate_ecl(data[i]))
        results.append(__validate_pid(data[i]))
        # print(results)
        # print()
        if False in results:
            data[i]["valid"] = False
    return data


def count_valid(data):
    valid = [d for d in data if d["valid"]]
    return len(valid)


def main():
    # data = __read_input_file("test_input1.txt")
    # data = __read_input_file("test_input2.txt")
    data = __read_input_file("input.txt")
    data = clean_input(data)
    data = validate_passport(data)
    # pprint([d for d in data], compact=True, width=120)
    num_valid = count_valid(data)
    print(f"Number of valid passwords are {num_valid}")


if __name__ == "__main__":
    main()
