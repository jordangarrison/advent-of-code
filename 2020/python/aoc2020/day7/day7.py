import os
from pathlib import Path


def __read_input_file(filename):
    my_dir = os.path.dirname(os.path.realpath(__file__))
    my_input_file = os.path.join(my_dir, filename)
    data = Path(my_input_file).read_text()
    return data


def clean_data(data):
    data = data.split("\n")
    data = [d.split(" contain ") for d in data]
    return data


def build_rules(data):
    rules = []
    for entry in data:
        item = {}
        item = {"type": " ".join(entry[0].split()[:-1]), "holds": []}
        for bag in entry[1].split(", "):
            bag_meta = bag.split()[:-1]
            if " ".join(bag_meta) == "no other":
                bag_entry = {"number": 0, "type": " ".join(bag_meta)}
            else:
                bag_entry = {"number": int(bag_meta[0]), "type": " ".join(bag_meta[1:])}
            item["holds"].append(bag_entry)
        rules.append(item)
    return rules


def holds_bag(bag, entry):
    for item in entry["holds"]:
        if bag == "no other":
            return False
        if bag == item["type"]:
            return True
    return False


def find_containing(bag, rules):
    items = []
    for rule in rules:
        if holds_bag(bag, rule):
            items.append(rule["type"])
            items.extend(find_containing(rule["type"], rules))
    return items


def count_holding(bag, rules, multiplier=1, count=0):
    print(f"Bag: {bag}, Multiplier: {multiplier}")
    for rule in rules:
        if rule["type"] == "no other":
            return multiplier
        if rule["type"] == bag:
            print(rule)
            for item in rule["holds"]:
                print(item)
                # count += multiplier * item["number"]
                count += multiplier * count_holding(item["type"], rules, item["number"])
                print(count)
    return count


def main():
    data = __read_input_file("test_input.txt")
    # data = __read_input_file("input.txt")
    data = clean_data(data)
    rules = build_rules(data)
    [print(rule) for rule in rules]
    # containing = [entry for entry in rules if holds_bag("shiny gold", entry)]
    containing = find_containing("shiny gold", rules)
    unique = list(set(containing))
    # [print(item) for item in unique]
    print(f"Number of bag types containing 'shiny gold': {len(unique)}")

    count = count_holding("shiny gold", rules)
    print(f"Number of bags inside shiny gold {count}")


if __name__ == "__main__":
    main()
