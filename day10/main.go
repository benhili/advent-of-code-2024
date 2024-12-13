package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/benhili/advent-of-code-2024/utils"
)

func main() {
	part, file, start := utils.Setup("input.txt")
	trailMap := parse(file)

	if part == 1 {
		ans := part1(trailMap)
		fmt.Println("Output:", ans)
	} else {
		ans := part2(trailMap)
		fmt.Println("Output:", ans)
	}
	utils.Finish(start)
}

func parse(file []byte) [][]int {
	var trailMap [][]int
	for _, line := range strings.Split(string(file), "\n") {
		var row []int
		for _, char := range strings.Split(line, "") {
			num, _ := strconv.Atoi(char)
			row = append(row, num)
		}
		trailMap = append(trailMap, row)
	}
	return trailMap
}

func inBounds(input [][]int, y, x int) bool {
	yLimit := len(input)
	xLimit := len(input[0])
	if y < yLimit && x < xLimit && x > -1 && y > -1 {
		return true
	}
	return false
}

func findPaths(trailMap [][]int, y, x int, peaks map[string]bool, rating bool) int {
	sum := 0
	curr := trailMap[y][x]
	if curr == 9 {
		if !rating {
			if _, ok := peaks[fmt.Sprintf("%d:%d", x, y)]; ok {
				return 0
			}
			peaks[fmt.Sprintf("%d:%d", x, y)] = true
		}
		return 1
	}

	if inBounds(trailMap, y-1, x) && trailMap[y-1][x] == curr+1 {
		sum += findPaths(trailMap, y-1, x, peaks, rating)
	}
	if inBounds(trailMap, y+1, x) && trailMap[y+1][x] == curr+1 {
		sum += findPaths(trailMap, y+1, x, peaks, rating)
	}
	if inBounds(trailMap, y, x-1) && trailMap[y][x-1] == curr+1 {
		sum += findPaths(trailMap, y, x-1, peaks, rating)
	}
	if inBounds(trailMap, y, x+1) && trailMap[y][x+1] == curr+1 {
		sum += findPaths(trailMap, y, x+1, peaks, rating)
	}

	return sum
}

func part1(trailMap [][]int) int {
	cumulitiveSum := 0
	for y, row := range trailMap {
		for x, height := range row {
			if height == 0 {
				peaks := make(map[string]bool)
				cumulitiveSum += findPaths(trailMap, y, x, peaks, false)
			}
		}
	}
	return cumulitiveSum
}

func part2(trailMap [][]int) int {
	cumulitiveSum := 0
	for y, row := range trailMap {
		for x, height := range row {
			if height == 0 {
				peaks := make(map[string]bool)
				cumulitiveSum += findPaths(trailMap, y, x, peaks, true)
			}
		}
	}
	return cumulitiveSum
}
