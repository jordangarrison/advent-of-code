package main

import (
	"fmt"
	"os"

	util "github.com/jordangarrison/advent-of-code/2023/go/util"

	day1 "github.com/jordangarrison/advent-of-code/2023/go/day1"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Please provide a day number")
		os.Exit(1)
	}
	// Given some numbers run the corresponding day for AOC
	for _, arg := range args {
		switch arg {
		case "1":
			data := util.GetData(1)
			day1.Run(data)
		}
	}
}
