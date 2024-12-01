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
		{"example", "example.txt", 11},
		{"input", "input.txt", 2000468},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, _ := os.Open(tt.filename)
			left, right := parse_input(file)

			if got := part1(left, right); got != tt.want {
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
		{"example", "example.txt", 31},
		{"input", "input.txt", 18567089},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, _ := os.Open(tt.filename)
			left, right := parse_input(file)

			if got := part2(left, right); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
