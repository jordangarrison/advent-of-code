package day7

import "testing"

var input = `16,1,2,0,4,2,7,1,2,14`

// Test Part 1
func TestPart1(t *testing.T) {
	expected := 37
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
	expected := 168
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

func TestNthTriangular(t *testing.T) {
	expectations := []int{0, 1, 3, 6, 10, 15, 21, 28, 36, 45, 55, 66}
	for i, e := range expectations {
		if nthTriangular(i) != e {
			t.Errorf("Expected %v, got %v", e, nthTriangular(i))
		}
	}
}
