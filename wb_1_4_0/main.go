package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// worker читает данные из канала и выводит их в stdout.
func worker(workerID int, dataChannel <-chan int, wg *sync.WaitGroup, quit chan struct{}) {
	defer wg.Done() // Уменьшаем счетчик WaitGroup при завершении воркера.

	for {
		select {
		case data := <-dataChannel:
			fmt.Printf("Worker %d: Received data %d\n", workerID, data)
			time.Sleep(time.Second) // Моделируем обработку данных
		case <-quit:
			// Если получен сигнал завершения, выходим из функции
			return
		}
	}
}

func main() {
	var numWorkers int
	fmt.Println("введите количество воркеров")
	fmt.Scanf("%d", &numWorkers)

	dataChannel := make(chan int) // Канал для передачи данных от главного потока к воркерам
	quit := make(chan struct{})   // Канал для уведомления воркеров о завершении
	var wg sync.WaitGroup

	// Обработчик сигналов для Ctrl+C
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	// Запуск воркеров
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, dataChannel, &wg, quit)
	}

	// Постоянная запись данных в канал из главного потока
	go func() {
		for i := 1; ; i++ {
			select {
			case dataChannel <- i:
				// Моделируем создание новых данных
			case <-quit:
				// Если получен сигнал завершения, выходим из цикла
				return
			}
			time.Sleep(time.Second)
		}
	}()

	// Ожидание сигнала завершения
	<-sigCh
	fmt.Println("Received termination signal. Stopping workers...")

	// Закрытие канала для уведомления воркеров о завершении
	close(quit)

	// Ждем завершения работы воркеров
	wg.Wait()
	fmt.Println("All workers stopped.")
}
