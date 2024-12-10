package main

import (
	"embed"
	_ "embed"
	"fmt"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/sti"
	"github.com/dbut2/advent-of-code/pkg/test"
)

//go:embed input.txt
var input string

//go:embed test*.txt
var tests embed.FS

func main() {
	t := test.Register(tests, solve)
	t.Expect(1, 20899048083289)
	fmt.Println(solve(input))
}

func solve(input string) int {
	tiles := strings.Split(input, "\n\n")

	p := make(map[int][]string)
	m := make(map[int]int)

	for _, tile := range tiles {
		t := strings.Split(tile, "\n")
		p[sti.Sti(t[0][5:9])] = getPossibleEdges(t[1:])
	}

	for tile, ps := range p {
		m[tile] = -8
		for _, pse := range ps {
			for _, ps2 := range p {
				for _, pse2 := range ps2 {
					if pse == pse2 {
						m[tile]++
					}
				}
			}
		}
	}

	total := 1
	for tile, c := range m {
		if c == 4 {
			total *= tile
		}
	}

	return total
}

func getPossibleEdges(tile []string) []string {
	var p []string

	p = append(p, tile[0])
	p = append(p, tile[9])

	left := ""
	right := ""
	for _, row := range tile {
		left += row[:1]
		right += row[9:]
	}
	p = append(p, left)
	p = append(p, right)

	for _, po := range p {
		b := ""
		for _, char := range strings.Split(po, "") {
			b = char + b
		}
		p = append(p, b)
	}

	return p
}
