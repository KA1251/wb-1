package main

import "fmt"

func removeBySlicing(slice []int, index int) []int {
	// Проверка на допустимость индекса
	if index < 0 || index >= len(slice) {
		return slice
	}

	// Изменяем оригинальный слайс, пропустив элемент
	result := append(slice[:index], slice[index+1:]...)

	return result
}

func main() {
	// Пример слайса
	originalSlice := []int{1, 2, 3, 4, 5}
	fmt.Printf("Исходный слайс: %v\n", originalSlice)
	// Удаляем элемент с индексом 2
	indexToRemove := 2
	modifiedSlice := removeBySlicing(originalSlice, indexToRemove)

	// Вывод результата

	fmt.Printf("Слайс после удаления элемента с индексом %d: %v\n", indexToRemove, modifiedSlice)
}
