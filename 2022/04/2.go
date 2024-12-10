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
	utils.Test(solve(test), 4)
	fmt.Println(solve(input))
}

func solve(input string) int {
	s := utils.ParseInput(input)
	overlap := 0

	for _, str := range s {
		elves := strings.Split(str, ",")

		e1 := sti.Stis(strings.Split(elves[0], "-"))
		e2 := sti.Stis(strings.Split(elves[1], "-"))

		e1s, e1e := e1[0], e1[1]
		e2s, e2e := e2[0], e2[1]

		if e1s <= e2s && e2s <= e1e {
			overlap++
			continue
		}

		if e2s <= e1s && e1s <= e2e {
			overlap++
			continue
		}
	}

	return overlap
}
