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
		{"example", "example.txt", 161},
		{"input", "input.txt", 166905464},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, _ := os.ReadFile(tt.filename)

			if got := part1(string(file)); got != tt.want {
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
		{"example", "example.txt", 161},
		{"example", "example2.txt", 48},
		{"input", "input.txt", 72948684},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, _ := os.ReadFile(tt.filename)

			if got := part2(string(file)); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
