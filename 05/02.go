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
	s := strings.Split(input, "\n")
	i := solve(s)
	fmt.Println(i)
}

func solve(s []string) int {
	grid := map[int]map[int]int{}

	for i := 0; i < 1000; i++ {
		grid[i] = map[int]int{}
		for j := 0; j < 1000; j++ {
			grid[i][j] = 0
		}
	}

	for _, line := range s {
		coords := strings.Split(line, " ")

		a := stringsToInts(strings.Split(coords[0], ","))
		b := stringsToInts(strings.Split(coords[2], ","))

		x := InclusiveRange(a[0], b[0])
		y := InclusiveRange(a[1], b[1])

		x, y = makeSameLength(x, y)

		for i := range x {
			grid[x[i]][y[i]]++
		}
	}

	count := 0

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] > 1 {
				count++
			}
		}
	}

	return count
}

// stringsToInts returns a slice of ints from a slice of strings
func stringsToInts(s []string) []int {
	var i []int
	for _, v := range s {
		s, err := strconv.Atoi(v)
		if err != nil {
			panic(err.Error())
		}
		i = append(i, s)
	}
	return i
}

func InclusiveRange(a, b int) []int {
	var r []int
	if a > b {
		r = make([]int, a-b+1)
		for i := range r {
			r[i] = a - i
		}
	} else {
		r = make([]int, b-a+1)
		for i := range r {
			r[i] = a + i
		}
	}
	return r
}

func makeSameLength(x, y []int) ([]int, []int) {
	if len(x) > len(y) {
		for i := len(y); i < len(x); i++ {
			y = append(y, y[0])
		}
	} else if len(y) > len(x) {
		for i := len(x); i < len(y); i++ {
			x = append(x, x[0])
		}
	}
	return x, y
}
