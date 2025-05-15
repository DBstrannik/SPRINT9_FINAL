package main

import (
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name string
		size int
		want int
	}{
		{"size 0", 0, 0},
		{"size 10", 10, 10},
		{"size 100", 100, 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := generateRandomElements(tt.size)
			if len(data) != tt.want {
				t.Errorf("Expected slice length %d, got %d", tt.want, len(data))
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"empty slice", []int{}, 0},
		{"single element", []int{42}, 42},
		{"multiple elements", []int{1, 3, 2, 5, 4}, 5},
		{"all equal", []int{7, 7, 7, 7}, 7},
		{"negative numbers", []int{-1, -5, -3}, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximum(tt.data)
			if got != tt.want {
				t.Errorf("Expected %d, got %d", tt.want, got)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"empty slice", []int{}, 0},
		{"single element", []int{42}, 42},
		{"multiple elements", []int{1, 5, 3, 8, 7, 4, 6, 9}, 9},
		{"size less than chunks", []int{1, 2}, 2},
		{"large size", generateRandomElements(1000), maximum(generateRandomElements(1000))},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxChunks(tt.data)
			if got != tt.want {
				t.Errorf("Expected %d, got %d", tt.want, got)
			}
		})
	}
}
