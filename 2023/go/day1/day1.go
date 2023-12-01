package day1

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type day1 struct {
	input string
}

func NewDay1(input string) *day1 {
	return &day1{input: input}
}

func Run(input string) int {
	d1 := NewDay1(input)
	fmt.Println("Day 1 Part 1:", d1.Part1())
	fmt.Println("Day 1 Part 2:", d1.Part2())
	return 0
}

func (d day1) Part1() int {
	// for each line in the input
	// scan the line for numbers and form an array
	// the first number of the array and last number
	// of the array are combined and returned
	sum := 0
	for _, line := range strings.Split(d.input, "\n") {
		digits := []rune{}
		for _, char := range line {
			if unicode.IsDigit(char) {
				digits = append(digits, char)
			}
		}
		rowStr := fmt.Sprintf("%c%c", digits[0], digits[len(digits)-1])
		row, _ := strconv.Atoi(rowStr)
		sum += row
	}

	return sum
}

func (d day1) Part2() int {
	// for each line in the input
	// scan the line for numbers and form an array
	// the first number of the array and last number
	// of the array are combined and returned
	sum := 0
	words := [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	// edgeCases := map[string]int{
	// 	"oneight":   18,
	// 	"twoone":    21,
	// 	"threeight": 38,
	// 	"fiveight":  58,
	// 	"sevenine":  79,
	// 	"eighthree": 83,
	// }
	for _, line := range strings.Split(d.input, "\n") {
		digits := []rune{}
		word := ""
		for _, char := range line {
			if unicode.IsDigit(char) {
				digits = append(digits, char-48)
				word = ""
			} else if unicode.IsLetter(char) {
				word += string(char)
				for index, w := range words {
					if strings.Contains(word, w) {
						digits = append(digits, rune(index+1))
						word = ""
					}
				}
			}
		}
		// get the first and last digits of the array
		rowStr := fmt.Sprintf("%v%v", digits[0], digits[len(digits)-1])
		row, err := strconv.Atoi(rowStr)
		if err != nil {
			panic(err)
		}
		sum += row
	}

	return sum
}
