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
		{"example", "example.txt", 18},
		{"input", "input.txt", 2613},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, _ := os.ReadFile(tt.filename)
			input := parse(file)

			if got := part1(input); got != tt.want {
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
		{"example", "example.txt", 9},
		{"input", "input.txt", 1905},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, _ := os.ReadFile(tt.filename)
			input := parse(file)
			if got := part2(input); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
