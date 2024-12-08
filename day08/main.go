package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type coords struct {
	x int
	y int
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)
	file, _ := os.ReadFile("input.txt")
	antennaCoordinates := parse(file)

	if part == 1 {
		ans := part1(antennaCoordinates)
		fmt.Println("Output:", ans)
	} else {
		ans := part2(antennaCoordinates)
		fmt.Println("Output:", ans)
	}
}

func parse(file []byte) map[string][]coords {
	antennas := make(map[string][]coords)

	for y, line := range strings.Split(string(file), "\n") {
		for x, symbol := range strings.Split(line, "") {
			if symbol != "." {
				antennas[symbol] = append(antennas[symbol], coords{x: x, y: y})
			}
		}
	}
	return antennas
}

func inBounds(yLimit, xLimit, y, x int) bool {
	if y < yLimit && x < xLimit && x > -1 && y > -1 {
		return true
	}
	return false
}

// turns coords into key in shape "y,x"
func coordsToKey(coord coords) string {
	return strconv.Itoa(coord.y) + "," + strconv.Itoa(coord.x)
}

func part1(antennaCoordinates map[string][]coords) int {
	seen := make(map[string]bool)
	yLimit, xLimit := 50, 50
	for k := range antennaCoordinates {
		coordList := antennaCoordinates[k]
		if len(coordList) < 2 {
			// need at least 2 antennas to make an antinode
			continue
		}

		for i := 0; i < len(coordList); i++ {
			for j := i + 1; j < len(coordList); j++ {
				a1 := coordList[i]
				a2 := coordList[j]
				antinode1 := coords{x: a1.x + (a1.x - a2.x), y: a1.y + (a1.y - a2.y)}
				antinode2 := coords{x: a2.x + (a2.x - a1.x), y: a2.y + (a2.y - a1.y)}
				if inBounds(yLimit, xLimit, antinode1.y, antinode1.x) {
					strKey := coordsToKey(antinode1)
					if !(seen[strKey]) {
						seen[strKey] = true
					}
				}
				if inBounds(yLimit, xLimit, antinode2.y, antinode2.x) {
					strKey := coordsToKey(antinode2)
					if !(seen[strKey]) {
						seen[strKey] = true
					}
				}
			}
		}
	}
	return len(seen)
}

func part2(antennaCoordinates map[string][]coords) int {
	seen := make(map[string]bool)
	yLimit, xLimit := 50, 50
	for k := range antennaCoordinates {
		coordList := antennaCoordinates[k]
		if len(coordList) < 2 {
			// need at least 2 antennas to make an antinode
			continue
		}

		for i := 0; i < len(coordList); i++ {
			for j := i + 1; j < len(coordList); j++ {
				a1 := coordList[i]
				a2 := coordList[j]
				seen[coordsToKey(a1)] = true
				seen[coordsToKey(a2)] = true
				var antinodes []coords
				xDiff1 := (a1.x - a2.x)
				yDiff1 := (a1.y - a2.y)
				xDiff2 := (a2.x - a1.x)
				yDiff2 := (a2.y - a1.y)

				for i := 0; i < yLimit; i++ {
					antinode1 := coords{x: a1.x + xDiff1, y: a1.y + yDiff1}
					antinode2 := coords{x: a2.x + xDiff2, y: a2.y + yDiff2}
					antinodes = append(antinodes, antinode1, antinode2)
					a1 = antinode1
					a2 = antinode2
				}

				for _, antinode := range antinodes {
					if inBounds(yLimit, xLimit, antinode.y, antinode.x) {
						strKey := coordsToKey(antinode)
						if !(seen[strKey]) {
							seen[strKey] = true
						}
					}
				}
			}
		}
	}
	return len(seen)
}
