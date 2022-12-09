package main

import (
	_ "embed"
	"fmt"

	"github.com/dbut2/advent-of-code/pkg/lists"
	"github.com/dbut2/advent-of-code/pkg/sti"
	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed input.txt
var input string

//go:embed test1.txt
var test string

func main() {
	utils.Test(solve(test), 21)
	fmt.Println(solve(input))
}

func solve(input string) int {
	s := utils.ParseInput(input)
	arr := sti.Stiss(s)

	visible := lists.Fill2D(len(arr), len(arr[0]), false)

	for i := 0; i < len(arr); i++ {
		// Scan each column downwards and check each tree against the biggest we've seen so far, if the tree is the
		// biggest for the column at that point we know it's visible from outside
		max := -1
		for j := 0; j < len(arr[0]); j++ {
			if arr[i][j] > max {
				max = arr[i][j]
				visible[i][j] = true
			}
		}
	}

	for i := 0; i < len(arr); i++ {
		// Scan each column updwards
		max := -1
		for j := len(arr[0]) - 1; j >= 0; j-- {
			if arr[i][j] > max {
				max = arr[i][j]
				visible[i][j] = true
			}
		}
	}

	for j := 0; j < len(arr[0]); j++ {
		// Scan each row to the right
		max := -1
		for i := 0; i < len(arr); i++ {
			if arr[i][j] > max {
				max = arr[i][j]
				visible[i][j] = true
			}
		}
	}

	for j := 0; j < len(arr[0]); j++ {
		// Scan each row to the left
		max := -1
		for i := len(arr) - 1; i >= 0; i-- {
			if arr[i][j] > max {
				max = arr[i][j]
				visible[i][j] = true
			}
		}
	}

	count := 0
	for _, a := range visible {
		for _, b := range a {
			if b {
				count++
			}
		}
	}

	return count
}
