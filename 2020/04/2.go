package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
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

			kvs[kv[0]] = kv[1]
		}

		v := true

		for req, rule := range map[string]func(string) bool{
			"byr": func(s string) bool {
				v, err := strconv.Atoi(s)
				if err != nil {
					return false
				}
				return v >= 1920 && v <= 2002
			},
			"iyr": func(s string) bool {
				v, err := strconv.Atoi(s)
				if err != nil {
					return false
				}
				return v >= 2010 && v <= 2020
			},
			"eyr": func(s string) bool {
				v, err := strconv.Atoi(s)
				if err != nil {
					return false
				}
				return v >= 2020 && v <= 2030
			},
			"hgt": func(s string) bool {
				if len(s) < 2 {
					return false
				}
				if s[len(s)-2:] == "cm" {
					s = strings.Trim(s, "cm")
					v, err := strconv.Atoi(s)
					if err != nil {
						return false
					}
					return v >= 150 && v <= 193
				}
				if s[len(s)-2:] == "in" {
					s = strings.Trim(s, "in")
					v, err := strconv.Atoi(s)
					if err != nil {
						return false
					}
					return v >= 59 && v <= 76
				}
				return false
			},
			"hcl": func(s string) bool {
				found, err := regexp.MatchString("^#[a-f0-9]{6}$", s)
				if err != nil || !found {
					return false
				}
				return true
			},
			"ecl": func(s string) bool {
				switch s {
				case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
					return true
				default:
					return false
				}
			},
			"pid": func(s string) bool {
				found, err := regexp.MatchString("^[0-9]{9}$", s)
				if err != nil || !found {
					return false
				}
				return true
			},
		} {
			if _, ok := kvs[req]; !ok {
				v = false
			}

			if !rule(kvs[req]) {
				v = false
			}
		}

		if v {
			valid++
		}
	}

	return valid
}
