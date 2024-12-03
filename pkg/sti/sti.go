package sti

import (
	"strconv"
	"strings"
)

func Int(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err.Error())
	}
	return i
}

func Ints(s []string) []int {
	ints := make([]int, len(s))
	for i, str := range s {
		ints[i] = Int(str)
	}
	return ints
}

func Intss(s []string, sep string) [][]int {
	ints := make([][]int, len(s))
	for i, line := range s {
		ints[i] = Ints(strings.Split(line, sep))
	}
	return ints
}
