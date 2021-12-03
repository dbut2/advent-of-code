package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed 2.txt
var measurements string

func main() {
	m := strings.Split(measurements, "\n")

	ints := strsToInts(m)

	c := 0
	for i, v := range ints {
		if i < 3 {
			continue
		}

		if v > ints[i-3] {
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
