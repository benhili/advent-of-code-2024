package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)
	file, _ := os.Open("input.txt")
	reports := parse_input(file)

	if part == 1 {
		ans := part1(reports)
		fmt.Println("Output:", ans)
	} else {
		ans := part2(reports)
		fmt.Println("Output:", ans)
	}
}

func parse_input(f *os.File) [][]int {
	scanner := bufio.NewScanner(f)
	var reports [][]int
	for scanner.Scan() {
		var line []int
		splitLines := strings.Split(scanner.Text(), " ")
		for _, v := range splitLines {
			number, err := strconv.Atoi(v)
			if err != nil {
				panic("Couldn't parse number")
			}
			line = append(line, number)
		}
		reports = append(reports, line)
	}
	return reports
}

func isReportSafe(report []int) bool {
	increase, decrease := false, false
	for i := 1; i < len(report); i++ {
		difference := report[i] - report[i-1]

		if difference > 0 {
			increase = true
		} else if difference < 0 {
			decrease = true
		} else {
			return false
		}

		if increase && decrease {
			return false
		}

		difference = max(difference, -difference)
		if difference > 3 || difference < 0 {
			return false
		}
	}
	return true
}

func part1(reports [][]int) int {
	var sumOfSafeReports int

	for _, report := range reports {
		if isReportSafe(report) {
			sumOfSafeReports = sumOfSafeReports + 1
		}
	}

	return sumOfSafeReports
}

func checkWithDeletion(report []int, deleteIndex int) bool {
	cloneReport := make([]int, len(report))
	_ = copy(cloneReport, report)
	// delete with slice fuckery
	// fmt.Println(deleteIndex)
	if deleteIndex == len(cloneReport)-1 {
		cloneReport = cloneReport[:deleteIndex]
	} else {
		cloneReport = append(cloneReport[:deleteIndex], cloneReport[deleteIndex+1:]...)
	}
	return isReportSafe(cloneReport)
}

func bruteForceMachine(report []int) bool {
	for i := 0; i < len(report); i++ {
		// grab a slice
		if checkWithDeletion(report, i) {
			return true
		}
	}
	return false
}

func part2(reports [][]int) int {
	var sumOfSafeReports int

	for _, report := range reports {
		if bruteForceMachine(report) {
			sumOfSafeReports = sumOfSafeReports + 1
		}
	}

	return sumOfSafeReports
}
