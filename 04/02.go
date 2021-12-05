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
	numbers := strings.Split(s[0], ",")

	called := stringsToInts(numbers)

	boards := []Board{}

	for i := 2; i < len(s); i += 6 {
		b := make(Board)
		for j := 0; j < 5; j++ {
			b[j] = make(map[int]*Cell)
			line := stringsToInts(strings.Split(s[j+i], " "))
			for k := 0; k < 5; k++ {
				b[j][k] = &Cell{
					Number: line[k],
					Marked: false,
				}
			}
		}
		boards = append(boards, b)
	}

	solved := map[int]int{}
	states := map[int]int{}

	for i, n := range called {
		for j, b := range boards {
			b.Mark(n)
			if b.IsComplete() {
				if solved[j] == 0 {
					solved[j] = i
					states[j] = b.SumOfRemainder()
				}
			}
		}
	}

	min := FindMax(solved)

	return states[min] * called[solved[min]]
}

type Board map[int]map[int]*Cell

func (b Board) IsComplete() bool {
	for i := 0; i < 5; i++ {
		complete := true
		for j := 0; j < 5; j++ {
			if b[i][j].Marked == false {
				complete = false
			}
		}
		if complete == true {
			return true
		}
	}

	for j := 0; j < 5; j++ {
		complete := true
		for i := 0; i < 5; i++ {
			if b[i][j].Marked == false {
				complete = false
			}
		}
		if complete == true {
			return true
		}
	}

	return false
}

func (b Board) SumOfRemainder() int {
	rem := 0

	for _, i := range b {
		for _, j := range i {
			if !j.Marked {
				rem += j.Number
			}
		}
	}

	return rem
}

func (b Board) Mark(n int) {
	for _, i := range b {
		for _, j := range i {
			if j.Number == n {
				j.Marked = true
			}
		}
	}
}

type Cell struct {
	Number int
	Marked bool
}

func stringsToInts(s []string) []int {
	ints := []int{}
	for _, str := range s {
		if str == "" {
			continue
		}
		i, err := strconv.Atoi(str)
		if err != nil {
			panic(err.Error())
		}
		ints = append(ints, i)
	}
	return ints
}

func FindMax(m map[int]int) int {
	max := -1
	index := -1

	for i, n := range m {
		if n > max {
			max = n
			index = i
		}
	}

	return index
}
