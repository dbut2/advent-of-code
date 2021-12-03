package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed 0.txt
var input string

func main() {
	s := strings.Split(input, "\n")
	i := solve(s)
	fmt.Println(i)
}

func solve(s []string) int {

}
