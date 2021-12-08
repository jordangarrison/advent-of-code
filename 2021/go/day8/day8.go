package day8

import (
	"fmt"
	"sort"
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

type inputOutput struct {
	input  []string
	output []string
}

func (d *Day) readInput() []inputOutput {
	var digits []inputOutput
	for _, line := range strings.Split(d.input, "\n") {
		data := strings.Split(line, "|")
		input := data[0]
		output := data[1]
		inputStrings := strings.Fields(input)
		outputStrings := strings.Fields(output)
		io := inputOutput{input: inputStrings, output: outputStrings}
		digits = append(digits, io)
	}
	return digits
}

var rules = map[int]int{
	1: 2,
	4: 4,
	7: 3,
	8: 7,
}

func (d *Day) Part1() int {
	digits := d.readInput()
	count := 0
	for i := range digits {
		for j := range digits[i].output {
			digitLength := len(digits[i].output[j])
			// fmt.Printf("%v: %v\n", digits[i][j], digitLength)
			for _, v := range rules {
				if digitLength == v {
					count += 1
				}
			}
		}
	}
	return count
}

type digitFormat struct {
	top        []string
	upperLeft  []string
	upperRight []string
	middle     []string
	lowerLeft  []string
	lowerRight []string
	bottom     []string
}

type process func([]string, []string) []string

func shouldAdd(processor process, chars []string, charsToAdd ...string) []string {
	charlist := processor(chars, charsToAdd)
	sort.Strings(charlist)
	return funk.Uniq(charlist).([]string)
}

func (d *Day) Part2() int {
	digits := d.readInput()
	for i := range digits {
		df := digitFormat{
			top:        []string{"a", "b", "c", "d", "e", "f", "g"},
			middle:     []string{"a", "b", "c", "d", "e", "f", "g"},
			bottom:     []string{"a", "b", "c", "d", "e", "f", "g"},
			upperLeft:  []string{"a", "b", "c", "d", "e", "f", "g"},
			upperRight: []string{"a", "b", "c", "d", "e", "f", "g"},
			lowerLeft:  []string{"a", "b", "c", "d", "e", "f", "g"},
			lowerRight: []string{"a", "b", "c", "d", "e", "f", "g"},
		}
		for j := range digits[i].input {
			digitString := digits[i].input[j]
			digitLength := len(digitString)
			chars := strings.Split(digitString, "")
			// fmt.Printf("%v: %v\n", digitString, digitLength)
			switch digitLength {
			case rules[1]:
				df.top = shouldAdd(funk.LeftJoinString, df.top, chars...)
				df.bottom = shouldAdd(funk.LeftJoinString, df.bottom, chars...)
				df.middle = shouldAdd(funk.LeftJoinString, df.middle, chars...)
				df.upperLeft = shouldAdd(funk.LeftJoinString, df.upperLeft, chars...)
				df.lowerLeft = shouldAdd(funk.LeftJoinString, df.lowerLeft, chars...)
				df.upperRight = shouldAdd(funk.InnerJoinString, df.upperRight, chars...)
				df.lowerRight = shouldAdd(funk.InnerJoinString, df.lowerRight, chars...)
			case rules[4]:
				df.top = shouldAdd(funk.LeftJoinString, df.top, chars...)
				df.bottom = shouldAdd(funk.LeftJoinString, df.bottom, chars...)
				df.upperLeft = shouldAdd(funk.InnerJoinString, df.upperLeft, chars...)
				df.middle = shouldAdd(funk.InnerJoinString, df.middle, chars...)
				df.middle = shouldAdd(funk.OuterJoinString, df.upperLeft, df.middle...)
				df.lowerLeft = shouldAdd(funk.LeftJoinString, df.lowerLeft, chars...)
				df.upperRight = shouldAdd(funk.InnerJoinString, df.upperRight, chars...)
				df.lowerRight = shouldAdd(funk.InnerJoinString, df.lowerRight, chars...)
			case rules[7]:
				df.top = shouldAdd(funk.InnerJoinString, df.top, chars...)
				df.bottom = shouldAdd(funk.LeftJoinString, df.bottom, chars...)
				df.middle = shouldAdd(funk.LeftJoinString, df.middle, chars...)
				df.upperLeft = shouldAdd(funk.LeftJoinString, df.upperLeft, chars...)
				df.lowerLeft = shouldAdd(funk.LeftJoinString, df.lowerLeft, chars...)
				df.upperRight = shouldAdd(funk.InnerJoinString, df.upperRight, chars...)
				df.lowerRight = shouldAdd(funk.InnerJoinString, df.lowerRight, chars...)
			case rules[8]:
				df.top = shouldAdd(funk.InnerJoinString, df.top, chars...)
				df.bottom = shouldAdd(funk.InnerJoinString, df.bottom, chars...)
				df.middle = shouldAdd(funk.InnerJoinString, df.middle, chars...)
				df.upperLeft = shouldAdd(funk.InnerJoinString, df.upperLeft, chars...)
				df.lowerLeft = shouldAdd(funk.InnerJoinString, df.lowerLeft, chars...)
				df.upperRight = shouldAdd(funk.InnerJoinString, df.upperRight, chars...)
				df.lowerRight = shouldAdd(funk.InnerJoinString, df.lowerRight, chars...)
			}
		}
		fmt.Printf("%+v\n", df)
	}
	return 0
}
