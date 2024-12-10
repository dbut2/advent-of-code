package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	s := strings.Split(input, "\n")
	i := solve(s)
	fmt.Println(i)
}

func solve(s []string) int {
	h := heightMap{}

	for _, line := range s {
		row := []int{}
		for _, v := range strings.Split(line, "") {
			n, err := strconv.Atoi(v)
			if err != nil {
				panic(err.Error())
			}
			row = append(row, n)
		}
		h = append(h, row)
	}

	cumRisk := 0

	for y, row := range h {
		for x, c := range row {
			u := h.get(x, y-1)
			d := h.get(x, y+1)
			l := h.get(x-1, y)
			r := h.get(x+1, y)

			if c < u && c < d && c < l && c < r {
				cumRisk += c + 1
			}
		}
	}

	return cumRisk
}

type heightMap [][]int

func (h heightMap) get(x, y int) int {
	if y < 0 || y >= len(h) {
		return 10
	}
	row := h[y]
	if x < 0 || x >= len(row) {
		return 10
	}
	return row[x]
}
