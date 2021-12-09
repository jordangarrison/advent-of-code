package day9

import "testing"

var input = `2199943210
3987894921
9856789892
8767896789
9899965678`

func TestPart1(t *testing.T) {
	expected := 15
	d := NewDay(input)
	result := d.Part1()
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func BenchmarkPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		d := NewDay(input)
		d.Part1()
	}
}

func TestPart2(t *testing.T) {
	expected := 1134
	d := NewDay(input)
	result := d.Part2()
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func BenchmarkPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		d := NewDay(input)
		d.Part2()
	}
}
