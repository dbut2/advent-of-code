package main

import (
	"embed"
	_ "embed"
	"fmt"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/sets"
	"github.com/dbut2/advent-of-code/pkg/sti"
	"github.com/dbut2/advent-of-code/pkg/test"
	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed input.txt
var input string

//go:embed test*.txt
var tests embed.FS

func main() {
	t := test.Register(tests, solve)
	t.Expect(1, 64)
	fmt.Println(solve(input))
}

func solve(input string) int {
	s := utils.ParseInput(input)

	cubes := sets.Set[[3]int]{}

	for _, str := range s {
		c := sti.Stis(strings.Split(str, ","))
		cubes.Add([3]int{c[0], c[1], c[2]})
	}

	total := 0

	for _, c := range cubes.Slice() {
		sides := [6][3]int{
			{1, 0, 0},
			{0, 1, 0},
			{0, 0, 1},
			{-1, 0, 0},
			{0, -1, 0},
			{0, 0, -1},
		}

		count := 6
		for _, d := range sides {
			x, y, z := c[0]+d[0], c[1]+d[1], c[2]+d[2]
			if cubes.Has([3]int{x, y, z}) {
				count--
			}
		}
		total += count
	}

	return total
}
