package main

import (
	_ "embed"
	"fmt"
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

	c := 0

	for _, line := range s {
		c += score(line)
	}

	return c
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
				return 25137
			}
		case "}":
			var p string
			p, stack = stack[len(stack)-1], stack[:len(stack)-1]
			if p != "{" {
				return 1197
			}
		case "]":
			var p string
			p, stack = stack[len(stack)-1], stack[:len(stack)-1]
			if p != "[" {
				return 57
			}
		case ")":
			var p string
			p, stack = stack[len(stack)-1], stack[:len(stack)-1]
			if p != "(" {
				return 3
			}
		}
	}
	return 0
}
