package day2

import (
	"testing"
)

var input = `forward 5
down 5
forward 8
up 3
down 8
forward 2`

func TestPart1(t *testing.T) {
	expected := 150
	d2 := NewDay2(input)
	result := d2.Part1(input)
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestPart2(t *testing.T) {
	expected := 900
	d2 := NewDay2(input)
	result := d2.Part2(input)
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
