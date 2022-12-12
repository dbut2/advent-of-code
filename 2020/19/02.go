package main

import (
	"embed"
	_ "embed"
	"fmt"
	"regexp"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/test"
)

//go:embed input.txt
var input string

//go:embed test*.txt
var tests embed.FS

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

	//parsedRules["42"] = []string{"\"A\""}
	//parsedRules["31"] = []string{"\"B\""}

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
	if rule == "8" {
		return "(" + genRegex(parsedRules, "42") + ")+"
	}
	if rule == "11" {
		r := "("
		for i := 1; i < 5; i++ {
			for j := 0; j < i; j++ {
				r += genRegex(parsedRules, "42")
			}
			for j := 0; j < i; j++ {
				r += genRegex(parsedRules, "31")
			}
			r += "|"
		}
		r = strings.TrimSuffix(r, "|") + ")"
		return r
	}
	r := ""
	hasOr := false
	rep := false
	for _, s := range parsedRules[rule] {
		if s[:1] == "\"" {
			r += strings.Trim(s, "\"")
			continue
		}
		switch s {
		case "|":
			r += s
			hasOr = true
		case rule:
			rep = true
		default:
			r += genRegex(parsedRules, s)
		}
	}
	if hasOr {
		r = "(" + r + ")"
	}
	if rep {
		r = "(" + r + ")*"
	}
	return r
}
