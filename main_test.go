package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	for _, size := range []int{-1, 0, 1, 100} {
		data := generateRandomElements(size)
		wantLength := size
		if size < 0 {
			wantLength = 0
		}
		if len(data) != wantLength {
			t.Fatalf("generateRandomElements(%d): длина %d, ожидается %d", size, len(data), wantLength)
		}
		for _, value := range data {
			if value <= 0 || value > maxRandomValue {
				t.Errorf("generateRandomElements(%d): число %d вне диапазона [1, %d]", size, value, maxRandomValue)
			}
		}
	}
}

func TestMaximumFunctions(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{name: "nil", data: nil, want: 0},
		{name: "пустой слайс", data: []int{}, want: 0},
		{name: "один элемент", data: []int{42}, want: 42},
		{name: "максимум в начале", data: []int{9, 3, 1}, want: 9},
		{name: "максимум в середине", data: []int{1, 9, 3}, want: 9},
		{name: "максимум в конце", data: []int{1, 3, 9}, want: 9},
		{name: "одинаковые элементы", data: []int{5, 5, 5}, want: 5},
		{name: "нулевые элементы", data: []int{0, 0, 0}, want: 0},
		{name: "отрицательные элементы", data: []int{-5, -1, -3}, want: -1},
		{name: "ровно восемь элементов", data: []int{1, 2, 3, 4, 5, 6, 7, 8}, want: 8},
		{name: "максимум в остатке", data: []int{1, 2, 3, 4, 5, 6, 7, 8, 99}, want: 99},
		{name: "отрицательные элементы с остатком", data: []int{-9, -8, -7, -6, -5, -4, -3, -2, -1}, want: -1},
	}
	functions := []struct {
		name string
		find func([]int) int
	}{
		{name: "maximum", find: maximum},
		{name: "maxChunks", find: maxChunks},
	}
	for _, function := range functions {
		t.Run(function.name, func(t *testing.T) {
			for _, test := range tests {
				t.Run(test.name, func(t *testing.T) {
					if got := function.find(test.data); got != test.want {
						t.Errorf("получено %d, ожидается %d", got, test.want)
					}
				})
			}
		})
	}
}

func TestMaxChunksBoundaries(t *testing.T) {
	for position := 0; position < 83; position++ {
		data := make([]int, 83)
		data[position] = 10
		if got := maxChunks(data); got != 10 {
			t.Errorf("максимум на позиции %d: получено %d, ожидается 10", position, got)
		}
	}
}
