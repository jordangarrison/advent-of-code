package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jordangarrison/advent-of-code/2021/go/day1"
	"github.com/jordangarrison/advent-of-code/2021/go/day2"
	"github.com/jordangarrison/advent-of-code/2021/go/day3"
	"github.com/jordangarrison/advent-of-code/2021/go/day4"
	"github.com/jordangarrison/advent-of-code/2021/go/day5"
	"github.com/jordangarrison/advent-of-code/2021/go/day6"
	"github.com/jordangarrison/advent-of-code/2021/go/day7"
	"github.com/jordangarrison/advent-of-code/2021/go/util"
)

func main() {
	// Get arguments from command line
	args := os.Args[1:]

	if args[0] == "pull" {
		day, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Error: Invalid day")
			os.Exit(1)
		}
		data, err := util.PullData(day)
		fmt.Printf("%s\n", data)
		return
	}

	for _, arg := range args {
		// Run day corresponding to argument
		switch arg {
		case "1":
			part1 := util.GetData(1, 1)
			util.Stats(day1.Run, part1)
		case "2":
			part1 := util.GetData(2, 1)
			util.Stats(day2.Run, part1)
		case "3":
			part1 := util.GetData(3, 1)
			util.Stats(day3.Run, part1)
		case "4":
			part1 := util.GetData(4, 1)
			util.Stats(day4.Run, part1)
		case "5":
			part1 := util.GetData(5, 1)
			util.Stats(day5.Run, part1)
		case "6":
			part1 := util.GetData(6, 1)
			util.Stats(day6.Run, part1)
		case "7":
			part1 := util.GetData(7, 1)
			util.Stats(day7.Run, part1)
		default:
			fmt.Printf("Day %s not implemented\n", arg)
		}
	}
}
