package day6

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/thoas/go-funk"
)

type Day struct {
	input string
}

func NewDay(input string) *Day {
	return &Day{input}
}

func Run(input string) {
	d := NewDay(input)

	part1Result := d.Parts(80)
	fmt.Printf("Part 1: %v\n", part1Result)

	part2Result := d.Parts(256)
	fmt.Printf("Part 2: %v\n", part2Result)
}

func (d *Day) ParseInput() []int {
	return funk.Map(strings.Split(d.input, ","), func(s string) int {
		i, _ := strconv.Atoi(s)
		return i
	}).([]int)
}

func (d *Day) Parts(numDays int) int {
	input := d.ParseInput()
	state := map[int]int{
		0: 0,
		1: 0,
		2: 0,
		3: 0,
		4: 0,
		5: 0,
		6: 0,
		7: 0,
		8: 0,
	}
	for _, i := range input {
		state[i] += 1
	}
	// fmt.Printf("%+v\n", state)
	for day := 0; day < numDays; day++ {
		// fmt.Printf("Day %v, %v\n", day+1, state)
		pastState := map[int]int{}
		for i, v := range state {
			pastState[i] = v
		}
		state[0] = pastState[1]
		state[1] = pastState[2]
		state[2] = pastState[3]
		state[3] = pastState[4]
		state[4] = pastState[5]
		state[5] = pastState[6]
		state[6] = pastState[7] + pastState[0]
		state[7] = pastState[8]
		state[8] = pastState[0]
	}
	sum := 0
	for _, i := range state {
		sum += i
	}
	return sum
}
