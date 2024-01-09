package main

import (
	"fmt"
	"time"
)

func main() {
	var N int
	fmt.Println("введите количество воркеров")
	fmt.Scanf("%d", &N)
	dataChannel := make(chan int) // Канал для передачи данных

	// Горутина для записи данных в канал
	go func() {
		for i := 1; ; i++ {
			dataChannel <- i
			time.Sleep(time.Second) // Моделируем задержку
		}
	}()

	// Горутина для чтения данных из канала
	go func() {
		for {
			select {
			case data := <-dataChannel:
				fmt.Println("Received:", data)
			case <-time.After(time.Second): // Таймер для задержки
			}
		}
	}()

	// Ждем N секунд
	time.Sleep(time.Duration(N) * time.Second)

	// Завершаем программу
	fmt.Println("Program finished.")
}
