package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// worker - функция, представляющая рабочую горутину
func worker(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			// При получении сигнала остановки (в результате отмены контекста), завершаем выполнение
			fmt.Println("Worker: Received stop signal. Exiting...")
			return
		default:
			// Рабочая нагрузка горутины
			fmt.Println("Worker: Doing some work...")
			time.Sleep(time.Second)
		}
	}
}

func main() {
	// Создаем контекст с функцией отмены
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	// Запуск горутины
	wg.Add(1)
	go worker(ctx, &wg)

	// Ждем какое-то время, а затем отменяем контекст
	time.Sleep(3 * time.Second)
	cancel()

	// Ждем завершения горутины
	wg.Wait()
	fmt.Println("Main: Program finished.")
}
