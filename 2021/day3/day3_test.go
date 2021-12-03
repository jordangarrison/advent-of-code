package day3

import "testing"

var input = `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`

// Test Part 1
func TestPart1(t *testing.T) {
	expected := 198
	d2 := NewDay3(input)
	result := d2.Part1()
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func BenchmarkPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		d2 := NewDay3(input)
		d2.Part1()
	}
}

func TestPart2(t *testing.T) {
	expected := 230
	d2 := NewDay3(input)
	result := d2.Part2()
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func BenchmarkPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		d2 := NewDay3(input)
		d2.Part2()
	}
}
