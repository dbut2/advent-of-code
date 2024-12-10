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
var test string

func main() {
	utils.Test(solve(test), 19)
	fmt.Println(solve(input))
}

func solve(input string) int {
	s := utils.ParseInput(input)
	str := s[0]

	for i := 0; i < len(str)-4; i++ {
		substr := strings.Split(str, "")[i : i+14]
		has := false
		for _, char := range substr {
			if timesInString(substr, char) > 1 {
				has = true
			}
		}
		if !has {
			return i + 14
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
