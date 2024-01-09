package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Создаем два больших числа
	a := new(big.Int)
	b := new(big.Int)

	// Устанавливаем значения a и b (значения > 2^20)
	a.SetString("1000000", 10) // Пример: 1000000
	b.SetString("2000000", 10) // Пример: 2000000

	// Перемножение
	resultMult := new(big.Int)
	resultMult.Mul(a, b)
	fmt.Printf("Умножение: %s\n", resultMult.String())

	// Деление
	resultDiv := new(big.Int)
	resultDiv.Div(a, b)
	fmt.Printf("Деление: %s\n", resultDiv.String())

	// Сложение
	resultAdd := new(big.Int)
	resultAdd.Add(a, b)
	fmt.Printf("Сложение: %s\n", resultAdd.String())

	// Вычитание
	resultSub := new(big.Int)
	resultSub.Sub(a, b)
	fmt.Printf("Вычитание: %s\n", resultSub.String())
}
