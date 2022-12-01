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
	input = strings.ReplaceAll(input, "\n", "+")
	fmt.Println(input)
	s := strings.Split(input, "++")
	i := solve(s)
	fmt.Println(i)
}

func solve(s []string) int {
	ints := []int{}

	for _, str := range s {
		sums := strings.Split(str, "+")

		s := 0
		for _, n := range sums {
			i, _ := strconv.Atoi(n)
			s += i
		}

		ints = append(ints, s)
	}

	max := 0

	for _, i := range ints {
		if i > max {
			max = i
		}
	}

	return max
}
