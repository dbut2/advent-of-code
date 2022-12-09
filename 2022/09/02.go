package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/dbut2/advent-of-code/utils"
)

//go:embed input.txt
var input string

//go:embed test2.txt
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

	knots := utils.Fill2D(10, 2, 0)

	visits.Add(fmt.Sprintf("%d,%d", knots[len(knots)-1][0], knots[len(knots)-1][1]))

	for _, line := range s {
		m := strings.Split(line, " ")
		dir := m[0]
		amt := utils.Sti(m[1])

		for i := 0; i < amt; i++ {
			switch dir {
			case "U":
				knots[0][0]--
			case "D":
				knots[0][0]++
			case "L":
				knots[0][1]--
			case "R":
				knots[0][1]++
			}

			for j := 1; j < len(knots); j++ {
				if utils.Abs(knots[j-1][0]-knots[j][0]) <= 1 && utils.Abs(knots[j-1][1]-knots[j][1]) <= 1 {
					continue
				}

				knots[j][0] += utils.Sign(knots[j-1][0] - knots[j][0])
				knots[j][1] += utils.Sign(knots[j-1][1] - knots[j][1])
			}

			visits.Add(fmt.Sprintf("%d,%d", knots[len(knots)-1][0], knots[len(knots)-1][1]))
		}
	}

	return len(visits)
}
