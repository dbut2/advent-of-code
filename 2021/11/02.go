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

	for _, line := range s {
		row := []int{}
		for _, char := range strings.Split(line, "") {
			v, err := strconv.Atoi(char)
			if err != nil {
				panic(err.Error())
			}
			row = append(row, v)
		}
		grid = append(grid, row)
	}

	for i := 0; true; i++ {
		if allZero(grid) {
			return i
		}
		grid, _ = step(grid)
	}

	return 0
}

func step(grid [][]int) ([][]int, int) {
	flashes := 0

	for i := range grid {
		for j := range grid[i] {
			grid[i][j]++
		}
	}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 0 {
				continue
			}
			if grid[i][j] > 9 {
				var f int
				grid, f = flash(grid, i, j)
				flashes += f
			}
		}
	}

	return grid, flashes
}

func flash(grid [][]int, i, j int) ([][]int, int) {
	flashes := 1

	grid[i][j] = 0

	rows := []int{i}
	cols := []int{j}

	if i > 0 {
		rows = append(rows, i-1)
	}
	if i < len(grid)-1 {
		rows = append(rows, i+1)
	}

	if j > 0 {
		cols = append(cols, j-1)
	}
	if j < len(grid[0])-1 {
		cols = append(cols, j+1)
	}

	for _, i2 := range rows {
		for _, j2 := range cols {
			if grid[i2][j2] == 0 {
				continue
			}
			grid[i2][j2]++
			if grid[i2][j2] > 9 {
				var f int
				grid, f = flash(grid, i2, j2)
				flashes += f
			}
		}
	}

	return grid, flashes
}

func allZero(grid [][]int) bool {
	all := true

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] != 0 {
				all = false
			}
		}
	}

	return all
}
