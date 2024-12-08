package main

import (
	"flag"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type calibration struct {
	target int
	nums   []int
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)
	file, _ := os.ReadFile("input.txt")
	calibrations := parse(file)

	if part == 1 {
		ans := part1(calibrations)
		fmt.Println("Output:", ans)
	} else {
		ans := part2(calibrations)
		fmt.Println("Output:", ans)
	}
}

func parse(file []byte) []calibration {
	var calibrations []calibration

	for _, line := range strings.Split(string(file), "\n") {
		strNums := strings.Split(line, " ")
		target, _ := strconv.Atoi(strNums[0][:len(strNums[0])-1])
		var nums []int
		for _, v := range strNums[1:] {
			num, _ := strconv.Atoi(v)
			nums = append(nums, num)
		}
		calibrations = append(calibrations, calibration{target: target, nums: nums})
	}
	return calibrations
}

// 292: 11 6 16 20

// 11 + 6 = 17
// 11 * 6 = 66

// [17, 66]

// 17 + 16 = 33
// 17 * 16 = 272
// 66 * 16 = 1056
// 66 + 16 = 82

// [33, 272, 1056, 82] remove larger than target? nah fuck it

// 33, 272, 1056, 82

func part1(calibrations []calibration) int {
	sum := 0
	for _, calibration := range calibrations {
		i := 1
		var combinations []int
		for i < len(calibration.nums) {
			if len(combinations) == 0 {
				combinations = append(combinations, calibration.nums[i-1]+calibration.nums[i])
				combinations = append(combinations, calibration.nums[i-1]*calibration.nums[i])
			} else {
				var newCombinations []int
				for _, v := range combinations {
					newCombinations = append(newCombinations, v+calibration.nums[i])
					newCombinations = append(newCombinations, v*calibration.nums[i])
				}
				combinations = newCombinations
			}
			i += 1
		}
		if slices.Contains(combinations, calibration.target) {
			sum += calibration.target
		}
	}
	return sum
}

func part2(calibrations []calibration) int {
	sum := 0
	for _, calibration := range calibrations {
		i := 1
		var combinations []int
		for i < len(calibration.nums) {
			if len(combinations) == 0 {
				combinations = append(combinations, calibration.nums[i-1]+calibration.nums[i])
				combinations = append(combinations, calibration.nums[i-1]*calibration.nums[i])
				concat, _ := strconv.Atoi(fmt.Sprintf("%d%d", calibration.nums[i-1], calibration.nums[i]))
				combinations = append(combinations, concat)
			} else {
				var newCombinations []int
				for _, v := range combinations {
					newCombinations = append(newCombinations, v+calibration.nums[i])
					newCombinations = append(newCombinations, v*calibration.nums[i])
					// concat
					concat, _ := strconv.Atoi(fmt.Sprintf("%d%d", v, calibration.nums[i]))
					newCombinations = append(newCombinations, concat)
				}
				combinations = newCombinations
			}
			i += 1
		}
		if slices.Contains(combinations, calibration.target) {
			sum += calibration.target
		}
	}
	return sum
}
