package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed input.txt
var input string

//go:embed test1.txt
var test string

func main() {
	utils.Test(solve(test), 45000)
	fmt.Println(solve(input))
}

func solve(input string) int {
	input = strings.ReplaceAll(input, "\n", "+")
	s := strings.Split(input, "++")
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
