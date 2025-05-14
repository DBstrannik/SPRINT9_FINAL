package main

import (
	"errors"
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
		data[i] = rand.Intn(1_000_000)
	}
	return data
}

// maximum возвращает максимальное число из слайса или ошибку, если слайс пустой
func maximum(data []int) (int, error) {
	if len(data) == 0 {
		return 0, errors.New("slice is empty")
	}
	max := data[0]
	for _, v := range data[1:] {
		if v > max {
			max = v
		}
	}
	return max, nil
}

// maxChunks параллельно ищет максимум по чанкам, возвращает итоговый максимум или ошибку
func maxChunks(data []int) (int, error) {
	if len(data) == 0 {
		return 0, errors.New("slice is empty")
	}

	chunkSize := len(data) / CHUNKS
	maxResults := make([]int, CHUNKS)
	var wg sync.WaitGroup
	wg.Add(CHUNKS)

	for i := 0; i < CHUNKS; i++ {
		go func(i int) {
			defer wg.Done()
			start := i * chunkSize
			end := start + chunkSize
			if i == CHUNKS-1 {
				end = len(data)
			}
			max := data[start]
			for _, v := range data[start+1 : end] {
				if v > max {
					max = v
				}
			}
			maxResults[i] = max
		}(i)
	}

	wg.Wait()
	return maximum(maxResults)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	rand.Seed(time.Now().UnixNano())
	data := generateRandomElements(SIZE)

	// Последовательный поиск
	start := time.Now()
	max, err := maximum(data)
	elapsed := time.Since(start).Microseconds()
	if err != nil {
		fmt.Printf("Ошибка при поиске (один поток): %v\n", err)
	} else {
		fmt.Printf("Максимальное значение элемента: %d\nВремя поиска (один поток): %d мкс\n", max, elapsed)
	}

	// Параллельный поиск
	start = time.Now()
	max, err = maxChunks(data)
	elapsed = time.Since(start).Microseconds()
	if err != nil {
		fmt.Printf("Ошибка при поиске (%d потоков): %v\n", CHUNKS, err)
	} else {
		fmt.Printf("Максимальное значение элемента: %d\nВремя поиска (%d потоков): %d мкс\n", max, CHUNKS, elapsed)
	}
}
