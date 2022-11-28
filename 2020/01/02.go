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
			for k := j + 1; k < len(numbers); k++ {
				a, b, c := numbers[i], numbers[j], numbers[k]

				if a+b+c == 2020 {
					return a * b * c
				}
			}
		}
	}

	return 0
}
