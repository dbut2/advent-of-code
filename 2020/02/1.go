package main

import (
	_ "embed"
	"fmt"
	"strconv"
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
	passed := 0

	for _, str := range s {
		split1 := strings.Split(str, " ")
		rule := split1[0]
		letter := strings.Trim(split1[1], ":")
		password := split1[2]

		split2 := strings.Split(rule, "-")
		min, err := strconv.Atoi(split2[0])
		if err != nil {
			panic(err.Error())
		}
		max, err := strconv.Atoi(split2[1])
		if err != nil {
			panic(err.Error())
		}

		counts := letterCounts(password)

		if counts[letter] >= min && counts[letter] <= max {
			passed++
		}
	}

	return passed
}

func letterCounts(s string) map[string]int {
	counts := map[string]int{}

	letters := strings.Split(s, "")
	for _, letter := range letters {
		if _, ok := counts[letter]; !ok {
			counts[letter] = 0
		}

		counts[letter]++
	}

	return counts
}
