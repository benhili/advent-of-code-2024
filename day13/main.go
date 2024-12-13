package main

import (
	"fmt"
	"image"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/benhili/advent-of-code-2024/utils"
)

type game struct {
	a     image.Point
	b     image.Point
	prize image.Point
}

func main() {
	part, file, start := utils.Setup("input.txt")
	games := parse(string(file))
	if part == 1 {
		ans := part1(games)
		fmt.Println("Output:", ans)
	} else {
		ans := part2(games)
		fmt.Println("Output:", ans)
	}
	utils.Finish(start)
}

func unsafeAtoi(str string) int {
	num, _ := strconv.Atoi(str)

	return num
}

func parse(file string) []game {
	var games []game
	for _, gameString := range strings.Split(file, "\n\n") {
		// fmt.Println("game", i)
		// fmt.Println(gameString)
		var newGame game
		r, _ := regexp.Compile(`X.(\d+), Y.(\d+)`)
		matches := r.FindAllStringSubmatch(gameString, -1)
		newGame.a.X = unsafeAtoi(matches[0][1])
		newGame.a.Y = unsafeAtoi(matches[0][2])
		newGame.b.X = unsafeAtoi(matches[1][1])
		newGame.b.Y = unsafeAtoi(matches[1][2])
		newGame.prize.X = unsafeAtoi(matches[2][1])
		newGame.prize.Y = unsafeAtoi(matches[2][2])
		games = append(games, newGame)
	}
	return games
}

func dfs(g game, point image.Point, cost int, aPressCount, bPressCount int, visited map[image.Point]int) (bool, int) {
	if g.prize == point {
		return true, cost
	}
	if prevCost, seen := visited[point]; seen && cost >= prevCost {
		return false, 0
	}
	visited[point] = cost
	aFound, aCost := dfs(g, point.Add(g.a), cost+3, aPressCount+1, bPressCount, visited)
	bFound, bCost := dfs(g, point.Add(g.b), cost+1, aPressCount, bPressCount+1, visited)

	if aFound && bFound {
		return true, min(aCost, bCost)
	} else if aFound {
		return true, aCost
	} else if bFound {
		return true, bCost
	}
	return false, 0
}

func part1(games []game) int {
	cost := 0

	for _, game := range games {
		// brute force
		visited := make(map[image.Point]int)
		found, cheapestCost := dfs(game, image.Point{0, 0}, 0, 0, 0, visited)
		if found {
			cost += cheapestCost
		}
	}
	return cost
}

func getPriceByEquation(g game, delta int) int {
	/*
		2x2 System Linear Equations
		a * aX + b * bX = pX
		a * aY + b * bY = pY
	*/
	pX := g.prize.X + delta
	pY := g.prize.Y + delta
	aX := g.a.X
	aY := g.a.Y
	bX := g.b.X
	bY := g.b.Y

	a := float64(pX*bY-pY*bX) / float64(aX*bY-aY*bX)
	b := float64(pY*aX-pX*aY) / float64(aX*bY-aY*bX)

	// if there is no decimals is valid
	if a == math.Trunc(a) && b == math.Trunc(b) {
		return int(a*3 + b)
	}
	return 0
}

func part2(games []game) int {
	cost := 0

	for _, game := range games {
		// maths 🤓
		cost += getPriceByEquation(game, 10000000000000)
	}
	return cost
}
