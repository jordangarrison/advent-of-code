#!/bin/sh

# get inputs for day
DAY=${1:?"Please provide a day"}

# validate aoc session token exists
if [ -z "$AOC_SESSION_TOKEN" ]; then
    echo "Please set AOC_SESSION_TOKEN environment variable"
    exit 1
fi

curl \
  -b session=${AOC_SESSION_TOKEN} \
  -H "User-Agent: Jordan Garrison/aoc2022 Pulling my data" \
  https://adventofcode.com/2022/day/${DAY}/input \
  -o data/day${DAY}.txt