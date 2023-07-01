package sti

import (
	"strconv"
	"strings"
)

func Sti(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err.Error())
	}
	return i
}

func Stis(s []string) []int {
	ints := make([]int, len(s))
	for i, str := range s {
		ints[i] = Sti(str)
	}
	return ints
}

func Stiss(s []string, sep string) [][]int {
	ints := make([][]int, len(s))
	for i, line := range s {
		ints[i] = Stis(strings.Split(line, sep))
	}
	return ints
}
