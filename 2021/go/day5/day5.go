package day5

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Day struct {
	input string
}

func NewDay(input string) *Day {
	return &Day{input}
}

func Run(input string) {
	d := NewDay(input)

	part1Result := d.Parts(false)
	fmt.Printf("Part 1: %v\n", part1Result)

	part2Result := d.Parts(true)
	fmt.Printf("Part 2: %v\n", part2Result)
}

func (d *Day) Parts(horizontals bool) int {
	lines := strings.Split(d.input, "\n")
	paths := [][][]int{}
	for _, line := range lines {
		change := strings.Split(line, " -> ")
		start := stringToInt(strings.Split(change[0], ","))
		end := stringToInt(strings.Split(change[1], ","))
		path, slope := getPath(start, end)
		// fmt.Printf("Path from %s -> %s: %v\n", change[0], change[1], path)
		if horizontals || slope == 0 || slope == math.Inf(1) || slope == math.Inf(-1) {
			paths = append(paths, path)
		}
	}
	travelledCoordinates := make(map[string]int)
	for _, path := range paths {
		for _, coordinate := range path {
			// check if key exists on travelledCoordinates
			coordinateString := strconv.Itoa(coordinate[0]) + "," + strconv.Itoa(coordinate[1])
			if _, ok := travelledCoordinates[coordinateString]; !ok {
				// fmt.Printf("Adding %v to travelledCoordinates\n", coordinateString)
				travelledCoordinates[coordinateString] = 0
			} else {
				// fmt.Printf("%v already exists in travelledCoordinates\n", coordinateString)
				travelledCoordinates[coordinateString]++
				// if travelledCoordinates[coordinateString] >= 1 {
				// 	fmt.Printf("%v is a duplicate\n", coordinateString)
				// }
			}
		}
	}
	// fmt.Printf("%+v\n", travelledCoordinates)
	sum := 0
	for _, v := range travelledCoordinates {
		if v >= 1 {
			sum++
		}
	}
	return sum
}

func stringToInt(s []string) []int {
	ints := []int{}
	for _, i := range s {
		int, _ := strconv.Atoi(i)
		ints = append(ints, int)
	}
	return ints
}

func min(a, b []int) ([]int, []int) {
	if a[0] < b[0] {
		return a, b
	}
	return b, a
}

func getPath(coord1 []int, coord2 []int) ([][]int, float64) {
	start, end := min(coord1, coord2)
	x1, y1 := start[0], start[1]
	x2, y2 := end[0], end[1]
	// fmt.Printf("x1: %v, y1: %v, x2: %v, y2: %v\n", x1, y1, x2, y2)
	path := [][]int{}
	// get the points between coordinates start and end
	slope := float64(y2-y1) / float64(x2-x1)
	// fmt.Printf("Slope: %v ", slope)
	for x1 <= x2 {
		path = append(path, []int{x1, y1})
		if x1 == x2 && y1 == y2 {
			break
		}
		if slope == 0 {
			x1++
		} else if (slope > 0 && slope < 1) || (slope > -1 && slope < 0) {
			y1++
			x1 += int(1 / slope)
		} else if (slope >= 1 && slope < math.Inf(1)) || (slope <= -1 && slope > math.Inf(-1)) {
			y1 += int(slope)
			x1++
		} else if slope == math.Inf(1) {
			y1++
		} else if slope == math.Inf(-1) {
			y1--
		} else {
			panic("Slope is not a number")
		}
	}
	return path, slope
}
