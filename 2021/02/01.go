package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed 01.txt
var input string

func main() {
	lines := strings.Split(input, "\n")
	inputs := inputToInput(lines)

	h := 0
	d := 0

	for _, i := range inputs {
		switch i.d {
		case forward:
			h += i.n
		case up:
			d -= i.n
		case down:
			d += i.n
		}
	}

	fmt.Println(h * d)
}

type Direction int

const (
	forward Direction = iota
	up
	down
)

type Input struct {
	d Direction
	n int
}

func inputToInput(s []string) []Input {
	inputs := []Input{}

	for _, v := range s {
		split := strings.Split(v, " ")

		a, err := strconv.Atoi(split[1])
		if err != nil {
			panic(err.Error())
		}

		i := Input{
			d: strToDirectio(split[0]),
			n: a,
		}

		inputs = append(inputs, i)
	}

	return inputs
}

func strToDirectio(s string) Direction {
	switch s {
	case "forward":
		return forward
	case "up":
		return up
	case "down":
		return down
	}
	return 0
}
