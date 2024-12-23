package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/benhili/advent-of-code-2024/utils"
)

func main() {
	part, file, start := utils.Setup("input.txt")
	var secretNums []int
	for _, strNum := range strings.Split(string(file), "\n") {
		num, _ := strconv.Atoi(strNum)
		secretNums = append(secretNums, num)
	}
	if part == 1 {
		ans := part1(secretNums)
		fmt.Println(ans)
	} else {
		ans := part2(secretNums)
		fmt.Println(ans)
	}
	utils.Finish(start)
}

func prune(value int) int {
	return value % 16777216
}

func mix(secret, value int) int {
	return value ^ secret
}

func evolveSecret(secret int) int {
	secret = prune(mix(secret, secret*64))
	secret = prune(mix(secret, secret/32))
	secret = prune(mix(secret, secret*2048))
	return secret
}

func ones(num int) int {
	return num % 10
}

func part1(secretNums []int) int {
	sum := 0
	for _, initialSecret := range secretNums {
		evolvedSecret := initialSecret
		for i := 0; i < 2000; i++ {
			evolvedSecret = evolveSecret(evolvedSecret)
		}
		sum += evolvedSecret
	}
	return sum
}

type Stock struct {
	value      int
	difference int
}

func part2(secretNums []int) int {
	max := 0
	var windows []map[string]int
	possibleKeys := make(map[string]bool)
	for _, initialSecret := range secretNums {
		var secretDiffs []Stock
		evolvedSecret := initialSecret
		for i := 0; i < 2000; i++ {
			prevSecret := evolvedSecret
			evolvedSecret = evolveSecret(evolvedSecret)
			secretDiffs = append(secretDiffs, Stock{difference: ones(evolvedSecret) - ones(prevSecret), value: ones(evolvedSecret)})
		}
		window := make(map[string]int)
		for i := 4; i < len(secretDiffs); i++ {
			var key string
			for j := i - 4; j < i; j++ {
				key += (strconv.Itoa(secretDiffs[j].difference))
				if j != i-1 {
					key += ","
				}
			}
			if _, has := window[key]; !has {
				window[key] = secretDiffs[i-1].value
				possibleKeys[key] = true
			}
		}
		windows = append(windows, window)
	}

	for key := range possibleKeys {
		var bananasum int
		for _, window := range windows {
			bananasum += window[key]
		}
		max = int(math.Max(float64(bananasum), float64(max)))
	}
	return max
}
