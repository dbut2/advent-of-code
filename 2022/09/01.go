package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/dbut2/advent-of-code/utils"
)

//go:embed input.txt
var input string

//go:embed test1.txt
var test string

func main() {
	fmt.Println("Test")
	fmt.Println(do(test))
	fmt.Println()
	fmt.Println("Solution")
	fmt.Println(do(input))
}

func do(s string) int {
	strs := strings.Split(s, "\n")
	return solve(strs)
}

func solve(s []string) int {

	visits := utils.Set[string]{}

	hx, hy := 0, 0
	tx, ty := 0, 0

	visits.Add(fmt.Sprintf("%d,%d", tx, ty))

	for _, line := range s {
		m := strings.Split(line, " ")
		dir := m[0]
		amt := utils.Sti(m[1])

		for i := 0; i < amt; i++ {
			switch dir {
			case "U":
				hx--
			case "D":
				hx++
			case "L":
				hy--
			case "R":
				hy++
			}

			if utils.Abs(hx-tx) <= 1 && utils.Abs(hy-ty) <= 1 {
				continue
			}

			tx += utils.Sign(hx - tx)
			ty += utils.Sign(hy - ty)

			visits.Add(fmt.Sprintf("%d,%d", tx, ty))
		}
	}

	return len(visits)
}
