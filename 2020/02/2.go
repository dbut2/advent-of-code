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
		first, err := strconv.Atoi(split2[0])
		if err != nil {
			panic(err.Error())
		}
		second, err := strconv.Atoi(split2[1])
		if err != nil {
			panic(err.Error())
		}

		indexesFound := 0

		if breakString(password)[first-1] == letter {
			indexesFound++
		}

		if breakString(password)[second-1] == letter {
			indexesFound++
		}

		if indexesFound == 1 {
			passed++
		}
	}

	return passed
}

func breakString(s string) []string {
	return strings.Split(s, "")
}
