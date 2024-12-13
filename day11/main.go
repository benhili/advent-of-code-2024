package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/benhili/advent-of-code-2024/utils"
)

func main() {
	part, file, start := utils.Setup("input.txt")

	stones := make(map[int]int)
	for _, v := range strings.Split(string(file), " ") {
		num, _ := strconv.Atoi(v)
		stones[num]++
	}

	if part == 1 {
		ans := part1(stones, 25)
		fmt.Println("Output:", ans)
	} else {
		ans := part2(stones, 75)
		fmt.Println("Output:", ans)
	}
	utils.Finish(start)
}

func part1(stones map[int]int, blinks int) int {
	for range blinks {
		newStones := make(map[int]int)
		for stone, count := range stones {
			strNum := strconv.Itoa(stone)
			if stone == 0 {
				newStones[1] += count
			} else if len(strNum)%2 == 0 {
				middle := len(strNum) / 2
				left, _ := strconv.Atoi(strNum[:middle])
				right, _ := strconv.Atoi(strNum[middle:])
				newStones[left] += count
				newStones[right] += count
			} else {
				newStones[stone*2024] += count
			}
		}
		stones = newStones
	}

	numberOfStones := 0

	for _, count := range stones {
		numberOfStones += count
	}

	return numberOfStones
}

func part2(stones map[int]int, blinks int) int {
	return part1(stones, blinks)
}
