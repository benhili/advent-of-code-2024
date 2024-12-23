package main

import (
	"fmt"
	"slices"
	"sort"
	"strings"

	"github.com/benhili/advent-of-code-2024/utils"
)

func main() {
	part, file, start := utils.Setup("input.txt")
	connMap := parse(string(file))
	if part == 1 {
		ans := part1(connMap)
		fmt.Println(ans)
	} else {
		ans := part2(connMap)
		fmt.Println(ans)
	}
	utils.Finish(start)
}

func parse(input string) map[string][]string {
	connectionMap := make(map[string][]string)
	for _, line := range strings.Split(input, "\n") {
		splitLine := strings.Split(line, "-")
		leftConnection := splitLine[0]
		rightConnection := splitLine[1]
		connectionMap[leftConnection] = append(connectionMap[leftConnection], rightConnection)
		connectionMap[rightConnection] = append(connectionMap[rightConnection], leftConnection)
	}
	return connectionMap
}

func chiefComputerInList(a, b, c string) bool {
	return string(a[0]) == "t" || string(b[0]) == "t" || string(c[0]) == "t"
}

func part1(connMap map[string][]string) int {
	threes := make(map[string]bool)
	for name, connections := range connMap {
		for _, c1 := range connections {
			for _, c2 := range connMap[c1] {
				if slices.Contains(connections, c2) && chiefComputerInList(name, c1, c2) {
					connectionGroup := []string{name, c1, c2}
					sort.Strings(connectionGroup)
					threes[strings.Join(connectionGroup, ",")] = true
				}
			}
		}
	}

	return (len(threes))
}

func part2(connMap map[string][]string) int {
	var largest []string
	for name, connections := range connMap {
		lan := make([]string, 1)
		lan[0] = name
		for _, c1 := range connections {
			c1Connections := connMap[c1]
			valid := true
			for _, pc := range lan {
				// fmt.Println("checking if", c1Connections, "contains", pc, "current lan is", lan, "if successful will add", c1)
				if !slices.Contains(c1Connections, pc) {
					valid = false
					break
				}
			}
			if valid {
				lan = append(lan, c1)
			}
		}
		if len(lan) > len(largest) {
			largest = lan
		}
	}
	sort.Strings(largest)
	fmt.Println(strings.Join(largest, ","))
	return (len(largest))
}
