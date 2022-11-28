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
	numbers := []int{}

	for _, line := range s {
		i, err := strconv.Atoi(line)
		if err != nil {
			panic(err.Error())
		}

		numbers = append(numbers, i)
	}

	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			a, b := numbers[i], numbers[j]

			if a+b == 2020 {
				return a * b
			}
		}
	}

	return 0
}
