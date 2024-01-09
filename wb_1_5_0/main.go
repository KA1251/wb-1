package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	var N int
	fmt.Println("введите количество воркеров")
	fmt.Scanf("%d", &N)

	var mu sync.Mutex // Mutex для защиты доступа к общему ресурсу (срезу data)
	var wg sync.WaitGroup

	data := make([]int, 0)
	dataChannel := make(chan int, 5) // Буферизованный канал для передачи данных от горутины записи к горутине чтения

	// Создаем контекст с таймаутом в N секунд
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(N)*time.Second)
	defer cancel()

	// Горутина для записи данных в канал
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; ; i++ {
			select {
			case <-ctx.Done():
				close(dataChannel) // Закрываем канал перед завершением горутины
				return             // Завершаем горутину при завершении контекста
			default:
				mu.Lock() // Блокируем доступ к общему ресурсу
				data = append(data, i)
				dataChannel <- i        // Отправляем данные в буферизованный канал
				mu.Unlock()             // Разблокируем доступ к общему ресурсу
				time.Sleep(time.Second) // Моделируем задержку
			}
		}
	}()

	// Горутина для чтения данных из канала
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return // Завершаем горутину при завершении контекста
			case d, ok := <-dataChannel:
				if !ok {
					return // Канал закрыт, завершаем горутину
				}
				fmt.Println("Received:", d)
				time.Sleep(time.Second) // Моделируем задержку
			}
		}
	}()

	// Ожидаем завершения всех горутин
	wg.Wait()
	fmt.Println("Program finished.")
}
