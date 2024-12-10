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

		if a[0] == b[0] {

			if a[1] > b[1] {
				for i := b[1]; i <= a[1]; i++ {
					grid[a[0]][i]++
				}
			} else {
				for i := a[1]; i <= b[1]; i++ {
					grid[a[0]][i]++
				}
			}
		}

		if a[1] == b[1] {

			if a[0] > b[0] {
				for i := b[0]; i <= a[0]; i++ {
					grid[i][a[1]]++
				}
			} else {
				for i := a[0]; i <= b[0]; i++ {
					grid[i][a[1]]++
				}
			}
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
