package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	s := strings.Split(input, "\n\n")
	i := solve(s)
	fmt.Println(i)
}

func solve(s []string) int {
	count := 0

	for _, str := range s {

		answers := strings.Split(str, "\n")

		gcount := len(lettersInAll(answers))

		count += gcount
	}

	return count
}

func lettersInAll(strs []string) string {
	letters := strs[0]
	for _, str := range strs {
		letters = union(letters, str)
	}
	return letters
}

func union(a, b string) string {
	u := ""
	for _, l := range strings.Split(a, "") {
		isinb := false
		for _, m := range strings.Split(b, "") {
			if l == m {
				isinb = true
			}
		}
		if isinb {
			u += l
		}
	}
	return u
}
