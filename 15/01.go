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
	grid := [][]int{}
	mins := [][]int{}
	for i := 0; i < len(s); i++ {
		grid = append(grid, make([]int, len(s[0])))
		mins = append(mins, make([]int, len(s[0])))
	}

	for i, line := range s {
		for j, cell := range strings.Split(line, "") {
			v, err := strconv.Atoi(cell)
			if err != nil {
				panic(err.Error())
			}
			grid[i][j] = v
		}
	}

	height := len(grid) - 1
	width := len(grid[0]) - 1

	for x := height + width; x >= 0; x-- {
		for i, line := range grid {
			for j, cell := range line {
				if i+j == x {
					if i == height && j == width {
						mins[i][j] = cell
						continue
					}

					if i == height {
						mins[i][j] = mins[i][j+1] + cell
						continue
					}
					if j == width {
						mins[i][j] = mins[i+1][j] + cell
						continue
					}

					mins[i][j] = min(mins[i+1][j], mins[i][j+1]) + cell
				}
			}
		}
	}

	return mins[0][0] - grid[0][0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
