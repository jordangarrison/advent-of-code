package util

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// Read a file and return its contents as a string.
func readFile(path string) (string, error) {
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(contents)), nil
}

// Get data for a given day and part
func GetData(day int, part int) string {
	filepath := "./data/day" + strconv.Itoa(day) + "/part" + strconv.Itoa(part) + ".txt"
	fmt.Println("Reading data from", filepath)
	contents, err := readFile(filepath)
	if err != nil {
		panic(err)
	}
	return contents
}
