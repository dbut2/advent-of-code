package main

import (
	_ "embed"
	"fmt"

	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed input.txt
var input string

//go:embed test1.txt
var test string

func main() {
	utils.Test(solve(test), 7)
	fmt.Println(solve(input))
}

func solve(input string) int {
	for i := 0; i < len(input)-4; i++ {
		m := make(map[rune]bool)
		for _, char := range input[i : i+4] {
			m[char] = true
		}
		if len(m) == 4 {
			return i + 4
		}
	}
	return -1
}
