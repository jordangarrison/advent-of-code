import os
from setuptools import setup, find_packages

setup(
    name="aoc2023",
    version="2023.0.0",
    packages=find_packages(),
    author="Jordan Garrison",
    author_email="jordan@jordangarrison.dev",
    description="Advent of Code 2023",
    url="https://github.com/jordangarrison/advent-of-code",
    package_data={
        # If any package contains *.txt or *.rst files, include them:
        "": ["*.txt", "*.rst"],
    },
    entry_points={
        "console_scripts": [
            "aoc2023=aoc2023.main:main",
        ],
    },
)
