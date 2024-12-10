package main

import (
	"embed"
	_ "embed"
	"fmt"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/test"
	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed input.txt
var input string

//go:embed test*.txt
var tests embed.FS

func main() {
	t := test.Register(tests, solve)
	t.Expect(1, "2=-1=0")
	fmt.Println(solve(input))
}

func solve(input string) string {
	s := utils.ParseInput(input)
	_ = s

	total := 0
	for _, str := range s {
		total += snafu2dec(str)
	}

	return dec2snafu(total)
}

func snafu2dec(snafu string) int {
	e := strings.Split(snafu, "")
	dec := 0
	for _, d := range e {
		dec *= 5
		dec += s2d[d]
	}
	return dec
}

var s2d = map[string]int{
	"=": -2,
	"-": -1,
	"0": 0,
	"1": 1,
	"2": 2,
}

func dec2snafu(dec int) string {
	snafu := ""
	for i := 0; dec > 0; i++ {
		rem := ((dec + 2) % 5) - 2
		dec -= rem
		snafu = d2s[rem] + snafu
		dec /= 5
	}
	return snafu
}

var d2s = map[int]string{
	-2: "=",
	-1: "-",
	0:  "0",
	1:  "1",
	2:  "2",
}
