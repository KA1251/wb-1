package main

import "fmt"

// printVarType принимает переменную типа interface{} и выводит ее тип.
func printVarType(variable interface{}) {
	switch variable.(type) {
	case int:
		fmt.Println("Тип переменной: int")
	case string:
		fmt.Println("Тип переменной: string")
	case bool:
		fmt.Println("Тип переменной: bool")
	case chan int:
		fmt.Println("Тип переменной: chan int")
	default:
		fmt.Println("Неизвестный тип переменной")
	}
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
