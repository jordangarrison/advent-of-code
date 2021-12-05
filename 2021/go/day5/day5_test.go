package day5

import "testing"

var input = `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`

// Test Part 1
func TestPart1(t *testing.T) {
	expected := 5
	d := NewDay(input)
	result := d.Parts(false)
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func BenchmarkPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		d := NewDay(input)
		d.Parts(false)
	}
}

func TestPart2(t *testing.T) {
	expected := 12
	d := NewDay(input)
	result := d.Parts(true)
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func BenchmarkPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		d := NewDay(input)
		d.Parts(true)
	}
}
