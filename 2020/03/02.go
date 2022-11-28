package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

var s = strings.Split(input, "\n")

func main() {
	s := strings.Split(input, "\n")
	i := solve(s)
	fmt.Println(i)
}

func solve(s []string) int {
	return treesFor(1, 1) * treesFor(3, 1) * treesFor(5, 1) * treesFor(7, 1) * treesFor(1, 2)
}

func treesFor(xi, yi int) int {
	trees := 0
	x, y := 0, 0
	for ; y < len(s); y += yi {
		if isTree(x, y) {
			trees++
		}
		x += xi
	}
	return trees
}

func isTree(x, y int) bool {
	x = x % len(s[0])
	return strings.Split(s[y], "")[x] == "#"
}
