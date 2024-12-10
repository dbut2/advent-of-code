package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

var max = 1500

func main() {
	s := strings.Split(input, "\n")
	i := solve(s)
	fmt.Println(i)
}

func solve(s []string) int {
	grid := [][]int{}
	for i := 0; i < max; i++ {
		grid = append(grid, make([]int, max))
	}

	firstFold := true
	for _, line := range s {
		if strings.Contains(line, ",") {
			coords := strings.Split(line, ",")

			x, err := strconv.Atoi(coords[0])
			if err != nil {
				panic(err.Error())
			}

			y, err := strconv.Atoi(coords[1])
			if err != nil {
				panic(err.Error())
			}

			grid[x][y] = 1
		}

		if strings.Contains(line, "=") {
			if firstFold {
				splitLine := strings.Split(line, " ")

				foldAt := strings.Split(splitLine[2], "=")

				along := foldAt[0]
				at, err := strconv.Atoi(foldAt[1])
				if err != nil {
					panic(err.Error())
				}

				count := 0

				for i := range grid {
					for j := range grid[i] {
						switch along {
						case "x":
							if i > at {
								if grid[i][j] > 0 {
									count++
								}
							}
							if i < at {
								if grid[2*at-i][j] == 0 {
									if grid[i][j] > 0 {
										count++
									}
								}
							}
						case "y":
							if j > at {
								if grid[i][j] > 0 {
									count++
								}
							}
							if j < at {
								if grid[i][2*at-j] == 0 {
									if grid[i][j] > 0 {
										count++
									}
								}
							}
						}
					}
				}

				return count
			}

			firstFold = false
		}
	}

	return 0
}

func fold(grid [][]int, along string, at int) [][]int {
	a := clear(grid, false, along, at)
	b := clear(grid, true, along, at)

	if at*2 < max {
		temp := a
		a = b
		b = temp
	}

	b = flip(b, along, at)

	c := add(a, b)

	return c
}

func add(a, b [][]int) [][]int {
	if len(a) != len(b) {
		panic("a and b not same length")
	}

	if len(a[0]) != len(b[0]) {
		panic("a and b not same length")
	}

	for i := range a {
		for j := range a[i] {
			a[i][j] += b[i][j]
		}
	}

	return a
}

func clear(grid [][]int, after bool, along string, at int) [][]int {
	switch along {
	case "x":
		for i := range grid {
			for j := range grid[i] {
				if after {
					if j >= at {
						grid[i][j] = 0
					}
				} else {
					if j <= at {
						grid[i][j] = 0
					}
				}
			}
		}
	case "y":
		for i := range grid {
			for j := range grid[i] {
				if after {
					if i >= at {
						grid[i][j] = 0
					}
				} else {
					if i <= at {
						grid[i][j] = 0
					}
				}
			}
		}
	}
	return grid
}

func flip(grid [][]int, along string, at int) [][]int {
	newGrid := [][]int{}
	for i := 0; i < max; i++ {
		newGrid = append(newGrid, make([]int, max))
	}

	for i := range newGrid {
		for j := range newGrid[i] {
			switch along {
			case "x":
				if max-j+2*at >= 0 && max-j+2*at < max {
					newGrid[i][max-j+2*at] = grid[i][j]
				}
			case "y":
				if max-i+2*at >= 0 && max-i+2*at < max {
					newGrid[max-i+2*at][j] = grid[i][j]
				}
			}
		}
	}

	return newGrid
}
