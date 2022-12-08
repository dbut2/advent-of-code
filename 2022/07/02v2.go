package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strings"

	"github.com/dbut2/advent-of-code/utils"
)

//go:embed input.txt
var input string

//go:embed test.txt
var test string

func main() {
	fmt.Println("Test")
	fmt.Println(do(test))
	fmt.Println()
	fmt.Println("Solution")
	fmt.Println(do(input))
}

func do(s string) int {
	strs := strings.Split(s, "\n")
	return solve(strs)
}

func solve(s []string) int {
	curr := ""
	var dirs utils.Set[string]
	files := make(map[string]int)
	for _, line := range s {
		if line == "$ cd /" {
			curr = ""
			continue
		}
		args := strings.Split(line, " ")
		switch args[0] {
		case "$":
			if args[1] == "cd" {
				switch args[2] {
				case "..":
					c := strings.Split(curr, "/")
					curr = strings.Join(c[:len(c)-1], "/")
				default:
					curr = curr + "/" + args[2]
					dirs.Add(curr)
				}
			}
		case "dir":
		default:
			files[curr+"/"+args[1]] = utils.Sti(args[0])
		}
	}
	ma := -1
	used := utils.SumMap(files)
	needToDelete := utils.Sti("30000000") - (utils.Sti("70000000") - used)
	for _, dir := range dirs.Slice() {
		dirsize := utils.SumMapIf(files, func(file string) bool {
			match, _ := regexp.MatchString(dir+"/.*", file)
			return match
		})
		if dirsize >= needToDelete && (ma == -1 || dirsize < ma) {
			ma = dirsize
		}
	}
	return ma
}
