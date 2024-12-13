package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/benhili/advent-of-code-2024/utils"
)

type coords struct {
	y int
	x int
}

type edge struct {
	y    int
	x    int
	side int
}

func main() {
	part, file, start := utils.Setup("input.txt")

	plantMap := fileTo2dStringSlice(string(file))
	if part == 1 {
		ans := part1(plantMap)
		fmt.Println("------------------")
		fmt.Println("Output:", ans)
	} else {
		ans := part2(plantMap)
		fmt.Println("Output:", ans)
	}
	utils.Finish(start)
}

func fileTo2dStringSlice(file string) [][]string {
	var plantMap [][]string
	for _, row := range strings.Split(file, "\n") {
		plantMap = append(plantMap, strings.Split(row, ""))

	}
	return plantMap
}

func inBounds(input [][]string, y, x int) bool {
	yLimit := len(input)
	xLimit := len(input[0])
	if y < yLimit && x < xLimit && x > -1 && y > -1 {
		return true
	}
	return false
}

func dfs(plantMap [][]string, y, x int, seen map[coords]bool) (int, int) {
	if seen[coords{y, x}] {
		return 0, 0 // If we've already seen this coordinate, no additional area or perimeter
	}
	directions := [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	seen[coords{y, x}] = true
	area := 1
	perimeter := 0
	for _, v := range directions {
		newY, newX := v[0]+y, v[1]+x
		if !inBounds(plantMap, newY, newX) || plantMap[newY][newX] != plantMap[y][x] {
			perimeter++
		} else if !seen[coords{newY, newX}] {
			newArea, newPerimeter := dfs(plantMap, newY, newX, seen)
			area += newArea
			perimeter += newPerimeter
		}
	}
	return area, perimeter
}

func part1(plantMap [][]string) int {
	sum := 0
	seen := make(map[coords]bool)
	for y, row := range plantMap {
		for x := range row {
			if !seen[coords{y, x}] {
				area, perimeter := dfs(plantMap, y, x, seen)
				sum += area * perimeter
			}
		}
	}
	return sum
}

func dfs2(plantMap [][]string, y, x int, seen map[coords]bool) (int, []edge) {
	if seen[coords{y, x}] {
		return 0, []edge{} // If we've already seen this coordinate, no additional area or perimeter
	}
	directions := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	seen[coords{y, x}] = true
	var edges []edge
	area := 1
	for i, v := range directions {
		newY, newX := v[0]+y, v[1]+x
		if !inBounds(plantMap, newY, newX) || plantMap[newY][newX] != plantMap[y][x] {
			edges = append(edges, edge{y, x, i})
		} else if !seen[coords{newY, newX}] {
			newArea, newEdges := dfs2(plantMap, newY, newX, seen)
			area += newArea
			edges = append(edges, newEdges...)
		}
	}
	return area, edges
}

func countSides(edges []edge) int {
	sides := 0
	seen := make(map[edge]bool)
	for _, e := range edges {
		if seen[e] {
			continue
		}
		if e.side == 1 || e.side == 3 {
			// horizontal
			sides++
			// go up
			up := edge{y: e.y - 1, x: e.x, side: e.side}
			for slices.Contains(edges, up) {
				seen[up] = true
				up = edge{y: up.y - 1, x: up.x, side: up.side}
			}
			// go down
			down := edge{y: e.y + 1, x: e.x, side: e.side}
			for slices.Contains(edges, down) {
				seen[down] = true
				down = edge{y: down.y + 1, x: down.x, side: down.side}
			}
		}
		if e.side == 0 || e.side == 2 {
			// vertical
			sides++

			// go left
			left := edge{y: e.y, x: e.x - 1, side: e.side}
			for slices.Contains(edges, left) {
				seen[left] = true
				left = edge{y: left.y, x: left.x - 1, side: left.side}
			}
			// go right
			right := edge{y: e.y, x: e.x + 1, side: e.side}
			for slices.Contains(edges, right) {
				seen[right] = true
				right = edge{y: right.y, x: right.x + 1, side: right.side}
			}
		}
	}
	return sides
}

func part2(plantMap [][]string) int {
	sum := 0
	seen := make(map[coords]bool)
	for y, row := range plantMap {
		for x := range row {
			if !seen[coords{y, x}] {
				area, edges := dfs2(plantMap, y, x, seen)
				sideCount := countSides(edges)
				sum += area * sideCount
			}
		}
	}
	return sum
}
