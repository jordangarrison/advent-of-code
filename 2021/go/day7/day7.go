package day7

import (
	"fmt"
	"math"
	"sort"
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

	part1Result := d.Part1()
	fmt.Printf("Part 1: %v\n", part1Result)

	part2Result := d.Part2()
	fmt.Printf("Part 2: %v\n", part2Result)
}

func (d *Day) ParseInput() []int {
	return funk.Map(strings.Split(d.input, ","), func(s string) int {
		i, _ := strconv.Atoi(s)
		return i
	}).([]int)
}

func (d *Day) Part1() int {
	input := d.ParseInput()
	sort.Ints(input)
	// fmt.Printf("Sorted: %v\n", input)
	distance, _ := leastDistance(input)
	// fmt.Printf("Least Distance number: %v, Total distance: %v", position, distance)
	return distance
}

func leastDistance(numbers []int) (int, int) {
	min := float64(-1)
	position := numbers[0]
	max := funk.MaxInt(numbers)
	for i := 0; i <= max; i++ {
		sum := 0.0
		for j := range numbers {
			sum += math.Abs(float64(numbers[j]) - float64(i))
		}
		if int(min) == -1 || min > sum {
			min = sum
			position = i
			// fmt.Printf("Updated position: %v\nNew distance: %v\n\n", position, min)
		}
	}
	return int(min), position
}

func (d *Day) Part2() int {
	input := d.ParseInput()
	sort.Ints(input)
	// fmt.Printf("Sorted: %v\n", input)
	cost, _ := leastDistanceNthTriangular(input)
	// fmt.Printf("Least Distance number: %v, Total distance: %v", position, distance)
	return cost
}

func leastDistanceNthTriangular(numbers []int) (int, int) {
	min := -1
	position := numbers[0]
	max := funk.MaxInt(numbers)
	for i := 0; i <= max; i++ {
		sum := 0
		for j := range numbers {
			diff := int(math.Abs(float64(i) - float64(numbers[j])))
			sum += nthTriangular(diff)
		}
		// fmt.Printf("Sum of %v: %v\n", i, sum)
		if min == -1 || min > sum {
			min = sum
			position = i
			// fmt.Printf("Updated position: %v\nNew distance: %v\n\n", position, min)
		}
	}
	return min, position
}

func nthTriangular(input int) int {
	cost := (input * (input + 1)) / 2
	return cost
}
