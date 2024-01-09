package main

import "fmt"

func setBit(value int64, position uint, bitValue bool) int64 {
	// Создаем битовую маску
	mask := int64(1) << position

	if bitValue {
		// Устанавливаем i-й бит в 1
		return value | mask
	}

	// Устанавливаем i-й бит в 0
	return value &^ mask
}
func main() {
	var num int64 = 42
	var position uint
	var bitValue bool
	var answer string

	fmt.Println("Введите позицию бита и его значение")
	fmt.Scanf("%d", &position)

	// Считываем символ новой строки после первого ввода
	fmt.Scanln()

	fmt.Println("Введите значение бита (0 или 1)")
	fmt.Scanf("%s", &answer)

	if answer == "1" {
		bitValue = true
	}
	if answer == "0" {
		bitValue = false
	}

	result := setBit(num, position, bitValue)
	fmt.Printf("Установлен i-й бит в %v: %d\n", bitValue, result)
}
