package main

import (
	"fmt"
	"runtime"
	"time"
)

// worker - функция, представляющая рабочую горутину.
// stopCh: Канал только для чтения, используемый для получения сигнала остановки работы горутины.
func worker(stopCh <-chan struct{}) {
	for {
		select {
		case <-stopCh:
			// При получении сигнала остановки, выводим сообщение, завершаем выполнение и выходим из текущей горутины
			fmt.Println("Worker: Received stop signal. Exiting...")
			runtime.Goexit()
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

	fmt.Println("Main: Program finished.")
}
