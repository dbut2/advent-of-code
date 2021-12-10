package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	s := strings.Split(input, "\n")
	i := solve(s)
	fmt.Println(i)
}

func solve(s []string) int {
	// arrow >
	// brace }
	// bracket ]
	// parentheses )

	var scores []int

	for _, line := range s {
		sc := score(line)
		if sc > 0 {
			scores = append(scores, sc)
		}
	}

	sort.Ints(scores)

	middle := scores[(len(scores)-1)/2]

	return middle
}

func score(s string) int {
	var stack []string

	for _, char := range strings.Split(s, "") {
		switch char {
		case "<", "{", "[", "(":
			stack = append(stack, char)
		case ">":
			var p string
			p, stack = stack[len(stack)-1], stack[:len(stack)-1]
			if p != "<" {
				return 0
			}
		case "}":
			var p string
			p, stack = stack[len(stack)-1], stack[:len(stack)-1]
			if p != "{" {
				return 0
			}
		case "]":
			var p string
			p, stack = stack[len(stack)-1], stack[:len(stack)-1]
			if p != "[" {
				return 0
			}
		case ")":
			var p string
			p, stack = stack[len(stack)-1], stack[:len(stack)-1]
			if p != "(" {
				return 0
			}
		}
	}

	c := 0

	for i := len(stack) - 1; i >= 0; i-- {
		char := stack[i]
		c *= 5
		switch char {
		case "<":
			c += 4
		case "{":
			c += 3
		case "[":
			c += 2
		case "(":
			c += 1
		}
	}

	return c
}
