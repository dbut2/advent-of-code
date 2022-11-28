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
	highest := 0

	for _, s := range str {
		seat := seatBin(s)
		if seat > highest {
			highest = seat
		}
	}

	return highest
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
