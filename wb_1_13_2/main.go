package main

import "fmt"

func main() {
	a := 5
	b := 10

	fmt.Printf("До обмена: a = %d, b = %d\n", a, b)

	// Простая замена значений без использования временной переменной
	a, b = b, a

	fmt.Printf("После обмена: a = %d, b = %d\n", a, b)
}
