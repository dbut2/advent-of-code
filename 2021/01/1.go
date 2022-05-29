package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed 1.txt
var measurements1 string

func main() {
	m := strings.Split(measurements1, "\n")

	ints := strsToInts(m)

	c := 0
	for i, v := range ints {
		if i == 0 {
			continue
		}

		if v > ints[i - 1] {
			c++
		}
	}

	fmt.Println(c)
}

func strsToInts(s []string) []int {
	ints := []int{}
	for _, str := range s {
		i, err := strconv.Atoi(str)
		if err != nil {
			panic(err.Error())
		}
		ints = append(ints, i)
	}
	return ints
}
