package day1

import (
	"fmt"
	"strconv"
	"strings"
)

type day1 struct {
	input string
}

// Return day1 struct
func NewDay1(input string) *day1 {
	return &day1{
		input: input,
	}
}

func Run(input string) {
	fmt.Println("Running Day 1 ")
	d1 := NewDay1(input)

	part1Result := d1.Part1()
	fmt.Printf("Part 1 Result: %d\n", part1Result)

	part2Result := d1.Part2()
	fmt.Printf("Part 2 Result: %d\n", part2Result)
}

func (d1 *day1) Part1() int {
	var previous int
	var sum int
	for i, numString := range strings.Split(d1.input, "\n") {
		num, _ := strconv.Atoi(numString)
		if i != 0 {
			if num > previous {
				sum++
			}
		}
		previous = num
	}
	return sum
}

func (d1 *day1) Part2() int {
	// create window of size 3
	// if window is increasing, add to sum
	var previous int
	var current int
	var sum int
	inputNumStrings := strings.Split(d1.input, "\n")
	for i := 2; i < len(inputNumStrings); i++ {
		num0, _ := strconv.Atoi(inputNumStrings[i-2])
		num1, _ := strconv.Atoi(inputNumStrings[i-1])
		num2, _ := strconv.Atoi(inputNumStrings[i])
		// fmt.Printf("%d %d %d\n", num0, num1, num2)
		current = num0 + num1 + num2
		// fmt.Printf("Current: %d\n", current)
		if i != 2 {
			if current > previous {
				sum++
			}
		}
		previous = current
	}
	return sum
}
