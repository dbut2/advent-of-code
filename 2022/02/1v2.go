package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed input.txt
var input string

//go:embed test1.txt
var test string

func main() {
	utils.Test(solve(test), 15)
	fmt.Println(solve(input))
}

func solve(input string) int {
	s := utils.ParseInput(input)

	m := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	count := 0

	for _, str := range s {
		plays := strings.Split(str, " ")

		pa := m[plays[0]]
		pb := m[plays[1]]

		score := ((pb - pa + 1 + 6) % 3) * 3

		count += pb + score
	}

	return count
}
