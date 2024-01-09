package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// WorkerData представляет данные, которые обрабатывают воркеры.
type WorkerData struct {
	ID   int
	Data int
}

// worker читает данные из канала и выводит их в stdout.
func worker(workerData *WorkerData, wg *sync.WaitGroup, mu *sync.Mutex, quit chan struct{}) {
	defer wg.Done() // Уменьшаем счетчик WaitGroup при завершении воркера.

	for {
		select {
		case <-quit:
			// Если получен сигнал завершения, выходим из функции
			return
		default:
			// Захватываем мьютекс для безопасного доступа к данным.
			mu.Lock()
			data := workerData.Data
			workerData.Data++
			mu.Unlock()

			fmt.Printf("Worker %d: Received data %d\n", workerData.ID, data)
			time.Sleep(time.Second) // Моделируем обработку данных
		}
	}
}

func main() {
	var numWorkers int // Количество воркеров (может быть изменено)
	fmt.Println("введите количество воркеров")
	fmt.Scanf("%d", &numWorkers)
	var wg sync.WaitGroup
	var mu sync.Mutex
	quit := make(chan struct{}) // Канал для уведомления воркеров о завершении

	// Обработчик сигналов для Ctrl+C
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	// Запуск воркеров
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		workerData := &WorkerData{ID: i}
		go worker(workerData, &wg, &mu, quit)
	}

	// Ожидание сигнала завершения
	<-sigCh
	fmt.Println("Received termination signal. Stopping workers...")

	// Закрытие канала для уведомления воркеров о завершении
	close(quit)

	// Ждем завершения работы воркеров
	wg.Wait()
	fmt.Println("All workers stopped.")
}
