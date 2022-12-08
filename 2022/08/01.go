package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/dbut2/advent-of-code/utils"
)

//go:embed input.txt
var input string

//go:embed test.txt
var test string

func main() {
	fmt.Println("Test")
	fmt.Println(do(test))
	fmt.Println()
	fmt.Println("Solution")
	fmt.Println(do(input))
}

func do(s string) int {
	strs := strings.Split(s, "\n")
	return solve(strs)
}

func solve(s []string) int {
	arr := utils.Stiss(s)

	visible := utils.Fill2D(len(arr), len(arr[0]), false)

	for i := 0; i < len(arr); i++ {
		max := -1
		for j := 0; j < len(arr[0]); j++ {
			if arr[i][j] > max {
				max = arr[i][j]
				visible[i][j] = true
			}
		}
	}

	for i := 0; i < len(arr); i++ {
		max := -1
		for j := len(arr[0]) - 1; j >= 0; j-- {
			if arr[i][j] > max {
				max = arr[i][j]
				visible[i][j] = true
			}
		}
	}

	for j := 0; j < len(arr[0]); j++ {
		max := -1
		for i := 0; i < len(arr); i++ {
			if arr[i][j] > max {
				max = arr[i][j]
				visible[i][j] = true
			}
		}
	}

	for j := 0; j < len(arr[0]); j++ {
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
