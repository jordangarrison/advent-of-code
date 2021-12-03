package day3

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/thoas/go-funk"
)

type Day3 struct {
	input string
}

func NewDay3(input string) *Day3 {
	return &Day3{input}
}

func Run(input string) {
	d := NewDay3(input)

	part1Result := d.Part1()
	fmt.Printf("Part 1: %v\n", part1Result)

	part2Result := d.Part2()
	fmt.Printf("Part 2: %v\n", part2Result)
}

func (d *Day3) Part1() int {
	lines := strings.Split(d.input, "\n")
	numBits := len(lines[0])
	diagnostic := make([]int, numBits)
	for _, line := range lines {
		runes := []rune(line)
		for i, char := range runes {
			buf := make([]byte, 1)
			_ = utf8.EncodeRune(buf, char)
			bit, _ := strconv.Atoi(string(buf))
			diagnostic[i] += bit
			// fmt.Printf("%v, %v, %v\n", i, bit, diagnostic)
		}
	}
	majority := len(lines) / 2
	gamma := ""
	epsilon := ""
	for _, v := range diagnostic {
		if v >= majority {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
		// fmt.Printf("%v, %v, %v\n", diagnostic, gamma, epsilon)
	}
	gammaValue, _ := strconv.ParseInt(gamma, 2, 64)
	epsilonValue, _ := strconv.ParseInt(epsilon, 2, 64)
	return int(gammaValue * epsilonValue)
}

type LifeSupport struct {
	zero []string
	one  []string
}

func (d *Day3) Part2() int {
	lines := strings.Split(d.input, "\n")
	numBits := len(lines[0])
	lifeSupport := make([]LifeSupport, numBits)
	for _, line := range lines {
		runes := []rune(line)
		for i, char := range runes {
			buf := make([]byte, 1)
			_ = utf8.EncodeRune(buf, char)
			bit, _ := strconv.Atoi(string(buf))
			if bit == 1 {
				lifeSupport[i].one = append(lifeSupport[i].one, line)
			} else {
				lifeSupport[i].zero = append(lifeSupport[i].zero, line)
			}
		}
	}

	majority := len(lines) / 2
	// fmt.Printf("Majority: %v\n%+v\n", majority, lifeSupport)
	o2 := []string{}
	co2 := []string{}
	o2String := ""
	co2String := ""
	for i, lf := range lifeSupport {
		if i == 0 {
			if len(lf.one) >= majority {
				o2 = append(o2, lf.one...)
				co2 = append(co2, lf.zero...)
			} else {
				o2 = append(o2, lf.zero...)
				co2 = append(co2, lf.one...)
			}
		} else {
			o2one := funk.IntersectString(lf.one, o2)
			o2zero := funk.IntersectString(lf.zero, o2)
			co2one := funk.IntersectString(lf.one, co2)
			co2zero := funk.IntersectString(lf.zero, co2)
			if len(o2one) >= len(o2zero) {
				o2 = o2one
			} else {
				o2 = o2zero
			}
			if len(co2zero) <= len(co2one) {
				co2 = co2zero
			} else {
				co2 = co2one
			}
		}
		if len(o2) == 1 {
			o2String = o2[0]
			// fmt.Printf("Final O2: %v\n", o2String)
		}
		if len(co2) == 1 {
			co2String = co2[0]
			// fmt.Printf("Final CO2: %v\n", co2String)
		}
		// fmt.Printf("LS: %+v\nO2: %+v\nCO2: %+v\n\n", lf, o2, co2)
	}

	o2Value, _ := strconv.ParseInt(o2String, 2, 64)
	co2Value, _ := strconv.ParseInt(co2String, 2, 64)
	return int(o2Value * co2Value)
}
