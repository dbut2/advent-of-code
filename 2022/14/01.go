package main

import (
	"embed"
	_ "embed"
	"fmt"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/lists"
	"github.com/dbut2/advent-of-code/pkg/math"
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
	t.Expect(1, 24)
	fmt.Println(solve(input))
}

func solve(input string) int {
	s := utils.ParseInput(input)

	grid := make(Grid)

	for _, str := range s {
		pairs := strings.Split(str, " -> ")
		last := parsePair(pairs[0])
		for i := 1; i < len(pairs); i++ {
			this := parsePair(pairs[i])
			coords := getCoords(last[0], last[1], this[0], this[1])
			for _, coord := range coords {
				grid.Set(coord[0], coord[1], Wall)
			}
			last = this
		}
	}

	maxx := 0
	for x := range grid {
		maxx = math.Max(maxx, x)
	}

	i := -1
	cont := true
	for cont {
		sand := [2]int{0, 500}
		for {
			x := sand[0]
			y := sand[1]

			// check in grid
			if x > maxx {
				cont = false
				break
			}

			// try move down
			if grid.Get(x+1, y) == Empty {
				sand[0]++
				continue
			}

			// try move down left
			if grid.Get(x+1, y-1) == Empty {
				sand[0]++
				sand[1]--
				continue
			}

			// try move down right
			if grid.Get(x+1, y+1) == Empty {
				sand[0]++
				sand[1]++
				continue
			}

			break
		}

		grid.Set(sand[0], sand[1], Sand)
		i++
	}

	return i
}

func parsePair(s string) [2]int {
	pair := sti.Stis(strings.Split(s, ","))
	return [2]int{pair[1], pair[0]}

}

func getCoords(x1, y1, x2, y2 int) [][2]int {
	// horizontal
	if x1 == x2 {
		var items [][2]int
		for _, y := range lists.Range(y1, y2) {
			items = append(items, [2]int{x1, y})
		}
		return items
	}

	//vertical
	if y1 == y2 {
		var items [][2]int
		for _, x := range lists.Range(x1, x2) {
			items = append(items, [2]int{x, y1})
		}
		return items
	}

	panic("not a straight line")
}

type Grid map[int]map[int]Cell

func (g *Grid) Set(x, y int, c Cell) {
	if g == nil {
		*g = make(Grid)
	}
	if _, ok := (*g)[x]; !ok {
		(*g)[x] = map[int]Cell{}
	}
	(*g)[x][y] = c
}

func (g *Grid) Get(x, y int) Cell {
	if _, ok := (*g)[x]; !ok {
		return Empty
	}
	return (*g)[x][y]
}

type Cell int

const (
	Empty Cell = iota
	Wall
	Sand
)

func stepSand(g Grid) {

}
