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
	var i []int
	for _, str := range s {
		i = append(i, Sti(str))
	}
	return i
}

func Stiss(s []string) [][]int {
	var i [][]int
	for _, line := range s {
		i = append(i, Stis(strings.Split(line, "")))
	}
	return i
}
