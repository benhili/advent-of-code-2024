package main

import (
	"os"
	"testing"
)

func Test_part1(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     int
	}{
		{"example", "example.txt", 143},
		{"input", "input.txt", 5166},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, _ := os.ReadFile(tt.filename)
			orderingRules, updates := parse(file)

			if got := part1(orderingRules, updates); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     int
	}{
		{"example", "example.txt", 123},
		{"input", "input.txt", 4679},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, _ := os.ReadFile(tt.filename)

			orderingRules, updates := parse(file)
			if got := part2(orderingRules, updates); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
