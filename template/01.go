package main

import (
	"embed"
	"fmt"

	"github.com/dbut2/advent-of-code/pkg/test"
	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed input.txt
var input string

//go:embed test*.txt
var tests embed.FS

func main() {
	t := test.Register(tests, solve)
	t.Expect(1, 0)
	fmt.Println(solve(input))
}

func solve(input string) int {
	s := utils.ParseInput(input)

}
