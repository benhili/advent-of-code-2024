package utils

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func Setup() (int, []byte, time.Time) {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Start")
	fmt.Println("------------------")
	fmt.Println("Running part", part)
	start := time.Now()
	file, _ := os.ReadFile("input.txt")
	return part, file, start
}
func Finish(start time.Time) {
	elapsed := time.Since(start)
	fmt.Println("------------------")
	fmt.Printf("Ran in %s", elapsed)
}
