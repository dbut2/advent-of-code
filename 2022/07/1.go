package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/sti"
	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed input.txt
var input string

//go:embed test1.txt
var test string

func main() {
	utils.Test(solve(test), 95437)
	fmt.Println(solve(input))
}

func solve(input string) int {
	s := strings.Split(strings.Trim(input, "$ "), "\n$ ")
	root := &Tree{
		Name: "/",
	}
	root.Root = root

	pointer := root

	for _, str := range s {
		lines := strings.Split(str, "\n")
		firstLine := strings.Split(lines[0], " ")

		cmd := firstLine[0]

		switch cmd {
		case "cd":
			pointer = pointer.MoveTo(firstLine[1])
		case "ls":
			for _, file := range lines[1:] {
				line := strings.Split(file, " ")

				size := 0
				ft := Dir
				if line[0] != "dir" {
					size = sti.Sti(line[0])
					ft = File
				}

				pointer.AddChild(line[1], size, ft)
			}
		}
	}

	pointer = pointer.MoveTo("/")

	return pointer.TotalOfChildDirectoriesLessThanOrEqualTo(100000)
}

type Tree struct {
	Name     string
	Type     Type
	Root     *Tree
	Parent   *Tree
	Children []*Tree
	Size     int
}

type Type int

const (
	Unknown Type = iota
	Dir
	File
)

func (t Tree) TotalOfChildDirectoriesLessThanOrEqualTo(size int) int {
	if t.Type == File {
		return 0
	}

	total := 0
	if t.SizeOf() <= size {
		total += t.SizeOf()
	}
	for _, child := range t.Children {
		total += child.TotalOfChildDirectoriesLessThanOrEqualTo(size)
	}
	return total
}

func (t *Tree) MoveUp() *Tree {
	return t.Parent
}

func (t *Tree) MoveDown(dir string) *Tree {
	for _, child := range t.Children {
		if child.Name == dir {
			return child
		}
	}
	child := &Tree{
		Name:   dir,
		Type:   Dir,
		Root:   t.Root,
		Parent: t,
	}
	t.Children = append(t.Children, child)
	return child
}

func (t *Tree) AddChild(name string, size int, ft Type) {
	child := &Tree{
		Name:   name,
		Type:   ft,
		Size:   size,
		Root:   t.Root,
		Parent: t,
	}
	t.Children = append(t.Children, child)
}

func (t *Tree) SizeOf() int {
	size := t.Size
	for _, child := range t.Children {
		size += child.SizeOf()
	}
	return size
}

func (t *Tree) MoveTo(dir string) *Tree {
	if dir == "" {
		return t
	}

	path := strings.Split(dir, "/")

	current := path[0]
	next := strings.Join(path[1:], "/")

	switch current {
	case ".":
		return t.MoveTo(next)
	case "..":
		return t.MoveUp().MoveTo(next)
	case "":
		return t.Root.MoveTo(next)
	default:
		return t.MoveDown(current).MoveTo(next)
	}
}
