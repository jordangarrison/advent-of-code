package day6

import "testing"

var input = `3,4,3,1,2`

// Test Part 1
func TestPart1(t *testing.T) {
	expected := 5934
	d := NewDay(input)
	result := d.Parts(80)
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func BenchmarkPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		d := NewDay(input)
		d.Parts(80)
	}
}

func TestPart2(t *testing.T) {
	expected := 26984457539
	d := NewDay(input)
	result := d.Parts(256)
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func BenchmarkPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		d := NewDay(input)
		d.Parts(256)
	}
}

func Benchmark1Day(b *testing.B) {
	for i := 0; i < b.N; i++ {
		d := NewDay(input)
		d.Parts(1)
	}
}

func Benchmark1000Days(b *testing.B) {
	for i := 0; i < b.N; i++ {
		d := NewDay(input)
		d.Parts(1000)
	}
}
