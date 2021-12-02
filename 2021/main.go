package main

import (
	"fmt"
	"os"

	"github.com/jordangarrison/advent-of-code/2021/day1"
	"github.com/jordangarrison/advent-of-code/2021/day2"
	"github.com/jordangarrison/advent-of-code/2021/util"
)

func main() {
	// Get arguments from command line
	args := os.Args[1:]

	for _, arg := range args {
		// Run day corresponding to argument
		switch arg {
		case "1":
			part1, err := util.GetData(1, 1)
			if err != nil {
				panic(err)
			}
			day1.Run(part1)
		case "2":
			part1, err := util.GetData(2, 1)
			if err != nil {
				panic(err)
			}
			day2.Run(part1)
		default:
			fmt.Printf("Day %s not implemented\n", arg)
		}
	}
}
