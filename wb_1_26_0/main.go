package main

import "fmt"

// IsUnique проверяет, что все символы в строке уникальны (регистронезависимо).
func IsUnique(s string) bool {
	// Создаем карту для отслеживания уникальных символов
	charMap := make(map[rune]bool)

	// Проходим по строке и проверяем уникальность каждого символа
	for _, char := range s {
		// Приводим символ к нижнему регистру
		lowerChar := toLowerCase(char)

		// Если символ уже встречался, строка не уникальна
		if charMap[lowerChar] {
			return false
		}

		// Добавляем символ в карту
		charMap[lowerChar] = true
	}

	// Все символы уникальны
	return true
}

// toLowerCase приводит символ к нижнему регистру
func toLowerCase(char rune) rune {
	if char >= 'A' && char <= 'Z' {
		return char + 'a' - 'A'
	}
	return char
}

func main() {
	// Примеры строк
	strings := []string{"abcd", "abCdefAaf", "aabcd", "AAAa"}

	// Проверяем уникальность символов в каждой строке
	for _, s := range strings {
		result := IsUnique(s)
		fmt.Printf("%s — %t\n", s, result)
	}
}
