package day4

import "testing"

var input = `7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7`

// Test Part 1
func TestPart1(t *testing.T) {
	expected := 4512
	d := NewDay(input)
	result := d.Parts(1)
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func BenchmarkPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		d := NewDay(input)
		d.Parts(1)
	}
}

func TestPart2(t *testing.T) {
	expected := 1924
	d := NewDay(input)
	result := d.Parts(2)
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func BenchmarkPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		d := NewDay(input)
		d.Parts(2)
	}
}
