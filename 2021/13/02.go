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
	empty := [][]int{}
	for i := 0; i < max; i++ {
		empty = append(empty, make([]int, max))
	}

	grid := empty

	dimensions := [2]int{}

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
			newGrid := empty

			splitLine := strings.Split(line, " ")

			foldAt := strings.Split(splitLine[2], "=")

			along := foldAt[0]
			at, err := strconv.Atoi(foldAt[1])
			if err != nil {
				panic(err.Error())
			}

			switch along {
			case "x":
				dimensions[0] = at + 1
			case "y":
				dimensions[1] = at + 1
			}

			for i := range grid {
				for j := range grid[i] {
					switch along {
					case "x":
						if i < at {
							newGrid[i][j] += grid[i][j]
						}
						if i > at {
							if 2*at-i >= 0 {
								newGrid[2*at-i][j] += grid[i][j]
							}
						}
					case "y":
						if j < at {
							newGrid[i][j] += grid[i][j]
						}
						if j > at {
							if 2*at-j >= 0 {
								newGrid[i][2*at-j] += grid[i][j]
							}
						}
					}
				}
			}

			grid = newGrid
		}
	}

	for j := 0; j < dimensions[1]; j++ {
		for i := 0; i < dimensions[0]; i++ {
			if grid[i][j] > 0 {
				fmt.Printf("▓")
			} else {
				fmt.Printf("░")
			}
		}
		fmt.Printf("\n")
	}

	return 0
}
