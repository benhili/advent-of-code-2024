package utils

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func Setup(fileName string) (int, []byte, time.Time) {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("------------------")
	fmt.Println("Running part", part)
	fmt.Println("------------------")
	start := time.Now()
	file, _ := os.ReadFile(fileName)
	return part, file, start
}
func Finish(start time.Time) {
	elapsed := time.Since(start)
	fmt.Println("------------------")
	fmt.Printf("Ran in %s", elapsed)
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
