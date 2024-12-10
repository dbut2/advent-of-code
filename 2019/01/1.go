package main

import (
	"embed"
	_ "embed"
	"fmt"

	"github.com/dbut2/advent-of-code/pkg/sti"
	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed input.txt
var input string

//go:embed test*.txt
var tests embed.FS

func main() {
	fmt.Println(solve(input))
}

func solve(input string) int {
	s := utils.ParseInput(input)

	total := 0

	for _, line := range sti.Stis(s) {
		lv := (line / 3) - 2
		total += lv
	}

	return total
}
