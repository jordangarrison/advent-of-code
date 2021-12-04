package util

import (
	"fmt"
	"time"
)

// function which get statistics about the given function such as timings, number of calls, etc.
func Stats(f func(string), s string) {
	start := time.Now()
	f(s)
	duration := time.Since(start)
	fmt.Printf("Call took %s\n", duration)
}
