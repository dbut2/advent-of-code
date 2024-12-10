package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/lists"
	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed input.txt
var input string

//go:embed test1.txt
var test string

func main() {
	utils.Test(solve(test), 157)
	fmt.Println(solve(input))
}

func solve(input string) int {
	s := utils.ParseInput(input)

	count := 0

	for _, str := range s {

		letters := strings.Split(str, "")

		first := []string{}
		second := []string{}

		for i, c := range letters {

			if i < len(letters)/2 {
				first = append(first, c)
			} else {
				second = append(second, c)
			}

		}

		inboth := lists.Intersection(first, second)[0]

		count += priority(inboth)

	}

	return count
}

func priority(a string) int {
	b := []byte(a)[0]

	if b > 96 {
		b -= 58
	}

	b -= 38

	return int(b)
}
