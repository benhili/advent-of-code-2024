package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)
	file, _ := os.ReadFile("input.txt")
	input := parse(file)

	if part == 1 {
		ans := part1(input)
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input)
		fmt.Println("Output:", ans)
	}
}

func parse(file []byte) [][]string {
	var result [][]string
	lines := strings.Fields(string(file))
	for _, v := range lines {
		result = append(result, strings.Split(v, ""))
	}
	return result
}

var targets = map[string]string{
	"X": "M",
	"M": "A",
	"A": "S",
}

func inBounds(input [][]string, y, x int) bool {
	yLimit := len(input)
	xLimit := len(input[0])
	if y < yLimit && x < xLimit && x > -1 && y > -1 {
		return true
	}
	return false
}

// Directions
// 0 UP
// 1 DIAGONAL UP RIGHT
// 2 RIGHT
// 3 DIAGONAL DOWN RIGHT
// 4 DOWN
// 5 DIAGONAL DOWN LEFT
// 6 LEFT
// 7 DIAGONAL UP LEFT

func seek(input [][]string, direction, y, x int, target string) bool {
	if input[y][x] == target {
		if target == "S" {
			return true
		}
		nextTarget := targets[target]
		switch direction {
		case 0:
			newY := y - 1
			newX := x
			if inBounds(input, newY, newX) {
				return seek(input, direction, newY, newX, nextTarget)
			}
		case 1:
			newY := y - 1
			newX := x + 1
			if inBounds(input, newY, newX) {
				return seek(input, direction, newY, newX, nextTarget)
			}
		case 2:
			newY := y
			newX := x + 1
			if inBounds(input, newY, newX) {
				return seek(input, direction, newY, newX, nextTarget)
			}
		case 3:
			newY := y + 1
			newX := x + 1
			if inBounds(input, newY, newX) {
				return seek(input, direction, newY, newX, nextTarget)
			}
		case 4:
			newY := y + 1
			newX := x
			if inBounds(input, newY, newX) {
				return seek(input, direction, newY, newX, nextTarget)
			}
		case 5:
			newY := y + 1
			newX := x - 1
			if inBounds(input, newY, newX) {
				return seek(input, direction, newY, newX, nextTarget)
			}
		case 6:
			newY := y
			newX := x - 1
			if inBounds(input, newY, newX) {
				return seek(input, direction, newY, newX, nextTarget)
			}
		case 7:
			newY := y - 1
			newX := x - 1
			if inBounds(input, newY, newX) {
				return seek(input, direction, newY, newX, nextTarget)
			}
		}
	}
	return false
}

func part1(input [][]string) int {
	directions := []int{0, 1, 2, 3, 4, 5, 6, 7}
	total := 0
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			if input[y][x] == "X" {
				for _, direction := range directions {
					found := seek(input, direction, y, x, "X")
					if found {
						total = total + 1
					}
				}
			}
		}
	}
	return total
}

func part2(input [][]string) int {
	total := 0
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			if input[y][x] == "A" {
				if !inBounds(input, y-1, x-1) || !inBounds(input, y+1, x-1) || !inBounds(input, y-1, x+1) || !inBounds(input, y+1, x+1) {
					continue
				}
				topLeft := input[y-1][x-1]
				topRight := input[y-1][x+1]
				botLeft := input[y+1][x-1]
				botRight := input[y+1][x+1]
				firstDiagonalMatches := topRight == "M" && botLeft == "S" || topRight == "S" && botLeft == "M"
				secondDiagonalMatches := botRight == "M" && topLeft == "S" || botRight == "S" && topLeft == "M"
				if firstDiagonalMatches && secondDiagonalMatches {
					total = total + 1
				}
			}
		}
	}
	return total
}
