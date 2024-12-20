package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/benhili/advent-of-code-2024/utils"
)

func main() {
	part, file, start := utils.Setup("input.txt")

	a, b, c, program := parse(string(file))
	if part == 1 {
		ans := part1(a, b, c, program)
		fmt.Println(ans)
	} else {
		ans := part2(0, program)
		fmt.Println(ans)
	}
	utils.Finish(start)
}

func parse(input string) (int, int, int, []int) {
	var program []int

	r, _ := regexp.Compile(`: (\d+)`)
	splitInput := strings.Split(input, "\n\n")
	registers := splitInput[0]
	programStr := splitInput[1]
	registerMatches := r.FindAllStringSubmatch(registers, -1)

	a, _ := strconv.Atoi(registerMatches[0][1])
	b, _ := strconv.Atoi(registerMatches[1][1])
	c, _ := strconv.Atoi(registerMatches[2][1])
	for _, strnum := range strings.Split(strings.Split(programStr, "Program: ")[1], ",") {
		num, _ := strconv.Atoi(strnum)
		program = append(program, num)
	}
	return a, b, c, program
}

func dv(operand, a int) int {
	denominator := math.Pow(2, float64(operand))
	a = a / int(denominator)
	return int(a)
}

func part1(a, b, c int, program []int) []int {
	i := 0
	var out []int
	for i < len(program)-1 {
		opcode := program[i]
		operand := program[i+1]
		var comboOperand int
		if operand == 4 {
			comboOperand = a
		} else if operand == 5 {
			comboOperand = b
		} else if operand == 6 {
			comboOperand = c
		}
		switch opcode {
		case 0:
			a = dv(operand, a)
		case 1:
			b = operand ^ b
		case 2:
			b = comboOperand % 8
		case 3:
			if a != 0 {
				i = operand
				continue
			}
		case 4:
			b = b ^ c
		case 5:
			out = append(out, comboOperand%8)
		case 6:
			b = dv(comboOperand, a)
		case 7:
			c = dv(comboOperand, a)

		}
		i += 2
		if len(out) > len(program) {
			break
		}
	}
	return out
}

func part2(ans int, program []int) int {
	fmt.Println(ans, program)
	if len(program) == 0 {
		return ans
	}
	for i := 0; i < 8; i++ {
		a := ans<<3 | i
		b := a % 8
		b = 3 ^ b
		c := dv(b, a)
		b = 5 ^ b
		b = b ^ c

		if b%8 == program[len(program)-1] {
			res := part2(a, program[:len(program)-1])
			if res == -1 {
				continue
			}
			return res
		}
	}
	return -1
}
