package main

import (
	"fmt"
	"image"
	"strings"
	"time"

	"github.com/benhili/advent-of-code-2024/utils"
	"github.com/gdamore/tcell"
)

func main() {
	part, file, start := utils.Setup("input.txt")

	splitInput := strings.Split(string(file), "\n\n")
	rawMap := splitInput[0]
	rawMoves := splitInput[1]
	warehouseMap, startPos := parseWarehouse(rawMap)
	wideMap, rPos := parseWideWarehouse(rawMap)
	moves := parseMoves(rawMoves)
	if part == 1 {
		ans := part1(warehouseMap, moves, startPos)
		fmt.Println("Output:", ans)
	} else {
		ans := part2(wideMap, moves, rPos)
		fmt.Println("Output:", ans)
	}
	utils.Finish(start)
}

func parseWarehouse(rawMap string) ([][]string, image.Point) {

	var warehouseMap [][]string
	var start image.Point
	for y, row := range strings.Split(rawMap, "\n") {
		var newRow []string
		for x, char := range strings.Split(row, "") {
			if char == "@" {
				start = image.Point{x, y}
				newRow = append(newRow, ".")
			} else {

				newRow = append(newRow, char)
			}
		}
		warehouseMap = append(warehouseMap, newRow)
	}
	return warehouseMap, start
}

func parseWideWarehouse(rawMap string) ([][]string, image.Point) {
	var warehouseMap [][]string
	var rPos image.Point
	for y, row := range strings.Split(rawMap, "\n") {
		var newRow []string
		for x, char := range strings.Split(row, "") {
			if char == "@" {
				rPos = image.Point{x * 2, y}
				newRow = append(newRow, ".")
				newRow = append(newRow, ".")
			} else if char == "#" {
				newRow = append(newRow, char)
				newRow = append(newRow, char)
			} else if char == "O" {
				newRow = append(newRow, "[")
				newRow = append(newRow, "]")

			} else if char == "." {
				newRow = append(newRow, char)
				newRow = append(newRow, char)
			}
		}
		warehouseMap = append(warehouseMap, newRow)
	}
	return warehouseMap, rPos
}

var DIRECTIONS = map[string]image.Point{
	"left":  image.Point{-1, 0},
	"right": image.Point{1, 0},
	"up":    image.Point{0, -1},
	"down":  image.Point{0, 1},
}

type Box struct {
	pos  image.Point
	side string
}

func parseMoves(rawMoves string) []image.Point {
	var moves []image.Point
	for _, v := range strings.Split(rawMoves, "") {
		if v == "^" {
			moves = append(moves, DIRECTIONS["up"])
		} else if v == ">" {
			moves = append(moves, DIRECTIONS["right"])

		} else if v == "<" {
			moves = append(moves, DIRECTIONS["left"])

		} else if v == "v" {
			moves = append(moves, DIRECTIONS["down"])
		}
	}

	return moves
}

func inBounds(input [][]string, pos image.Point) bool {
	yLimit := len(input)
	xLimit := len(input[0])
	if pos.Y < yLimit && pos.X < xLimit && pos.X > -1 && pos.Y > -1 {
		return true
	}
	return false
}

func print2DArray(grid [][]string, robot image.Point) {
	for y, row := range grid {
		for x, val := range row {
			// Adjust the width here (e.g., %3d) to format the spacing
			if robot.X == x && robot.Y == y {
				fmt.Print("@")
			} else {
				fmt.Print(val)
			}
		}
		fmt.Println() // Move to the next line after printing a row
	}
}

func part1(warehouseMap [][]string, moves []image.Point, rPos image.Point) int {
	for _, move := range moves {
		newPos := rPos.Add(move)
		itemInFront := warehouseMap[newPos.Y][newPos.X]

		// free space
		if itemInFront == "." {
			rPos = newPos
			continue
		}
		// wall
		if itemInFront == "#" {
			continue
		}
		// box
		if itemInFront == "O" {
			next := newPos
			for {
				next = next.Add(move)
				if !inBounds(warehouseMap, next) || warehouseMap[next.Y][next.X] == "#" {
					break
				}
				if warehouseMap[next.Y][next.X] == "." {
					rPos = newPos
					warehouseMap[newPos.Y][newPos.X] = "."
					warehouseMap[next.Y][next.X] = "O"
					break
				}
			}
			continue
		}
	}

	sumCoords := 0
	for y, row := range warehouseMap {
		for x, char := range row {
			if char == "O" {
				sumCoords += (y * 100) + x

			}
		}
	}
	return sumCoords
}

