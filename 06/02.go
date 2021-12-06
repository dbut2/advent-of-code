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
	days := 256

	counts := map[int]int{}

	for i := 0; i < 9; i++ {
		counts[i] = 0
	}

	for _, f := range stringsToInts(s) {
		counts[f]++
	}

	for i := 0; i < days; i++ {
		newFishes := counts[0]
		for j := 1; j < 9; j++ {
			counts[j-1] = counts[j]
		}
		counts[8] = newFishes
		counts[6] = counts[6] + newFishes
	}

	count := 0

	for _, c := range counts {
		count += c
	}

	return count
}

// return a list of ints from a list of strings
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
