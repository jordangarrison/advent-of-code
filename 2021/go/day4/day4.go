package day4

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/thoas/go-funk"
)

type Day struct {
	input string
}

func NewDay(input string) *Day {
	return &Day{input}
}

func Run(input string) {
	d := NewDay(input)

	part1Result := d.Parts(1)
	fmt.Printf("Part 1: %v\n", part1Result)

	part2Result := d.Parts(2)
	fmt.Printf("Part 2: %v\n", part2Result)
}

func (d *Day) Parts(part int) int {
	blocks := strings.Split(d.input, "\n\n")
	inputString := blocks[0]
	input := funk.Map(strings.Split(inputString, ","), func(s string) int {
		i, _ := strconv.Atoi(s)
		return i
	}).([]int)
	boardsString := blocks[1:]
	boards := funk.Map(boardsString, func(s string) Board {
		return parseBoard(s)
	}).([]Board)
	var bingos []Board
	for _, num := range input {
		for _, board := range boards {
			board.Mark(num)
			bingo := board.CheckBingo()
			if bingo && part == 1 {
				return board.sumBoard(num)
			} else if bingo && part == 2 {
				if !funk.Contains(bingos, board) {
					bingos = append(bingos, board)
				}
				if len(bingos) == len(boards) {
					return board.sumBoard(num)
				}
			}
		}
	}
	return 0
}

func (b *Board) sumBoard(num int) int {
	sum := 0
	for _, row := range b.Places {
		for _, place := range row {
			if !place.marked {
				sum += place.number
			}
		}
	}
	return sum * num
}

func parseBoard(input string) Board {
	var board Board
	for _, line := range strings.Split(input, "\n") {
		var places []Place
		for _, field := range strings.Fields(line) {
			num, _ := strconv.Atoi(field)
			place := Place{num, false}
			places = append(places, place)
		}
		board.Places = append(board.Places, places)
	}
	return board
}

type Board struct {
	Places [][]Place
}

type Place struct {
	number int
	marked bool
}

func (b *Board) Mark(num int) {
	for i, row := range b.Places {
		for j, place := range row {
			if place.number == num && !place.marked {
				b.Places[i][j].marked = true
			}
		}
	}
}

// Apparently diagonals don't count
func (b *Board) CheckBingo() bool {
	// var diagonal []Place
	for _, row := range b.Places {
		if isBingo(row) {
			return true
		}
		tb := transpose(b.Places)
		for _, row := range tb {
			if isBingo(row) {
				return true
			}
		}
	}
	// return isBingo(diagonalT)
	return false
}

func isBingo(places []Place) bool {
	for _, place := range places {
		if !place.marked {
			return false
		}
	}
	return true
}

func transpose(a [][]Place) [][]Place {
	transposed := make([][]Place, len(a[0]))
	for i := range transposed {
		transposed[i] = make([]Place, len(a))
	}
	for i, row := range a {
		for j, val := range row {
			transposed[j][i] = val
		}
	}
	return transposed
}
