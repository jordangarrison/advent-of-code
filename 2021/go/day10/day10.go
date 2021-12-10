package day10

import (
	"fmt"
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

func (d *Day) readInput() [][]string {
	var symbols [][]string
	for _, line := range strings.Split(d.input, "\n") {
		symbols = append(symbols, strings.Split(line, ""))
	}
	return symbols
}

type Stack []string

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(str string) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

func (d *Day) Part1() int {
	symbols := d.readInput()
	score := map[string]int{
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}
	var invalidSymbols []string
	for _, line := range symbols {
		var stack Stack
		for _, symbol := range line {
			switch symbol {
			case "(", "[", "{", "<":
				stack.Push(symbol)
			case ")", "]", "}", ">":
				matchSymbol, err := match(symbol)
				if err != nil {
					invalidSymbols = append(invalidSymbols, symbol)
				}
				stackSymbol, _ := stack.Pop()
				if stack.IsEmpty() {
				} else if stackSymbol != matchSymbol {
					// fmt.Printf("%v %v\n", stackSymbol, symbol)
					invalidSymbols = append(invalidSymbols, symbol)
				}
			}
		}
	}
	// fmt.Printf("Invalid Symbols: %v\n", invalidSymbols)
	return funk.SumInt(funk.Map(invalidSymbols, func(s string) int { return score[s] }).([]int))
}

func match(char string) (string, error) {
	switch char {
	case ")":
		return "(", nil
	case "]":
		return "[", nil
	case "}":
		return "{", nil
	case ">":
		return "<", nil
	default:
		return char, fmt.Errorf("invalid symbol: %v", char)
	}
}

func (d *Day) Part2() int {
	return 0
}
