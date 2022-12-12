package main

import (
	"embed"
	_ "embed"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/dbut2/advent-of-code/pkg/test"
	"github.com/dbut2/advent-of-code/pkg/watch"
)

//go:embed input.txt
var input string

//go:embed test*.txt
var tests embed.FS

var inc = watch.Incrementer(time.Second)

func main() {
	t := test.Register(tests, solve)
	t.Expect(1, 2)
	fmt.Println(solve(input))
}

func solve(input string) int {
	parts := strings.Split(input, "\n\n")
	rulesRaw := strings.Split(parts[0], "\n")
	s := strings.Split(parts[1], "\n")

	parsedRules := make(map[string][]string)
	for _, ruleRaw := range rulesRaw {
		broken := strings.Split(ruleRaw, " ")
		parsedRules[strings.Trim(broken[0], ":")] = broken[1:]
	}

	r := "^" + genRegex(parsedRules, "0") + "$"

	valid := 0
	for _, str := range s {
		match, err := regexp.MatchString(r, str)
		if err != nil {
			panic(err)
		}
		if match {
			valid++
		}
	}
	return valid
}

func genRegex(parsedRules map[string][]string, rule string) string {
	inc()
	r := ""
	hasOr := false
	for _, s := range parsedRules[rule] {
		if s[:1] == "\"" {
			r += strings.Trim(s, "\"")
			continue
		}
		switch s {
		case "|":
			r += s
			hasOr = true
		default:
			r += genRegex(parsedRules, s)
		}
	}
	if hasOr {
		r = "(" + r + ")"
	}
	return r
}
