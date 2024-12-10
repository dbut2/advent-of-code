package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/math"
	"github.com/dbut2/advent-of-code/pkg/sets"
	"github.com/dbut2/advent-of-code/pkg/sti"
	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed input.txt
var input string

//go:embed test1.txt
var test string

func main() {
	utils.Test(solve(test), 24933642)
	fmt.Println(solve(input))
}

func solve(input string) int {
	s := utils.ParseInput(input)
	curr := ""
	var dirs sets.Set[string]
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
			files[curr+"/"+args[1]] = sti.Sti(args[0])
		}
	}
	ma := -1
	used := math.SumMap(files)
	needToDelete := sti.Sti("30000000") - (sti.Sti("70000000") - used)
	for _, dir := range dirs.Slice() {
		dirsize := math.SumMapIf(files, func(file string) bool {
			match, _ := regexp.MatchString(dir+"/.*", file)
			return match
		})
		if dirsize >= needToDelete && (ma == -1 || dirsize < ma) {
			ma = dirsize
		}
	}
	return ma
}
