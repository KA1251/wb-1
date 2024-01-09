package main

import (
	"fmt"
	"strings"
)

// Функция для переворачивания слов в строке
func reverseWords(input string) string {
	// Разбиваем строку на слова
	words := strings.Fields(input)

	// Переворачиваем порядок слов
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	// Собираем слова обратно в строку
	reversedString := strings.Join(words, " ")

	return reversedString
}

func main() {
	// Пример строки
	inputString := "snow dog sun mega_word!"

	// Вызываем функцию для переворачивания слов
	reversedWords := reverseWords(inputString)

	// Выводим результат
	fmt.Printf("Исходная строка: %s\n", inputString)
	fmt.Printf("Строка с перевернутыми словами: %s\n", reversedWords)
}
