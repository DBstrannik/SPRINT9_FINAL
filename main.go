package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements генерирует случайные положительные числа
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}
	data := make([]int, size)
	for i := range data {
		data[i] = rand.Int()
	}
	return data
}

// maximum возвращает максимальное число из слайса
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}
	max := data[0]
	for _, v := range data[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

// maxChunks параллельно ищет максимум по чанкам, возвращает итоговый максимум
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	chunkSize := len(data) / CHUNKS
	maxResults := make([]int, CHUNKS)
	var wg sync.WaitGroup
	wg.Add(CHUNKS)

	for i := 0; i < CHUNKS; i++ {
		go func(i int, chunk []int) {
			defer wg.Done()
			maxResults[i] = maximum(chunk)
		}(i, data[i*chunkSize:min((i+1)*chunkSize, len(data))])
	}

	wg.Wait()
	return maximum(maxResults)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	rand.Seed(time.Now().UnixNano())
	data := generateRandomElements(SIZE)

	// Последовательный поиск
	start := time.Now()
	max := maximum(data)
	elapsed := time.Since(start).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска (один поток): %d мкс\n", max, elapsed)

	// Параллельный поиск
	start = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(start).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска (%d потоков): %d мкс\n", max, CHUNKS, elapsed)
}
