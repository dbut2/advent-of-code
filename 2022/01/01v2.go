package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/math"
	"github.com/dbut2/advent-of-code/pkg/sti"
	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed input.txt
var input string

//go:embed test1.txt
var test string

func main() {
	utils.Test(solve(test), 24000)
	fmt.Println(solve(input))
}

func solve(input string) int {
	s := strings.Split(input, "\n\n")
	m := 0
	for _, str := range s {
		m = math.Max(m, math.Sum(sti.Stis(strings.Split(str, "\n"))))
	}
	return m
}
