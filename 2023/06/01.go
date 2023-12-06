package main

import (
	"embed"
	"fmt"

	"github.com/dbut2/advent-of-code/pkg/strings"
	"github.com/dbut2/advent-of-code/pkg/test"
	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed input.txt
var input string

//go:embed test*.txt
var tests embed.FS

func main() {
	t := test.Register(tests, solve)
	t.Expect(1, 288)
	fmt.Println(solve(input))
}

func solve(input string) int {
	s := utils.ParseInput(input)

	in := [2][]int{
		strings.Ints(s[0]),
		strings.Ints(s[1]),
	}

	total := 1
	for i := range in[0] {
		time := in[0][i]
		distance := in[1][i]

		perms := 0
		for j := 0; j <= time; j++ {
			movedDistance := j * (time - j)
			if movedDistance > distance {
				perms++
			}
		}
		total *= perms
	}

	return total
}
