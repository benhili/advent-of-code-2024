package main

import (
	"container/heap"
	"fmt"
	"image"
	"strings"

	"github.com/benhili/advent-of-code-2024/utils"
)

func main() {
	part, file, start := utils.Setup("input.txt")

	racetrack, sPos, ePos := parse(string(file))
	if part == 1 {
		ans := part1(racetrack, sPos, ePos, 141)
		fmt.Println(ans)
	} else {
		ans := part1(racetrack, sPos, ePos, 141)
		fmt.Println(ans)
	}
	utils.Finish(start)
}

func parse(input string) (map[image.Point]string, image.Point, image.Point) {
	var sPos, ePos image.Point
	racetrack := make(map[image.Point]string)
	for y, row := range strings.Split(input, "\n") {
		for x, char := range strings.Split(row, "") {
			if char == "S" {
				sPos = image.Point{x, y}
				racetrack[image.Point{x, y}] = "."
			} else if char == "E" {
				ePos = image.Point{x, y}
				racetrack[image.Point{x, y}] = "."
			} else {
				racetrack[image.Point{x, y}] = char
			}
		}
	}
	return racetrack, sPos, ePos
}

type Node struct {
	Position image.Point
	Priority int
	Index    int
}

// PriorityQueue implementation for heap interface.
type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	node := x.(*Node)
	node.Index = len(*pq)
	*pq = append(*pq, node)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	node.Index = -1
	*pq = old[0 : n-1]
	return node
}

func inBounds(limit, x, y int) bool {
	if y <= limit && x <= limit && x >= 0 && y >= 0 {
		return true
	}
	return false
}

func findShortestPath(racetrack map[image.Point]string, sPos, ePos image.Point, size int, ignoreWall image.Point) []image.Point {
	pq := &PriorityQueue{}
	heap.Init(pq)

	// Track costs for each state (position)
	costSoFar := make(map[image.Point]int)
	costSoFar[sPos] = 0
	cameFrom := make(map[image.Point]image.Point)

	heap.Push(pq, &Node{
		Position: sPos,
		Priority: 0,
	})

	directions := []image.Point{
		{X: 0, Y: -1}, // Up
		{X: 0, Y: 1},  // Down
		{X: -1, Y: 0}, // Left
		{X: 1, Y: 0},  // Right
	}

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*Node)
		currentPos := current.Position

		if currentPos == ePos {
			path := []image.Point{}
			for pos := ePos; pos != sPos; pos = cameFrom[pos] {
				path = append([]image.Point{pos}, path...)
			}
			path = append([]image.Point{sPos}, path...)
			return path
		}

		for _, dir := range directions {
			nextPos := currentPos.Add(dir)

			if (racetrack[nextPos] == "#" && ignoreWall != nextPos) || !inBounds(size, nextPos.X, nextPos.Y) {
				continue
			}

			newCost := costSoFar[currentPos] + 1

			if oldCost, exists := costSoFar[nextPos]; !exists || newCost < oldCost {
				costSoFar[nextPos] = newCost
				cameFrom[nextPos] = currentPos
				heap.Push(pq, &Node{
					Position: nextPos,
					Priority: newCost,
				})
			}
		}
	}

	return []image.Point{}
}

func quickAbs(val int) int {
	if val > -val {
		return val
	}
	return -val
}

func part1(racetrack map[image.Point]string, sPos, ePos image.Point, size int) int {
	shortestPath := findShortestPath(racetrack, sPos, ePos, size, image.Point{-9, -9})
	possibleSkips := 0
	disableCollisionMaxTime := 20

	for i := 0; i < len(shortestPath)-1; i++ {
		for j := i + 1; j < len(shortestPath); j++ {
			// calc distance between the points on path
			timeSaved := j - i

			// find distance as the crow flies
			xDiff := quickAbs(shortestPath[i].X - shortestPath[j].X)
			yDiff := quickAbs(shortestPath[i].Y - shortestPath[j].Y)

			if xDiff+yDiff <= disableCollisionMaxTime {
				saved := timeSaved - (xDiff + yDiff)
				if saved >= 100 {
					possibleSkips++
				}
			}

		}
	}

	return possibleSkips
}
