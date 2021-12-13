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
	fmt.Println(i)
}

func solve(s []string) int {
	counts := map[int]int{}

	half := len(s) / 2
	_ = half

	for i := 0; i < len(s[0]); i++ {
		counts[i] = 0
	}

	for _, line := range s {
		for i, char := range line {
			if string(char) == "1" {
				counts[i]++
			}
		}
	}

	gamma := 0
	epsilon := 0

	fmt.Println(counts)

	for _, c := range counts {
		if c > half {
			gamma++
		} else {
			epsilon++
		}

		gamma *= 2
		epsilon *= 2
	}

	return gamma * epsilon
}
