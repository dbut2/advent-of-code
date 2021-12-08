package main

import (
	_ "embed"
	"fmt"
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
	var ssds []*ssd

	ofsoe := 0

	for _, line := range s {
		ss := &ssd{}

		numbers := strings.Split(line, " ")

		observation := true

		for _, n := range numbers {
			if observation {
				ss.observations = append(ss.observations, n)
			} else {
				//           1              7              4              8
				if len(n) == 2 || len(n) == 3 || len(n) == 4 || len(n) == 7 {
					ofsoe++
				}
				ss.outputs = append(ss.outputs, n)
			}

			if n == "|" {
				observation = false
			}
		}

		ssds = append(ssds, ss)
	}

	return ofsoe
}

type ssd struct {
	observations []string
	outputs      []string

	lights []bool

	// index points to correct value in .lights
	correction []int
}

func (s *ssd) correct(test []string) {

	for i, t := range test {
		_ = i
		if len(t) == 2 {

		}
	}

}
