package main

import "fmt"

func removeByCopying(slice []int, index int) []int {
	// Проверка на допустимость индекса
	if index < 0 || index >= len(slice) {
		return slice
	}

	// Создаем новый слайс с элементами до удаляемого индекса
	result := append(slice[:index], slice[index+1:]...)

	return result
}

func main() {
	// Пример слайса
	originalSlice := []int{1, 2, 3, 4, 5}
	fmt.Printf("Исходный слайс: %v\n", originalSlice)
	// Удаляем элемент с индексом 2
	indexToRemove := 2
	modifiedSlice := removeByCopying(originalSlice, indexToRemove)

	// Вывод результата

	fmt.Printf("Слайс после удаления элемента с индексом %d: %v\n", indexToRemove, modifiedSlice)
}
