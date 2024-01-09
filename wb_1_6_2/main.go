package main

import (
	"fmt"
	"time"
)

// worker - функция, представляющая рабочую горутину
// stopCh: Канал, используемый для получения сигнала остановки работы горутины.
func worker(stopCh chan struct{}) {
	for {
		select {
		case <-stopCh:
			// При получении сигнала остановки, выводим сообщение и завершаем выполнение
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
	stopCh := make(chan struct{})

	// Запуск горутины
	go worker(stopCh)

	// Ждем какое-то время, а затем отправляем сигнал остановки
	time.Sleep(3 * time.Second)
	close(stopCh)

	// Ждем завершения горутины
	time.Sleep(time.Second)
	fmt.Println("Main: Program finished.")
}
