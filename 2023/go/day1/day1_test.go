package day1

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := `1abc2
	pqr3stu8vwx
	a1b2c3d4e5f
	treb7uchet`
	expected := 142
	d := NewDay1(input)
	result := d.Part1()
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestPart2(t *testing.T) {
	input := `two1nine
	eightwothree
	abcone2threexyz
	xtwone3four
	4nineeightseven2
	zoneight234
	7pqrstsixteen`
	expected := 281
	d1 := NewDay1(input)
	result := d1.Part2()
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
