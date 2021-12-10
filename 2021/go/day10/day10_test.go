package day10

import "testing"

var input = `[({(<(())[]>[[{[]{<()<>>
[(()[<>])]({[<{<<[]>>(
{([(<{}[<>[]}>{[]{[(<()>
(((({<>}<{<{<>}{[]{[]{}
[[<[([]))<([[{}[[()]]]
[{[{({}]{}}([{[{{{}}([]
{<[[]]>}<{[{[{[]{()[[[]
[<(<(<(<{}))><([]([]()
<{([([[(<>()){}]>(<<{{
<{([{{}}[<[[[<>{}]]]>[]]`

func TestPart1(t *testing.T) {
	expected := 26397
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
