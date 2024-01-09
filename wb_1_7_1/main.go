package main

import (
	"fmt"
	"sync"
)

// Запись данных в map
func writeData(data map[string]int, key string, value int, wg *sync.WaitGroup, ch chan struct{}) {
	data[key] = value
	ch <- struct{}{} // Отправляем сигнал об окончании записи данных
	fmt.Printf("Data written: %s -> %d\n", key, value)
	wg.Done()
}

func main() {
	// Создаем map для хранения данных
	data := make(map[string]int)

	// Используем WaitGroup для ожидания завершения всех горутин
	var wg sync.WaitGroup

	// Канал для сигнала об окончании записи данных
	ch := make(chan struct{})

	// Запускаем горутины для конкурентной записи данных
	for i := 0; i < 5; i++ {
		key := fmt.Sprintf("Key%d", i)
		value := i * 10

		wg.Add(1)
		go writeData(data, key, value, &wg, ch)

		// Ожидаем сигнала об окончании записи данных
		<-ch
	}

	// Ожидаем завершения всех горутин
	wg.Wait()

	// Выводим окончательное состояние map
	fmt.Println("Final data in map:", data)
}
