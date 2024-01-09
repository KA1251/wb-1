package main

import (
	"fmt"
	"sync"
)

// Использует sync.WaitGroup для синхронизации горутин.
func calculateSquare(wg *sync.WaitGroup, num int) {
	defer wg.Done() // Уменьшаем счетчик WaitGroup при завершении горутины.

	result := num * num
	fmt.Printf("%d squared is %d\n", num, result)
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup // WaitGroup используется для ожидания завершения всех горутин.

	for _, num := range numbers {
		wg.Add(1)                    // Увеличиваем счетчик WaitGroup перед запуском каждой горутины.
		go calculateSquare(&wg, num) // Запускаем горутину для расчета квадрата числа.
	}

	wg.Wait() // Ожидаем завершения всех горутин перед завершением программы.
}
