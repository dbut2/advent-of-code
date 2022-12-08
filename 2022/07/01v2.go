package main

import (
	_ "embed"
	"fmt"
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
	dirs := map[string]bool{"": true}
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
					dirs[curr] = true
				}
			}
		case "dir":
			dirs[curr+"/"+args[1]] = true
		default:
			files[curr+"/"+args[1]] = utils.Sti(args[0])
		}
	}
	total := 0
	for dir := range dirs {
		dirsize := 0
		for file, size := range files {
			if len(dir) >= len(file) {
				continue
			}
			match := file[len(dir)] == "/"[0]
			for i := range dir {
				if dir[i] != file[i] {
					match = false
				}
			}
			if match {
				dirsize += size
			}
		}
		if dirsize <= 100000 {
			total += dirsize
		}
	}
	return total
}
