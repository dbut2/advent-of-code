package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed input.txt
var input string

//go:embed test1.txt
var test1 string

//go:embed test2.txt
var test2 string

func main() {
	utils.Test(solve(test1), 0)
	fmt.Println(solve(input))
}

func solve(input string) int {
	s := strings.Split(input, "\n")

}