func findConnectedBoxes(warehouseMap [][]string, lPos, rPos, move image.Point) ([]Box, bool) {
	var boxesToMove []Box
	var valid = true
	nextL := lPos.Add(move)
	nextR := rPos.Add(move)
	if !inBounds(warehouseMap, nextL) || warehouseMap[nextL.Y][nextL.X] == "#" || warehouseMap[nextR.Y][nextR.X] == "#" {
		return []Box{}, false
	}
	if warehouseMap[nextL.Y][nextL.X] == "[" {
		boxesToMove = append(boxesToMove, Box{pos: nextL, side: "["})
		boxesToMove = append(boxesToMove, Box{pos: nextR, side: "]"})
		newBoxes, nvalid := findConnectedBoxes(warehouseMap, nextL, nextR, move)
		if !nvalid {
			return []Box{}, false
		}
		boxesToMove = append(boxesToMove, newBoxes...)
	}
	if warehouseMap[nextL.Y][nextL.X] == "]" {
		boxesToMove = append(boxesToMove, Box{pos: nextL, side: "]"})
		boxesToMove = append(boxesToMove, Box{pos: nextL.Add(DIRECTIONS["left"]), side: "["})
		newBoxes, nvalid := findConnectedBoxes(warehouseMap, nextL.Add(DIRECTIONS["left"]), nextL, move)
		if !nvalid {
			return []Box{}, false
		}
		boxesToMove = append(boxesToMove, newBoxes...)

	}
	if warehouseMap[nextR.Y][nextR.X] == "[" {
		boxesToMove = append(boxesToMove, Box{pos: nextR, side: "["})
		boxesToMove = append(boxesToMove, Box{pos: nextR.Add(DIRECTIONS["right"]), side: "]"})
		newBoxes, nvalid := findConnectedBoxes(warehouseMap, nextR, nextR.Add(DIRECTIONS["right"]), move)
		if !nvalid {
			return []Box{}, false
		}
		boxesToMove = append(boxesToMove, newBoxes...)

	}

	return boxesToMove, valid
}
func draw2DArray(screen tcell.Screen, grid [][]string, robot image.Point) {
	for y, row := range grid {
		for x, char := range row {
			if x == robot.X && y == robot.Y {
				drawCell(screen, x, y, '@', tcell.ColorLightSkyBlue) // Draw the robot
			} else if char == "#" {
				drawCell(screen, x, y, rune(char[0]), tcell.ColorForestGreen)
			} else if char == "[" || char == "]" {
				drawCell(screen, x, y, rune(char[0]), tcell.ColorViolet)
			} else if char == "." {
				drawCell(screen, x, y, rune(char[0]), tcell.ColorWhite)
			}
		}
	}
}

func drawCell(screen tcell.Screen, x, y int, char rune, color tcell.Color) {
	style := tcell.StyleDefault.Foreground(color)
	screen.SetContent(x, y, char, nil, style)
}

func part2(warehouseMap [][]string, moves []image.Point, rPos image.Point) int {
	screen, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	}
	err = screen.Init()
	if err != nil {
		panic(err)
	}
	defer screen.Fini()
	print2DArray(warehouseMap, rPos)
	for _, move := range moves {
		// fmt.Print("\033[H\033[2J")
		// fmt.Println()
		// print2DArray(warehouseMap, rPos)
		// fmt.Println()
		screen.Clear()
		time.Sleep(200 * time.Millisecond)
		draw2DArray(screen, warehouseMap, rPos)
		screen.Show()
		newPos := rPos.Add(move)
		itemInFront := warehouseMap[newPos.Y][newPos.X]

		// free space
		if itemInFront == "." {
			rPos = newPos
			continue
		}
		// wall
		if itemInFront == "#" {
			continue
		}
		// box
		if itemInFront == "[" || itemInFront == "]" {
			// moving horizontal
			var boxesToMove []Box
			boxesToMove = append(boxesToMove, Box{pos: newPos, side: itemInFront})
			if move == DIRECTIONS["right"] || move == DIRECTIONS["left"] {
				next := newPos
				for {
					next = next.Add(move)
					if !inBounds(warehouseMap, next) || warehouseMap[next.Y][next.X] == "#" {
						break
					}
					if warehouseMap[next.Y][next.X] == "[" || warehouseMap[next.Y][next.X] == "]" {
						boxesToMove = append(boxesToMove, Box{pos: next, side: warehouseMap[next.Y][next.X]})
					}
					if warehouseMap[next.Y][next.X] == "." {
						rPos = newPos
						for _, box := range boxesToMove {
							newBoxPos := box.pos.Add(move)
							warehouseMap[newBoxPos.Y][newBoxPos.X] = box.side
						}
						warehouseMap[newPos.Y][newPos.X] = "."
						break
					}
				}
			} else {
				var leftPos, rightPos image.Point
				if itemInFront == "[" {
					leftPos = newPos
					rightPos = newPos.Add(DIRECTIONS["right"])
				} else {
					rightPos = newPos
					leftPos = newPos.Add(DIRECTIONS["left"])
				}
				boxesToMove = append(boxesToMove, Box{pos: leftPos, side: warehouseMap[leftPos.Y][leftPos.X]})
				boxesToMove = append(boxesToMove, Box{pos: rightPos, side: warehouseMap[rightPos.Y][rightPos.X]})
				newBoxes, valid := findConnectedBoxes(warehouseMap, leftPos, rightPos, move)
				boxesToMove = append(boxesToMove, newBoxes...)
				if valid {
					writtenTo := make(map[image.Point]bool)
					rPos = newPos
					for _, box := range boxesToMove {
						newBoxPos := box.pos.Add(move)
						warehouseMap[newBoxPos.Y][newBoxPos.X] = box.side
						writtenTo[newBoxPos] = true
						// if old position hasn't been written to before
						if !writtenTo[box.pos] {
							warehouseMap[box.pos.Y][box.pos.X] = "."
						}
					}
					warehouseMap[leftPos.Y][leftPos.X] = "."
					warehouseMap[rightPos.Y][rightPos.X] = "."
					continue
				}
			}
		}
	}
	print2DArray(warehouseMap, rPos)
	sumCoords := 0
	for y, row := range warehouseMap {
		for x, char := range row {
			if char == "[" {
				sumCoords += (y * 100) + x
			}
		}
	}

	return sumCoords
}
