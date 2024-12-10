package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/sti"
	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed input.txt
var input string

//go:embed test1.txt
var test string

func main() {
	utils.Test(solve(test), 2)
	fmt.Println(solve(input))
}

func solve(input string) int {
	s := utils.ParseInput(input)
	full := 0

	for _, str := range s {

		elves := strings.Split(str, ",")

		e1 := sti.Stis(strings.Split(elves[0], "-"))
		e2 := sti.Stis(strings.Split(elves[1], "-"))

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
