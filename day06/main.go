package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)
	file, _ := os.ReadFile("input.txt")
	labMap := parse(file)

	if part == 1 {
		ans := part1(labMap)
		fmt.Println("Output:", ans)
	} else {
		ans := part2(labMap)
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

func inBounds(input [][]string, y, x int) bool {
	yLimit := len(input)
	xLimit := len(input[0])
	if y < yLimit && x < xLimit && x > -1 && y > -1 {
		return true
	}
	return false
}

func getGuardPos(input [][]string) (int, int) {
	for y, row := range input {
		for x, char := range row {
			if char == "^" {
				return y, x
			}
		}
	}
	panic("NO GUARD IN MAP")
}

func getForwardCoords(y, x, direction int) (int, int) {
	switch direction {
	case 0:
		return y - 1, x
	case 1:
		return y, x + 1
	case 2:
		return y + 1, x
	case 3:
		return y, x - 1
	}
	panic("invalid direction")
}

// directions
// 0 NORTH, 1 EAST, 2 SOUTH, 3 WEST
func part1(labMap [][]string) int {
	sum := 1
	visited := make([][]string, len(labMap))
	copy(visited, labMap)
	y, x := getGuardPos(labMap)
	direction := 0
	for {
		forwardY, forwardX := getForwardCoords(y, x, direction)
		if !inBounds(labMap, forwardY, forwardX) {
			return sum
		}
		// obstacle
		if labMap[forwardY][forwardX] == "#" {
			direction = (direction + 1) % 4
		} else {
			y = forwardY
			x = forwardX
			if visited[y][x] != "X" {
				sum += 1
				visited[forwardY][forwardX] = "X"
			}
		}
	}
}

func hasLoop(labMap [][]string) bool {
	visited := make([][]string, len(labMap))
	copy(visited, labMap)
	y, x := getGuardPos(labMap)
	direction := 0
	for {
		forwardY, forwardX := getForwardCoords(y, x, direction)
		if !inBounds(labMap, forwardY, forwardX) {
			// guard escaped
			return false
		}
		// obstacle
		if labMap[forwardY][forwardX] == "#" {
			direction = (direction + 1) % 4
		} else {
			y = forwardY
			x = forwardX
			stringDirection := strconv.Itoa(direction)
			if visited[y][x] == stringDirection {
				return true
			}
			visited[y][x] = stringDirection
		}
	}
}

func part2(labMap [][]string) int {
	sum := 0
	for y := range labMap {
		for x := range labMap[y] {
			currentOccupant := labMap[y][x]
			if currentOccupant == "#" || currentOccupant == "^" {
				continue
			}
			extraObstacleMap := make([][]string, len(labMap))
			for i := range labMap {
				extraObstacleMap[i] = make([]string, len(labMap[i]))
				copy(extraObstacleMap[i], labMap[i])
			}

			extraObstacleMap[y][x] = "#"
			if hasLoop(extraObstacleMap) {
				sum += 1
			}
		}
	}
	return sum
}
