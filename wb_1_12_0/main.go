package main

import "fmt"

func main() {
	// Последовательность строк
	words := []string{"cat", "cat", "dog", "cat", "tree"}

	// Создаем множество
	wordSet := make(map[string]bool)

	// Добавляем уникальные элементы в множество
	for _, word := range words {
		wordSet[word] = true
	}

	// Выводим уникальные элементы множества
	fmt.Println("Множество:", wordSet)

	// Печать каждого элемента множества
	fmt.Println("Элементы множества:")
	for word := range wordSet {
		fmt.Println(word)
	}
}
