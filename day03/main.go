package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)
	file, _ := os.ReadFile("input.txt")

	if part == 1 {
		ans := part1(string(file))
		fmt.Println("Output:", ans)
	} else {
		ans := part2(string(file))
		fmt.Println("Output:", ans)
	}
}

func part1(memory string) int {
	var sumOfInstructions int
	r, _ := regexp.Compile(`mul\((\d+),(\d+)\)`)
	matches := r.FindAllStringSubmatch(memory, -1)
	for _, v := range matches {
		l, _ := strconv.Atoi(v[1])
		r, _ := strconv.Atoi(v[2])
		sumOfInstructions = sumOfInstructions + (l * r)
	}

	return sumOfInstructions
}

func part2(memory string) int {
	r, _ := regexp.Compile(`(?s)(don't\(\).+?(do\(\)|\z))`)
	cleanedMemory := r.ReplaceAllString(memory, "")
	return part1(cleanedMemory)
}
