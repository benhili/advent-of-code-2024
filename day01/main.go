package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)
	file, _ := os.Open("input.txt")
	left, right := parse_input(file)

	if part == 1 {
		ans := part1(left, right)
		fmt.Println("Output:", ans)
	} else {
		ans := part2(left, right)
		fmt.Println("Output:", ans)
	}
}

func parse_input(f *os.File) ([]int, []int) {
	scanner := bufio.NewScanner(f)
	var left []int
	var right []int
	r, _ := regexp.Compile(`(\d+)\s+(\d+)`)
	for scanner.Scan() {
		line := scanner.Text()
		matches := r.FindAllStringSubmatch(line, -1)
		l, _ := strconv.Atoi(matches[0][1])
		r, _ := strconv.Atoi(matches[0][2])
		left = append(left, l)
		right = append(right, r)
	}
	return left, right
}

func part1(left, right []int) int {
	if len(left) != len(right) {
		panic("left and right are not the same length!")
	}

	sort.Ints(left)
	sort.Ints(right)

	var sumDifferences int

	for i := range left {
		difference := left[i] - right[i]
		sumDifferences = sumDifferences + max(difference, -difference)
	}

	return sumDifferences
}

func part2(left, right []int) int {
	var similarity int
	freqMap := make(map[int]int)

	for _, v := range right {
		freqMap[v] = freqMap[v] + 1
	}

	for _, v := range left {
		freq, ok := freqMap[v]
		if ok {
			similarity = similarity + freq*v
		}
	}

	return similarity
}
