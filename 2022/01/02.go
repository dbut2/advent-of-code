package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	input = strings.ReplaceAll(input, "\n", "+")
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

	sort.Ints(ints)

	c := 0

	for i := len(ints) - 1; i > len(ints)-4; i-- {
		c += ints[i]
	}

	return c
}
