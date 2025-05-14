package main

import (
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	data := generateRandomElements(0)
	if len(data) != 0 {
		t.Errorf("Expected empty slice for size 0, got %d", len(data))
	}

	data = generateRandomElements(10)
	if len(data) != 10 {
		t.Errorf("Expected slice of length 10, got %d", len(data))
	}
}

func TestMaximum(t *testing.T) {
	max, err := maximum([]int{42})
	if err != nil || max != 42 {
		t.Error("Failed on single-element slice")
	}

	max, err = maximum([]int{1, 3, 2, 5, 4})
	if err != nil || max != 5 {
		t.Error("Incorrect maximum result")
	}

	_, err = maximum([]int{})
	if err == nil {
		t.Error("Expected error for empty slice")
	}
}

func TestMaxChunks(t *testing.T) {
	// Проверка на обычный массив
	data := []int{1, 5, 3, 8, 7, 4, 6, 9}
	max, err := maxChunks(data)
	if err != nil || max != 9 {
		t.Errorf("Expected 9, got %d", max)
	}

	// Проверка на пустой массив
	_, err = maxChunks([]int{})
	if err == nil {
		t.Error("Expected error for empty slice")
	}
}
