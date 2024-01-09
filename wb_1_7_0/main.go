package main

import (
	"fmt"
	"sync"
)

func main() {
	// Создаем мьютекс для защиты доступа к map
	var mutex sync.Mutex

	// Создаем map для хранения данных
	data := make(map[string]int)

	// Функция для конкурентной записи данных в map
	writeData := func(key string, value int, wg *sync.WaitGroup) {
		defer wg.Done()

		// Захватываем мьютекс перед записью данных
		mutex.Lock()
		defer mutex.Unlock()

		// Записываем данные в map
		data[key] = value

		fmt.Printf("Data written: %s -> %d\n", key, value)
	}

	// Используем WaitGroup для ожидания завершения всех горутин
	var wg sync.WaitGroup

	// Запускаем горутины для конкурентной записи данных
	for i := 0; i < 5; i++ {
		key := fmt.Sprintf("Key%d", i)
		value := i * 10

		wg.Add(1)
		go writeData(key, value, &wg)
	}

	// Ждем завершения всех горутин
	wg.Wait()

	// Выводим окончательное состояние map
	fmt.Println("Final data in map:", data)
}
