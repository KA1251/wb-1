package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	ch := make(chan int)
	var wg sync.WaitGroup

	// Горутина для обработки чисел из массива
	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, num := range numbers {
			ch <- num
		}
		close(ch)
	}()

	// Горутина для умножения чисел и вывода результатов
	go func() {
		for num := range ch {
			result := num * 2
			fmt.Println(result)
		}
	}()

	wg.Wait()
}
