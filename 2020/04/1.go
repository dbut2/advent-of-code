package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	s := strings.Split(input, "\n\n")
	i := solve(s)
	fmt.Println(i)
}

func solve(s []string) int {
	valid := 0

	for _, str := range s {
		str = strings.ReplaceAll(str, "\n", " ")

		kvs := map[string]string{}

		fields := strings.Split(str, " ")
		for _, field := range fields {
			kv := strings.Split(field, ":")

			kvs[kv[0]] = kvs[kv[1]]
		}

		v := true

		for _, req := range []string{
			"byr",
			"iyr",
			"eyr",
			"hgt",
			"hcl",
			"ecl",
			"pid",
		} {
			if _, ok := kvs[req]; !ok {
				v = false
			}
		}

		if v {
			valid++
		}
	}

	return valid
}
