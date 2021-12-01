package day1

// Source: day1.go
import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := `199
200
208
210
200
207
240
269
260
263`
	expected := 7
	d1 := NewDay1(input)
	result := d1.Part1()
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestPart2(t *testing.T) {
	input := `199
200
208
210
200
207
240
269
260
263`
	expected := 5
	d1 := NewDay1(input)
	result := d1.Part2()
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
