package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input string

//go:embed test.txt
var test string

func main() {
	fmt.Println("Test")
	fmt.Println(do(test))
	fmt.Println()
	fmt.Println("Solution")
	fmt.Println(do(input))
}

func do(s string) int {
	return solve(s)
}

func solve(s string) int {
	l := 4
	for i := 0; i < len(s)-l; i++ {
		m := make(map[rune]bool)
		for _, char := range s[i : i+l] {
			m[char] = true
		}
		if len(m) == l {
			return i + l
		}
	}
	return -1
}
