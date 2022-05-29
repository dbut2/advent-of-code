package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	s := strings.Split(input, ",")
	i := solve(s)
	fmt.Println(i)
}

func solve(s []string) int {
	ints := stringsToInts(s)

	max := 0
	for _, i := range ints {
		if i > max {
			max = i
		}
	}

	best := -1
	for i := 0; i < max; i++ {
		total := 0
		for _, crab := range ints {
			total += abs(crab - i)
		}

		if total < best || best == -1 {
			best = total
		} else {
		}
	}

	return best
}

// return a slice of ints from a slice of strings
func stringsToInts(s []string) []int {
	ints := []int{}
	for _, str := range s {
		i, err := strconv.Atoi(str)
		if err != nil {
			panic(err.Error())
		}
		ints = append(ints, i)
	}
	return ints
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
