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

	rMap, sPos, ePos := parse(string(file))
	if part == 1 {
		ans := part1(rMap, sPos, ePos)
		fmt.Println("Output:", ans)
	} else {
		ans := part2(rMap, sPos, ePos)
		fmt.Println("Output:", ans)
	}
	utils.Finish(start)
}

func parse(input string) (map[image.Point]string, image.Point, image.Point) {
	rMap := make(map[image.Point]string)
	var sPos image.Point
	var ePos image.Point
	for y, row := range strings.Split(input, "\n") {
		for x, char := range strings.Split(row, "") {
			if char == "E" {
				ePos = image.Point{x, y}
				rMap[ePos] = "."
			} else if char == "S" {
				sPos = image.Point{x, y}
				rMap[sPos] = "."
			} else {
				rMap[image.Point{x, y}] = char
			}
		}
	}
	return rMap, sPos, ePos
}

type Node struct {
	Position  image.Point
	Direction image.Point
	Priority  int
	Index     int
	Path      []PathStep
}

type PathStep struct {
	Position  image.Point
	Direction image.Point
	Cost      int
}

// State represents a unique state in the path (position + direction)
type State struct {
	Position  image.Point
	Direction image.Point
}

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

func findShortestPath(rMap map[image.Point]string, sPos, ePos image.Point) (int, int) {
	pq := &PriorityQueue{}
	heap.Init(pq)

	// Track costs for each state (position + direction)
	costSoFar := make(map[State]int)

	shortestPaths := [][]PathStep{}
	initialState := State{Position: sPos, Direction: image.Point{1, 0}}
	costSoFar[initialState] = 0
	heap.Push(pq, &Node{
		Position:  sPos,
		Direction: image.Point{1, 0},
		Priority:  0,
	})
	minCostFound := -1

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*Node)
		currentState := State{Position: current.Position, Direction: current.Direction}

		// Skip if we've found a better path to this state
		if cost, exists := costSoFar[currentState]; exists && cost < current.Priority {
			continue
		}

		if current.Position == ePos {
			if minCostFound == -1 || current.Priority == minCostFound {
				minCostFound = current.Priority
				shortestPaths = append(shortestPaths, current.Path)
			} else if current.Priority > minCostFound {
				continue // Skip if we've found a better path
			}
		}

		// Get possible next directions (forward, left, right)
		possibleDirs := []image.Point{
			current.Direction, // continue forward
			{X: -current.Direction.Y, Y: current.Direction.X}, // turn left
			{X: current.Direction.Y, Y: -current.Direction.X}, // turn right
		}

		for _, nextDir := range possibleDirs {
			nextPos := current.Position.Add(nextDir)

			// Check if the next position is valid
			if rMap[nextPos] != "." {
				continue
			}

			// Calculate rotation cost
			rotationCost := 0
			if nextDir != current.Direction {
				rotationCost = 1000 // Cost for 90-degree rotation
			}

			nextState := State{Position: nextPos, Direction: nextDir}
			newCost := current.Priority + 1 + rotationCost // Base movement cost + rotation cost

			if oldCost, exists := costSoFar[nextState]; !exists || newCost <= oldCost {
				costSoFar[nextState] = newCost
				newPath := make([]PathStep, len(current.Path))
				copy(newPath, current.Path)
				newPath = append(newPath, PathStep{
					Position:  nextPos,
					Direction: nextDir,
					Cost:      newCost,
				})
				heap.Push(pq, &Node{
					Position:  nextPos,
					Direction: nextDir,
					Priority:  newCost,
					Path:      newPath,
				})
			}
		}
	}

	seatPositions := 0
	seen := make(map[image.Point]bool)
	for _, pathSteps := range shortestPaths {
		for _, step := range pathSteps {
			if !seen[step.Position] {
				seatPositions++
				seen[step.Position] = true
			}
		}
	}
	return minCostFound, seatPositions + 1
}

func part1(rMap map[image.Point]string, sPos, ePos image.Point) int {
	minCost, _ := findShortestPath(rMap, sPos, ePos)
	return minCost
}

func part2(rMap map[image.Point]string, sPos, ePos image.Point) int {
	_, pathSum := findShortestPath(rMap, sPos, ePos)
	return pathSum
}
