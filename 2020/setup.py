import os
from setuptools import setup, find_packages

setup(
    name="aoc2020",
    version="2020",
    packages=find_packages(),
    author="Jordan Garrison",
    author_email="jordan.andrew.garrison@gmail.com",
    description="Advent of Code 2020",
    url="https://github.com/jordangarrison/advent-of-code",
    package_data={
        # Include any packages inside of *.txt and *.rst files
        "": ["*.txt", "*.rst"]
    },
    long_description=open("README.md").read(),
)
