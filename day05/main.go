package main

import (
	"flag"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)
	file, _ := os.ReadFile("input.txt")
	orderingRules, updates := parse(file)

	if part == 1 {
		ans := part1(orderingRules, updates)
		fmt.Println("Output:", ans)
	} else {
		ans := part2(orderingRules, updates)
		fmt.Println("Output:", ans)
	}
}

func parse(file []byte) (map[int][]int, [][]int) {
	var updates [][]int
	orderingRules := make(map[int][]int)
	splitFile := strings.Split(string(file), "\n\n")
	partOne, partTwo := splitFile[0], splitFile[1]
	for _, rule := range strings.Split(partOne, "\n") {
		splitRule := strings.Split(rule, "|")
		left, err := strconv.Atoi(splitRule[0])
		if err != nil {
			panic("Failed to parse left side of rule")
		}
		right, err := strconv.Atoi(splitRule[1])
		if err != nil {
			panic("Failed to parse right side of rule")
		}

		orderingRules[left] = append(orderingRules[left], right)
	}

	for _, update := range strings.Split(partTwo, "\n") {
		strNumbers := strings.Split(update, ",")
		var numbers []int
		for _, strNumber := range strNumbers {
			number, _ := strconv.Atoi(strNumber)
			numbers = append(numbers, number)
		}
		updates = append(updates, numbers)
	}

	return orderingRules, updates
}

func isValidUpdate(orderingRules map[int][]int, update []int) bool {
	seen := make(map[int]bool)
	for _, num := range update {
		mustNotBeBeforeList := orderingRules[num]
		for _, mustNotBeBefore := range mustNotBeBeforeList {
			if seen[mustNotBeBefore] {
				return false
			}
		}
		seen[num] = true
	}
	return true
}

func part1(orderingRules map[int][]int, updates [][]int) int {
	var sum int
	for _, update := range updates {
		if isValidUpdate(orderingRules, update) {
			middle := update[len(update)/2]
			sum += middle
		}
	}
	return sum
}

func sortUpdate(orderingRules map[int][]int, update []int) []int {
	sort.Slice(update, func(i, j int) bool {
		cantBeBeforeList := orderingRules[update[i]]
		return slices.Contains(cantBeBeforeList, update[j])
	})
	return update
}

func part2(orderingRules map[int][]int, updates [][]int) int {
	var sum int
	var invalidUpdates [][]int
	for _, update := range updates {
		if !isValidUpdate(orderingRules, update) {
			invalidUpdates = append(invalidUpdates, update)
		}
	}
	for _, invalidUpdate := range invalidUpdates {

		sortedUpdate := sortUpdate(orderingRules, invalidUpdate)
		middle := sortedUpdate[len(sortedUpdate)/2]
		sum += middle
	}
	return sum
}
