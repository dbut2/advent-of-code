package main

import (
	"embed"
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/sti"
	"github.com/dbut2/advent-of-code/pkg/test"
	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed input.txt
var input string

//go:embed test*.txt
var tests embed.FS

func main() {
	t := test.Register(tests, solve)
	t.Expect(1, 231)
	t.Expect(2, 51)
	t.Expect(3, 46)
	t.Expect(4, 1445)
	t.Expect(5, 669060)
	t.Expect(6, 23340)
	fmt.Println(solve(input))
}

func solve(input string) int {
	s := utils.ParseInput(input)

	sum := 0
	for _, eq := range s {
		sum += math(eq)
	}
	return sum
}

func math(s string) int {
	ops := strings.ReplaceAll(s, " ", "")
	level := 0

	// Avoid real parsing at all cost, this just forces multiplication to be executed last by wrapping everything else in brackets
	ops = "(" + ops + ")"
	ops = strings.ReplaceAll(ops, "(", "((")
	ops = strings.ReplaceAll(ops, ")", "))")
	ops = strings.ReplaceAll(ops, "*", ")*(")

	number := map[int]string{0: "0"}
	operand := map[int]string{0: "+"}
	buffer := map[int]string{0: ""}

	for _, char := range strings.Split(ops, "") {
		switch char {
		case "(":
			level++
			number[level] = "0"
			operand[level] = "+"
			buffer[level] = ""
		case ")":
			n := levelDown(level, number, operand, buffer)
			level--
			buffer[level] = n
		case "+":
			number[level] = doMath(number[level], operand[level], buffer[level])
			operand[level] = "+"
			buffer[level] = ""
		case "*":
			number[level] = doMath(number[level], operand[level], buffer[level])
			operand[level] = "*"
			buffer[level] = ""
		default:
			buffer[level] += char
		}
	}

	return sti.Sti(levelDown(level, number, operand, buffer))
}

func levelDown(level int, number map[int]string, operand map[int]string, buffer map[int]string) string {
	n := doMath(number[level], operand[level], buffer[level])
	delete(number, level)
	delete(operand, level)
	delete(buffer, level)
	return n
}

func doMath(a string, operand string, b string) string {
	if a == "" {
		a = "0"
	}
	if b == "" {
		b = "0"
	}
	ai, bi := sti.Sti(a), sti.Sti(b)
	switch operand {
	case "+":
		return strconv.Itoa(ai + bi)
	case "*":
		return strconv.Itoa(ai * bi)
	default:
		panic(fmt.Sprintf("operand not defined: %s", operand))
	}
}
