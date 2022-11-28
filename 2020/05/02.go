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

func solve(str []string) int {
	found := map[int]bool{}

	highest := 0

	for _, s := range str {
		seat := seatBin(s)
		found[seat] = true
		if seat > highest {
			highest = seat
		}
	}

	for i := highest - 1; i > 0; i-- {
		if _, ok := found[i]; !ok {
			return i
		}
	}

	return 0
}

func seatBin(s string) int {
	s = strings.ReplaceAll(s, "B", "1")
	s = strings.ReplaceAll(s, "F", "0")
	s = strings.ReplaceAll(s, "R", "1")
	s = strings.ReplaceAll(s, "L", "0")

	i, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		return 0
	}
	return int(i)
}
