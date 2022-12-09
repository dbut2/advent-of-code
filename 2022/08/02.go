package main

import (
	_ "embed"
	"fmt"

	"github.com/dbut2/advent-of-code/pkg/math"
	"github.com/dbut2/advent-of-code/pkg/sti"
	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed input.txt
var input string

//go:embed test1.txt
var test string

func main() {
	utils.Test(solve(test), 8)
	fmt.Println(solve(input))
}

func solve(input string) int {
	s := utils.ParseInput(input)

	arr := sti.Stiss(s)
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
			max = math.Max(max, val)
		}
	}
	return max
}
