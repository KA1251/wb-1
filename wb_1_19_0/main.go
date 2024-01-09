package main

import (
	"fmt"
)

// Функция для переворачивания строки с учетом символов Unicode
func reverseString(input string) string {
	// Преобразуем строку в слайс рун
	runes := []rune(input)

	// Получаем длину строки
	length := len(runes)

	// Индексы начала и конца для переворачивания
	i, j := 0, length-1

	// Переворачиваем строку
	for i < j {
		// Меняем местами символы
		runes[i], runes[j] = runes[j], runes[i]
		// Увеличиваем начальный индекс, уменьшаем конечный индекс
		i++
		j--
	}

	// Возвращаем перевернутую строку
	return string(runes)
}

func main() {
	// Пример строки с символами Unicode
	inputString := "главрыба"

	// Вызываем функцию для переворачивания строки
	reversedString := reverseString(inputString)

	// Выводим результат
	fmt.Printf("Исходная строка: %s\n", inputString)
	fmt.Printf("Перевернутая строка: %s\n", reversedString)
}
