package main

import (
	"fmt"
	"sync"
	"time"
)

// worker - функция, представляющая рабочую горутину
func worker(stopCh <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done() // Уменьшаем счетчик ожидания при выходе из функции

	// Бесконечный цикл для выполнения рабочей нагрузки
	for {
		select {
		case <-stopCh:
			// При получении сигнала остановки, выводим сообщение и завершаем функцию
			fmt.Println("Worker: Received stop signal. Exiting...")
			return
		default:
			// Рабочая нагрузка горутины
			fmt.Println("Worker: Doing some work...")
			time.Sleep(time.Second) // Имитация выполнения работы
		}
	}
}

func main() {
	stopCh := make(chan struct{}) // Создаем канал для отправки сигнала остановки
	var wg sync.WaitGroup         // Создаем WaitGroup для ожидания завершения горутины

	// Запуск горутины
	wg.Add(1) // Увеличиваем счетчик ожидания
	go worker(stopCh, &wg)

	// Ждем какое-то время, а затем отправляем сигнал остановки
	time.Sleep(3 * time.Second)
	close(stopCh) // Закрываем канал, чтобы отправить сигнал остановки горутины

	// Ждем завершения горутины
	wg.Wait() // Блокируем выполнение до завершения горутины
	fmt.Println("Main: Program finished.")
}
