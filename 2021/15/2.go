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
	grid0 := [][]int{}
	grid := [][]int{}
	mins := [][]int{}
	for i := 0; i < len(s); i++ {
		grid0 = append(grid0, make([]int, len(s[0])))
	}
	for i := 0; i < len(s)*5; i++ {
		grid = append(grid, make([]int, len(s[0])*5))
		mins = append(mins, make([]int, len(s[0])*5))
	}

	height := len(grid0)
	width := len(grid0[0])

	for i, line := range s {
		for j, cell := range strings.Split(line, "") {
			v, err := strconv.Atoi(cell)
			if err != nil {
				panic(err.Error())
			}
			grid0[i][j] = v
		}
	}

	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[0]); y++ {

			i := x % height
			j := y % width

			row := (x - i) / height
			col := (y - j) / width

			grid[x][y] = (grid0[i][j]+row+col-1)%9 + 1
		}
	}

	height = height*5 - 1
	width = width*5 - 1

	mins[height][width] = grid[height][width]
	heap := [][2]int{{height - 1, width}, {height, width - 1}}

	for len(heap) > 0 {
		var check [2]int
		heap, check = heap[:len(heap)-1], heap[len(heap)-1]

		i, j := check[0], check[1]

		prev := mins[i][j]

		var neighbours [][2]int
		if i > 0 {
			neighbours = append(neighbours, [2]int{i - 1, j})
		}
		if j > 0 {
			neighbours = append(neighbours, [2]int{i, j - 1})
		}
		if i < height {
			neighbours = append(neighbours, [2]int{i + 1, j})
		}
		if j < width {
			neighbours = append(neighbours, [2]int{i, j + 1})
		}

		for _, n := range neighbours {
			x, y := n[0], n[1]
			v := mins[x][y]

			if v != 0 && (v+grid[i][j] < prev || prev == 0) {
				mins[i][j] = v + grid[i][j]

				heap = append(heap, neighbours...)
			}
		}
	}

	return mins[0][0] - grid[0][0]
}
