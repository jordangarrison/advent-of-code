package day2

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Day2 struct {
	input string
}

var directions = map[string]int{
	"forward": 1,
	"down":    1,
	"up":      -1,
}

func NewDay2(input string) *Day2 {
	return &Day2{input}
}

func Run(input string) {
	d := NewDay2(input)
	part1Result := d.Part1(input)
	fmt.Printf("Part 1: %d\n", part1Result)
	part2Result := d.Part2(input)
	fmt.Printf("Part 2: %d\n", part2Result)
}

func (d *Day2) Part1(input string) int {
	position := [2]int{0, 0}

	lines := strings.Split(input, "\n")
	re := regexp.MustCompile(`\s+`)
	for _, line := range lines {
		words := re.Split(line, -1)
		// fmt.Printf("%v\n", words)
		direction := words[0]
		distance, _ := strconv.Atoi(words[1])
		if direction == "up" || direction == "down" {
			position[1] += directions[direction] * distance
		} else if direction == "forward" {
			position[0] += directions[direction] * distance
		} else {
			fmt.Printf("Invalid direction: %s\n", direction)
		}
		// fmt.Printf("Position: %v\n", position)
	}
	return position[0] * position[1]
}

func (d *Day2) Part2(input string) int {
	// position is [x, y, aim]
	position := [3]int{0, 0, 0}

	lines := strings.Split(input, "\n")
	re := regexp.MustCompile(`\s+`)
	for _, line := range lines {
		words := re.Split(line, -1)
		direction := words[0]
		distance, _ := strconv.Atoi(words[1])
		if direction == "up" || direction == "down" {
			position[2] += directions[direction] * distance
		} else if direction == "forward" {
			forward := directions[direction] * distance
			depth := position[2] * forward
			position[0] += forward
			position[1] += depth
		} else {
			fmt.Printf("Invalid direction: %s\n", direction)
		}
	}
	return position[0] * position[1]
}
