package main

import (
	"fmt"
	"image"
	"strings"

	"github.com/benhili/advent-of-code-2024/utils"
)

func main() {
	part, _, start := utils.Setup("example.txt")

	if part == 1 {
		ans := part1()
		fmt.Println(ans)
	} else {
		// ans := part1(racetrack, sPos, ePos, 141)
		// fmt.Println(ans)
	}
	utils.Finish(start)
}

var Numpad = [][]string{
	{"7", "8", "9"},
	{"4", "5", "6"},
	{"1", "2", "3"},
	{" ", "0", "A"}}

var NumpadMap = func() map[string]image.Point {
	temp := make(map[string]image.Point)
	for y, row := range Numpad {
		for x, char := range row {
			temp[char] = image.Point{x, y}
		}
	}
	return temp
}()
var DirectionalPad = [][]string{
	{" ", "^", "A"},
	{"<", "v", ">"}}

var DPadMap = func() map[string]image.Point {
	temp := make(map[string]image.Point)
	for y, row := range DirectionalPad {
		for x, char := range row {
			temp[char] = image.Point{x, y}
		}
	}
	return temp
}()

func distanceToButton(pos image.Point, char string, buttonMap map[string]image.Point) (int, int) {
	dest := buttonMap[char]
	return dest.X - pos.X, dest.Y - pos.Y
}

func distToMoves(xDist, yDist int) []string {
	var moves []string
	if yDist > 0 {
		// down
		for range yDist {
			moves = append(moves, "v")
		}
	}

	if xDist < 0 {
		// left
		for range -xDist {
			moves = append(moves, "<")
		}
	}

	if xDist > 0 {
		// right
		for range xDist {
			moves = append(moves, ">")
		}
	}

	if yDist < 0 {
		// up
		for range -yDist {
			moves = append(moves, "^")
		}
	}

	return moves
}

func part1() int {
	sequence := "029A"
	numPos := image.Point{2, 3}
	dPadPos := image.Point{2, 0}
	dPad2Pos := image.Point{2, 0}

	var dpad1Moves []string
	var dpad2Moves []string
	// var humanMoves []string
	for _, char := range strings.Split(sequence, "") {
		// numpad
		numPadXDist, numPadYDist := distanceToButton(numPos, char, NumpadMap)
		numpadMoves := distToMoves(numPadXDist, numPadYDist)
		// dpad 1
		fmt.Println("numpad", numpadMoves)
		for _, move := range numpadMoves {
			if DirectionalPad[dPadPos.Y][dPadPos.X] == move {
				dpad2Moves = append(dpad1Moves, "A")
				continue
			} else if DirectionalPad[dPadPos.Y][dPadPos.X] != "A" {
				returnDistX, returnDistY := distanceToButton(dPadPos, "A", DPadMap)
				dpad1Moves = append(dpad1Moves, distToMoves(returnDistX, returnDistY)...)
				dpad1Moves = append(dpad1Moves, "A")
				dPadPos = dPadPos.Add(image.Point{returnDistX, returnDistY})
			}
			dPadXDist, dPadYDist := distanceToButton(dPadPos, move, DPadMap)

			dpad1Moves = append(dpad1Moves, distToMoves(dPadXDist, dPadYDist)...)
			fmt.Println("adding moves", distToMoves(dPadXDist, dPadYDist))

			// press the button
			dpad1Moves = append(dpad1Moves, "A")
			// update the pointer
			dPadPos = dPadPos.Add(image.Point{dPadXDist, dPadYDist})

		}
		returnDistX, returnDistY := distanceToButton(dPadPos, "A", DPadMap)
		dpad1Moves = append(dpad1Moves, distToMoves(returnDistX, returnDistY)...)
		dpad1Moves = append(dpad1Moves, "A")
		dPadPos = dPadPos.Add(image.Point{returnDistX, returnDistY})
		numPos = numPos.Add(image.Point{numPadXDist, numPadYDist})
	}

	for _, move := range dpad1Moves {
		if DirectionalPad[dPad2Pos.Y][dPad2Pos.X] == move {
			dpad2Moves = append(dpad2Moves, "A")
			continue
		} else if DirectionalPad[dPad2Pos.Y][dPad2Pos.X] != "A" {
			returnDistX, returnDistY := distanceToButton(dPad2Pos, "A", DPadMap)
			dpad2Moves = append(dpad2Moves, distToMoves(returnDistX, returnDistY)...)
			dpad2Moves = append(dpad2Moves, "A")
			dPad2Pos = dPad2Pos.Add(image.Point{returnDistX, returnDistY})
		}
		dPadXDist, dPadYDist := distanceToButton(dPad2Pos, move, DPadMap)

		dpad2Moves = append(dpad2Moves, distToMoves(dPadXDist, dPadYDist)...)
		fmt.Println("adding moves", distToMoves(dPadXDist, dPadYDist))

		// press the button
		dpad2Moves = append(dpad2Moves, "A")
		// update the pointer
		dPad2Pos = dPad2Pos.Add(image.Point{dPadXDist, dPadYDist})

	}
	fmt.Println(dpad2Moves)

	return len(dpad1Moves)
}
