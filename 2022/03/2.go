package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/lists"
	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed input.txt
var input string

//go:embed test1.txt
var test string

func main() {
	utils.Test(solve(test), 70)
	fmt.Println(solve(input))
}

func solve(input string) int {
	s := utils.ParseInput(input)

	count := 0

	for i := 0; i < len(s); i += 3 {

		l1 := strings.Split(s[i], "")
		l2 := strings.Split(s[i+1], "")
		l3 := strings.Split(s[i+2], "")

		inboth := lists.Intersection(lists.Intersection(l1, l2), l3)[0]

		count += priority(inboth)
	}

	return count
}

func priority(a string) int {
	b := []byte(a)[0]

	if b > 96 {
		b -= 58
	}

	b -= 38

	return int(b)
}
