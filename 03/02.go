package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed 01.txt
var input string

func main() {
	s := strings.Split(input, "\n")
	i := solve(s)
	fmt.Println(i, "done")
}

func solve(s []string) int {
	oxygenStr := filter(true, 0, s)[0]
	co2Str := filter(false, 0, s)[0]

	o := binToDec(oxygenStr)
	c := binToDec(co2Str)

	return o * c
}

func filter(max bool, bit int, s []string) []string {
	if len(s) <= 1 || bit > len(s[0]) {
		return s
	}

	common, equal := mostCommonNthDigit(s, bit)

	fmt.Println(s)

	var r []string

	f := common
	if equal {
		f = "1"
	}

	if max {
		for _, line := range s {
			if string(line[bit]) == f {
				r = append(r, line)
			}
		}
	} else {
		for _, line := range s {
			if string(line[bit]) != f {
				r = append(r, line)
			}
		}
	}

	fmt.Println(r)

	return filter(max, bit+1, r)
}

func mostCommonNthDigit(s []string, n int) (string, bool) {
	c := 0

	for _, line := range s {
		if string(line[n]) == "1" {
			c++
		}
	}

	if 2*c == len(s) {
		return "", true
	}

	if 2*c > len(s) {
		return "1", false
	} else {
		return "0", false
	}
}

func binToDec(b string) int {
	d := 0
	for _, char := range b {
		d *= 2
		if string(char) == "1" {
			d++
		}
	}
	return d
}
