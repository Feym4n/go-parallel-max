package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE           = 100_000_000
	CHUNKS         = 8
	maxRandomValue = 1_000_000
)

// generateRandomElements возвращает случайные положительные числа.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return nil
	}
	data := make([]int, size)
	for i := range data {
		data[i] = rand.Intn(maxRandomValue) + 1
	}
	return data
}

// maximum возвращает наибольший элемент или ноль для пустого слайса.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}
	max := data[0]
	for _, value := range data[1:] {
		if value > max {
			max = value
		}
	}
	return max
}

// maxChunks находит максимум, обрабатывая части слайса в восьми горутинах.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}
	maxima := make([]int, CHUNKS)
	chunkSize := len(data) / CHUNKS
	var wg sync.WaitGroup
	for i := 0; i < CHUNKS; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if i == CHUNKS-1 {
			end = len(data)
		}
		wg.Add(1)
		go func(index int, chunk []int) {
			defer wg.Done()
			maxima[index] = data[0]
			if len(chunk) > 0 {
				maxima[index] = maximum(chunk)
			}
		}(i, data[start:end])
	}
	wg.Wait()
	return maximum(maxima)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	max := maximum(data)
	elapsed := time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d горутинах\n", CHUNKS)
	start = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", max, elapsed)
}
