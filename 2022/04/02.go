package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"

	"github.com/dbut2/advent-of-code/utils"
)

//go:embed input.txt
var input string

func main() {
	start := time.Now()
	s := strings.Split(input, "\n")
	i := solve(s)
	fmt.Println(i)
	fmt.Println(time.Now().Sub(start))
}

func solve(s []string) int {
	overlap := 0

	for _, str := range s {
		elves := strings.Split(str, ",")

		e1 := utils.Stis(strings.Split(elves[0], "-"))
		e2 := utils.Stis(strings.Split(elves[1], "-"))

		e1s, e1e := e1[0], e1[1]
		e2s, e2e := e2[0], e2[1]

		if e1s <= e2s && e2s <= e1e {
			overlap++
			continue
		}

		if e2s <= e1s && e1s <= e2e {
			overlap++
			continue
		}
	}

	return overlap
}
