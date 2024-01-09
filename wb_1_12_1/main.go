package main

import "fmt"

func main() {
	// Последовательность строк
	words := []string{"cat", "cat", "dog", "cat", "tree"}

	// Создаем map для хранения уникальных элементов
	wordMap := make(map[string]struct{})

	// Добавляем уникальные элементы в map
	for _, word := range words {
		wordMap[word] = struct{}{}
	}

	// Выводим уникальные элементы множества
	fmt.Println("Множество:", wordMap)

	// Печать каждого элемента множества
	fmt.Println("Элементы множества:")
	for item := range wordMap {
		fmt.Println(item)
	}
}
