from argparse import ArgumentParser
import importlib as imp
import os
import util


def main():
    parser = ArgumentParser()
    parser.add_argument(
        "days",
        help="List of days to run",
        nargs="+",
    )
    args = parser.parse_args()

    for day in args.days:
        # make sure the module exists for the day module withing the package relative to this file
        current_dir = os.path.dirname(os.path.realpath(__file__))
        if int(day) < 0 and int(day) > 25:
            print(f"Day {day} does not exist on the advent calendar")
            continue
        if not os.path.exists(f"{current_dir}/day{day}.py"):
            print(f"Day {day} does not exist, creating...")
            util.create_day(day)
            continue

        # Load the day's module and run
        data = util.read_file(day)
        day = imp.import_module(f"day{day}")
        day.main(data)


if __name__ == "__main__":
    main()
