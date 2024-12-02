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
		{"example", "example.txt", 2},
		{"input", "input.txt", 371},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, _ := os.Open(tt.filename)
			reports := parse_input(file)

			if got := part1(reports); got != tt.want {
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
		{"example", "example.txt", 4},
		{"input", "input.txt", 426},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, _ := os.Open(tt.filename)
			reports := parse_input(file)

			if got := part2(reports); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
