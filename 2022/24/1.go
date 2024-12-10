package main

import (
	"embed"
	_ "embed"
	"fmt"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/grid"
	"github.com/dbut2/advent-of-code/pkg/sets"
	"github.com/dbut2/advent-of-code/pkg/test"
	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed input.txt
var input string

//go:embed test*.txt
var tests embed.FS

func main() {
	t := test.Register(tests, solve)
	t.Expect(1, 18)
	fmt.Println(solve(input))
}

func solve(input string) int {
	s := utils.ParseInput(input)

	cols, rows := len(s[0]), len(s)
	ec, er := cols-2, rows-2

	g := grid.Grid[string]{}

	for y, line := range s {
		for x, str := range strings.Split(line, "") {
			g.Set(x, y, str)
		}
	}

	lm := sets.Set[[2]int]{}
	lm.Add([2]int{1, 0})

	for i := 1; ; i++ {
		npm := sets.Set[[2]int]{}
		npm.Add([2]int{1, 0})

		for m := range lm {
			x, y := m[0], m[1]

			for _, move := range [][2]int{{0, 0}, {0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
				nx, ny := x+move[0], y+move[1]

				if nx == cols-2 && ny == rows-1 {
					return i
				}

				if nx <= 0 || nx >= cols-1 {
					continue
				}

				if ny <= 0 || ny >= rows-1 {
					continue
				}

				if g.Get(nx, ny) == "#" && g.Get(nx, ny) != "" {
					continue
				}

				if g.Get((nx-1-i+ec*(i/ec+1))%ec+1, (ny-1+er*(i/er+1))%er+1) == ">" {
					continue
				}

				if g.Get((nx-1+ec*(i/ec+1))%ec+1, (ny-1-i+er*(i/er+1))%er+1) == "v" {
					continue
				}

				if g.Get((nx-1+i+ec*(i/ec+1))%ec+1, (ny-1+er*(i/er+1))%er+1) == "<" {
					continue
				}

				if g.Get((nx-1+ec*(i/ec+1))%ec+1, (ny-1+i+er*(i/er+1))%er+1) == "^" {
					continue
				}

				npm.Add([2]int{nx, ny})
			}
		}

		lm = npm
	}
}
