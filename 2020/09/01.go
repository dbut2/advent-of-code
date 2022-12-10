package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/sti"
	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed input.txt
var input string

//go:embed test1.txt
var test string

func main() {
	utils.Test(solve(test, 5), 127)
	fmt.Println(solve(input, 25))
}

func solve(input string, prev int) int {
	s := strings.Split(input, "\n")

	ints := sti.Stis(s)

	for i := prev; i < len(ints); i++ {
		hasPair := false
		for x := i - prev; x < i; x++ {
			for y := x + 1; y < i; y++ {
				if ints[x]+ints[y] == ints[i] {
					hasPair = true
				}
			}
		}

		if !hasPair {
			return ints[i]
		}
	}

	return -1
}
