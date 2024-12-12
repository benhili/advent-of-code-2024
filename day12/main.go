package main

import (
	"fmt"

	"github.com/benhili/advent-of-code-2024/utils"
)

func main() {
	part, file, start := utils.Setup()
	if part == 1 {
		ans := part1(string(file))
		fmt.Println("Output:", ans)
	} else {
		// ans := part2(stones, 75)
		// fmt.Println("Output:", ans)
	}
	utils.Finish(start)
}

func part1(input string) int {
	fmt.Println(input)
	return 10
}
