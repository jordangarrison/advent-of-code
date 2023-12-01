package util

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Read a file and return its contents as a string.
func readFile(filepath string) (string, error) {
	contents, err := os.ReadFile(filepath)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(contents)), nil
}

// Get data for a given day and part
func GetData(day int) string {
	filepath := "./data/day" + strconv.Itoa(day) + ".txt"
	fmt.Println("Reading data from", filepath)
	contents, err := readFile(filepath)
	if err != nil {
		panic(err)
	}
	return contents
}
