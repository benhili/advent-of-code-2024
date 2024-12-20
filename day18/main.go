package main

import (
	"container/heap"
	"fmt"
	"image"
	"strconv"
	"strings"

	"github.com/benhili/advent-of-code-2024/utils"
)

func main() {
	part, file, start := utils.Setup("input.txt")

	res := parse(string(file))
	if part == 1 {
		ans := part1(res, 70)
		fmt.Println(ans)
	} else {
		ans := part2(res, 70)
		fmt.Println(ans)
	}
	utils.Finish(start)
}

func parse(input string) map[image.Point]int {
	points := make(map[image.Point]int)
	for i, row := range strings.Split(input, "\n") {
		x, _ := strconv.Atoi(strings.Split(row, ",")[0])
		y, _ := strconv.Atoi(strings.Split(row, ",")[1])
		points[image.Point{x, y}] = i
	}
	return points
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

// findShortestPath calculates the shortest path from start (sPos) to end (ePos) on a given map.
func findShortestPath(invalidPositions map[image.Point]int, sPos, ePos image.Point, size, bytesToInclude int) int {
	pq := &PriorityQueue{}
	heap.Init(pq)

	// Track costs for each state (position)
	costSoFar := make(map[image.Point]int)
	costSoFar[sPos] = 0

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
			return current.Priority
		}

		for _, dir := range directions {
			nextPos := currentPos.Add(dir)

			if id, ok := invalidPositions[nextPos]; (ok && id <= bytesToInclude) || !inBounds(size, nextPos.X, nextPos.Y) {
				continue
			}

			newCost := costSoFar[currentPos] + 1

			if oldCost, exists := costSoFar[nextPos]; !exists || newCost < oldCost {
				costSoFar[nextPos] = newCost
				heap.Push(pq, &Node{
					Position: nextPos,
					Priority: newCost,
				})
			}
		}
	}

	return -1
}

func part1(invalidPositions map[image.Point]int, size int) int {
	start := image.Point{0, 0}
	end := image.Point{size, size}
	return findShortestPath(invalidPositions, start, end, size, 1024)
}

func part2(invalidPositions map[image.Point]int, size int) int {
	start := image.Point{0, 0}
	end := image.Point{size, size}

	var cost int
	includedBytes := 0
	for cost != -1 {
		cost = findShortestPath(invalidPositions, start, end, size, includedBytes)
		includedBytes++
	}

	return includedBytes
}
