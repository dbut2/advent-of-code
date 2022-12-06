package main

import (
	_ "embed"
	"fmt"
	"strings"
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
	strs := strings.Split(s, "\n")
	return solve(strs)
}

func solve(s []string) int {
	str := s[0]

	for i := 0; i < len(str)-4; i++ {
		substr := strings.Split(str, "")[i : i+4]
		has := false
		for _, char := range substr {
			if timesInString(substr, char) > 1 {
				has = true
			}
		}
		if !has {
			return i + 4
		}
	}

	return -1
}

func timesInString(s []string, c string) int {
	count := 0
	for _, str := range s {
		if str == c {
			count++
		}
	}
	return count
}

func duplicateLetters(s string) []string {
	letters := make(map[string]bool)
	duplicates := make(map[string]bool)

	for _, char := range strings.Split(s, "") {
		if _, ok := letters[char]; ok {
			duplicates[char] = true
		} else {
			letters[char] = true
		}
	}

	l := []string{}

	for d := range duplicates {
		l = append(l, d)
	}

	return l
}

func stringHasChar(s []string, c string) bool {
	for _, str := range s {
		if str == c {
			return true
		}
	}
	return false
}
