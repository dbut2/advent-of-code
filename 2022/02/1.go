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

	mapb := map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	count := 0

	for _, str := range s {
		plays := strings.Split(str, " ")

		add := 0

		pb := mapb[plays[1]]

		switch plays[0] {
		case "A":
			switch plays[1] {
			case "X":
				add = 3
			case "Y":
				add = 6
			case "Z":
				add = 0
			}
		case "B":
			switch plays[1] {
			case "X":
				add = 0
			case "Y":
				add = 3
			case "Z":
				add = 6
			}
		case "C":
			switch plays[1] {
			case "X":
				add = 6
			case "Y":
				add = 0
			case "Z":
				add = 3
			}
		}

		count += pb + add
	}

	return count
}
