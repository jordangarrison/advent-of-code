package day9

import (
	"fmt"
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

func (d *Day) readInput() [][]int {
	var digits [][]int
	for _, line := range strings.Split(d.input, "\n") {
		var lineArray []int
		for _, digitString := range strings.Split(line, "") {
			digit, _ := strconv.Atoi(digitString)
			lineArray = append(lineArray, int(digit))
		}
		digits = append(digits, lineArray)
	}
	return digits
}
func (d *Day) Part1() int {
	data := d.readInput()
	lowest := getLowest(data)
	// fmt.Printf("Found lowest points: %v\n", lowest)
	return funk.SumInt(funk.Map(lowest, func(coordinate [2]int) int {
		return data[coordinate[0]][coordinate[1]] + 1
	}).([]int))
}

func getLowest(data [][]int) [][2]int {
	numLines := len(data)
	numDigits := len(data[0])
	var lowest [][2]int
	for i := 0; i < numLines; i++ {
		for j := 0; j < numDigits; j++ {
			if (i == 0 || i == numLines-1) && (j == 0 || j == numDigits-1) {
				if i == 0 && j == 0 {
					if data[i][j] < data[i+1][j] && data[i][j] < data[i][j+1] {
						lowest = append(lowest, [2]int{i, j})
					}
				} else if i == 0 && j == numDigits-1 {
					if data[i][j] < data[i+1][j] && data[i][j] < data[i][j-1] {
						lowest = append(lowest, [2]int{i, j})
					}
				} else if i == numLines-1 && j == 0 {
					if data[i][j] < data[i-1][j] && data[i][j] < data[i][j+1] {
						lowest = append(lowest, [2]int{i, j})
					}
				} else {
					if data[i][j] < data[i-1][j] && data[i][j] < data[i][j-1] {
						lowest = append(lowest, [2]int{i, j})
					}
				}
			} else if i == 0 || i == numLines-1 || j == 0 || j == numDigits-1 {
				if i == 0 {
					if data[i][j] < data[i+1][j] && data[i][j] < data[i][j+1] && data[i][j] < data[i][j-1] {
						lowest = append(lowest, [2]int{i, j})
					}
				} else if i == numLines-1 {
					if data[i][j] < data[i-1][j] && data[i][j] < data[i][j+1] && data[i][j] < data[i][j-1] {
						lowest = append(lowest, [2]int{i, j})
					}
				} else if j == 0 {
					if data[i][j] < data[i+1][j] && data[i][j] < data[i-1][j] && data[i][j] < data[i][j+1] {
						lowest = append(lowest, [2]int{i, j})
					}
				} else {
					if data[i][j] < data[i+1][j] && data[i][j] < data[i-1][j] && data[i][j] < data[i][j-1] {
						lowest = append(lowest, [2]int{i, j})
					}
				}
			} else {
				if data[i][j] < data[i+1][j] && data[i][j] < data[i-1][j] && data[i][j] < data[i][j+1] && data[i][j] < data[i][j-1] {
					lowest = append(lowest, [2]int{i, j})
				}
			}
		}
	}
	return lowest
}

func (d *Day) Part2() int {
	data := d.readInput()
	lowPoints := getLowest(data)
	numLines := len(data)
	numDigits := len(data[0])
	var visited [][]bool
	for i := 0; i < numLines; i++ {
		visited = append(visited, make([]bool, numDigits))
		for j := 0; j < numDigits; j++ {
			if data[i][j] == 9 {
				data[i][j] = 1
			} else {
				data[i][j] = 0
			}
			visited[i][j] = false
		}
	}
	basins := funk.Map(lowPoints, func(point [2]int) int { return search(data, visited, point[0], point[1]) }).([]int)
	sort.Ints(basins)
	basins = funk.Reverse(basins).([]int)
	fmt.Printf("Found basins: %v\n", basins)
	return basins[0] * basins[1] * basins[2]
}

func search(data [][]int, visited [][]bool, i int, j int) int {
	visited[i][j] = true
	basin := 1
	if data[i][j] == 0 && i > 0 && !visited[i-1][j] {
		visited[i-1][j] = true
		basin += search(data, visited, i-1, j)
	}
	if data[i][j] == 0 && j > 0 && !visited[i][j-1] {
		visited[i][j-1] = true
		basin += search(data, visited, i, j-1)
	}
	if data[i][j] == 0 && i < len(data)-1 && !visited[i+1][j] {
		visited[i+1][j] = true
		basin += search(data, visited, i+1, j)
	}
	if data[i][j] == 0 && j < len(data[0])-1 && !visited[i][j+1] {
		visited[i][j+1] = true
		basin += search(data, visited, i, j+1)
	}
	return basin
}
