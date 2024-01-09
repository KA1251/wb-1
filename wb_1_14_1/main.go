package main

import (
	"fmt"
	"reflect"
)

// printVarType принимает переменную типа interface{} и выводит ее тип, используя пакет reflect.
func printVarType(variable interface{}) {
	// Получаем тип переменной с использованием пакета reflect.
	varType := reflect.TypeOf(variable)

	// Выводим тип переменной.
	fmt.Printf("Тип переменной: %v\n", varType)
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
