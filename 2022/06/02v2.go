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
	utils.Test(solve(test), 19)
	fmt.Println(solve(input))
}

func solve(input string) int {
	for i := 0; i < len(input)-14; i++ {
		m := make(map[rune]bool)
		for _, char := range input[i : i+14] {
			m[char] = true
		}
		if len(m) == 14 {
			return i + 14
		}
	}
	return -1
}
