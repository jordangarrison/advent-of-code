package util

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// Read a file and return its contents as a string.
func readFile(filepath string) (string, error) {
	contents, err := ioutil.ReadFile(filepath)
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

func PullData(day int) ([]byte, error) {
	client := new(http.Client)
	url := "https://adventofcode.com/2019/day/" + strconv.Itoa(day) + "/input"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.AddCookie(&http.Cookie{Name: "session", Value: os.Getenv("AOC_COOKIE")})
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
