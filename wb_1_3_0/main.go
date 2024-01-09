package main

import (
	"fmt"
	"sync"
)

// calculateSquareSum рассчитывает сумму квадратов чисел в заданной последовательности.
// Использует канал resultChan для передачи частичных результатов.
func calculateSquareSum1(wg *sync.WaitGroup, number int, resultChan chan int) {
	defer wg.Done() // Уменьшаем счетчик WaitGroup при завершении горутины.
	sum := 0
	sum += number * number

	resultChan <- sum // Отправляем частичную сумму в канал.
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	numChunks := len(numbers)

	var wg sync.WaitGroup
	resultChan := make(chan int, numChunks) // Канал для передачи частичных результатов.

	// Делим последовательность чисел на части и вычисляем сумму квадратов конкурентно.
	for i := 0; i < len(numbers); i++ {
		wg.Add(1)
		go calculateSquareSum1(&wg, numbers[i], resultChan)
	}

	// Закрываем канал после завершения всех горутин.
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Суммируем результаты из канала.
	totalSum := 0
	for partialSum := range resultChan {
		totalSum += partialSum
	}

	fmt.Printf("Сумма квадратов: %d\n", totalSum)
}
