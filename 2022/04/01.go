package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/dbut2/advent-of-code/utils"
)

//go:embed input.txt
var input string

func main() {
	s := strings.Split(input, "\n")
	i := solve(s)
	fmt.Println(i)
}

func solve(s []string) int {
	full := 0

	for _, str := range s {

		elves := strings.Split(str, ",")

		e1 := utils.Stis(strings.Split(elves[0], "-"))
		e2 := utils.Stis(strings.Split(elves[1], "-"))

		if e1[0] <= e2[0] && e1[1] >= e2[1] {
			full++
		}
		if e2[0] <= e1[0] && e2[1] >= e1[1] {
			full++
		}
		if e1[0] == e2[0] && e1[1] == e2[1] {
			full--
		}

	}

	return full
}
