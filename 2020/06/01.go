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
		gcount := 0
		yesses := map[string]bool{}

		answers := strings.Split(str, "\n")

		for _, answer := range answers {

			for _, question := range strings.Split(answer, "") {
				if _, ok := yesses[question]; !ok {
					gcount++
					yesses[question] = true
				}
			}

		}

		count += gcount
	}

	return count
}
