package main

import (
	"fmt"
)

// printVarType принимает переменную типа interface{} и выводит ее тип, используя форматированный вывод.
func printVarType(variable interface{}) {
	// Получаем строковое представление типа переменной с использованием форматированной строки.
	varType := fmt.Sprintf("%T", variable)

	// Выводим тип переменной.
	fmt.Printf("Тип переменной: %s\n", varType)
}

func main() {
	// Инициализация переменных разных типов
	var integerVar int = 42
	var stringVar string = "Hello, Gophers!"
	var boolVar bool = true
	var channelVar chan int = make(chan int)

	// Вывод типов переменных с использованием функции printVarType
	printVarType(integerVar)
	printVarType(stringVar)
	printVarType(boolVar)
	printVarType(channelVar)
}
