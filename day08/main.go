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
	file, _ := os.ReadFile("example.txt")
	antennaMap := parse(file)

	if part == 1 {
		ans := part1(antennaMap)
		fmt.Println("Output:", ans)
	} else {
		// ans := part2(antennaMap)
		// fmt.Println("Output:", ans)
	}
}

func parse(file []byte) [][]string {
	var antennaMap [][]string

	for _, line := range strings.Split(string(file), "\n") {
		antennaMap = append(antennaMap, strings.Split(line, ""))
	}
	return antennaMap
}

func inBounds(yLimit, xLimit, y, x int) bool {
	if y < yLimit && x < xLimit && x > -1 && y > -1 {
		return true
	}
	return false
}

func getAtennas(antennaMap [][]string) map[string][]coords {
	antennas := make(map[string][]coords)
	for y, row := range antennaMap {
		for x, symbol := range row {
			if symbol != "." {
				antennas[symbol] = append(antennas[symbol], coords{x: x, y: y})
			}
		}
	}
	return antennas
}

func calculatePairAntinodeCoords(a1, a2 coords) (coords, coords) {
	antinode1 := coords{x: -999, y: -999}
	antinode2 := coords{x: -999, y: -999}
	xDist := max(a1.x-a2.x, a2.x-a1.x)
	yDist := max(a1.y-a2.y, a2.y-a1.y)
	// diagonal (x and y are different)
	if a1.y < a2.y {
		// a1 higher
		antinode1.y = a1.y - yDist
		antinode2.y = a2.y + yDist
	} else if a2.y < a1.y {
		// a2 higher
		antinode1.y = a1.y + yDist
		antinode2.y = a2.y - yDist
	} else {
		// same height
		antinode1.y = a1.y
		antinode2.y = a2.y
	}

	if a1.x < a2.x {
		// a1 leftmost
		antinode1.x = a1.x - xDist
		antinode2.x = a2.x + xDist
	} else if a2.x < a1.x {
		// a2 leftmost
		antinode1.x = a1.x + xDist
		antinode2.x = a2.x - xDist
	} else {
		// same horizontal
		antinode1.x = a1.x
		antinode2.x = a2.x
	}

	if antinode1.x == -999 || antinode1.y == -999 || antinode2.x == -999 || antinode2.y == -999 {
		panic("SOMETHINGS WRONG")
	}

	return antinode1, antinode2
}

// turns coords into key in shape "y,x"
func coordsToKey(coord coords) string {
	return strconv.Itoa(coord.y) + "," + strconv.Itoa(coord.x)
}

func part1(antennaMap [][]string) int {
	antennas := getAtennas(antennaMap)
	sum := 0
	seen := make(map[string]bool)

	for k := range antennas {
		coordList := antennas[k]
		if len(coordList) < 2 {
			// need at least 2 antennas to make an antinode
			continue
		}

		for i := 0; i < len(coordList); i++ {
			for j := i + 1; j < len(coordList); j++ {
				a1 := coordList[i]
				a2 := coordList[j]
				antinode1, antinode2 := calculatePairAntinodeCoords(a1, a2)
				if inBounds(len(antennaMap), len(antennaMap[0]), antinode1.y, antinode1.x) {
					strKey := coordsToKey(antinode1)
					if !(seen[strKey]) {
						sum += 1
						seen[strKey] = true
					}
				}
				if inBounds(len(antennaMap), len(antennaMap[0]), antinode2.y, antinode2.x) {
					strKey := coordsToKey(antinode2)
					if !(seen[strKey]) {
						sum += 1
						seen[strKey] = true
					}
				}
			}
		}
	}
	return sum
}
