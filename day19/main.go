package main

import (
	"fmt"
	"strings"

	"github.com/benhili/advent-of-code-2024/utils"
)

func main() {
	part, file, start := utils.Setup("input.txt")

	avail, desired := parse(string(file))
	if part == 1 {
		ans := part1(avail, desired)
		fmt.Println(ans)
	} else {
		ans := part2(avail, desired)
		fmt.Println(ans)
	}
	utils.Finish(start)
}

func parse(input string) ([]string, []string) {
	splitString := strings.Split(input, "\n\n")
	return strings.Split(splitString[0], ", "), strings.Split(splitString[1], "\n")
}

func search(avail []string, pattern string, cache map[string]bool) bool {
	if pattern == "" {
		return true
	}
	if cachedRes, seen := cache[pattern]; seen {
		return cachedRes
	}

	var found bool
	for _, availablePattern := range avail {
		if strings.HasPrefix(pattern, availablePattern) {
			found = search(avail, pattern[len(availablePattern):], cache)
			if found {
				cache[pattern] = true
				return true
			}
		}
	}
	cache[pattern] = false
	return false
}

func part1(avail []string, desiredDesigns []string) int {
	count := 0
	cache := make(map[string]bool)
	for _, design := range desiredDesigns {
		if search(avail, design, cache) {
			count++
		}
	}
	return count
}

func searchAll(avail []string, pattern string, cache map[string]int) int {
	if pattern == "" {
		return 1
	}
	if count, seen := cache[pattern]; seen {
		return count
	}

	var total int
	for _, availablePattern := range avail {
		if strings.HasPrefix(pattern, availablePattern) {
			total += searchAll(avail, pattern[len(availablePattern):], cache)
		}
	}
	cache[pattern] = total
	return total
}

func part2(avail []string, desiredDesigns []string) int {
	count := 0
	cache := make(map[string]int)
	for _, design := range desiredDesigns {
		count += searchAll(avail, design, cache)
	}
	return count
}
