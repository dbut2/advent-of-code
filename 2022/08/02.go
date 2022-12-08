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
	max := 0
	for i, row := range arr {
		for j, tree := range row {

			up := 0
			for x := i - 1; x >= 0; x-- {
				up++
				if arr[x][j] >= tree {
					break
				}
			}
			right := 0
			for x := j + 1; x < len(row); x++ {
				right++
				if arr[i][x] >= tree {
					break
				}
			}
			down := 0
			for x := i + 1; x < len(arr); x++ {
				down++
				if arr[x][j] >= tree {
					break
				}
			}
			left := 0
			for x := j - 1; x >= 0; x-- {
				left++
				if arr[i][x] >= tree {
					break
				}
			}
			val := up * right * down * left
			max = utils.Max(max, val)
		}
	}
	return max
}
