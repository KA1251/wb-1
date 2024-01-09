package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	// Создаем atomic.Value для хранения данных
	var data atomic.Value
	data.Store(make(map[string]int))

	// Используем WaitGroup для ожидания завершения всех горутин
	var wg sync.WaitGroup

	// Функция для конкурентной записи данных в map
	writeData := func(key string, value int) {
		defer wg.Done()

		// Загружаем текущий map из atomic.Value
		currentData := data.Load().(map[string]int)

		// Копируем текущий map, чтобы избежать гонок данных
		newData := make(map[string]int)
		for k, v := range currentData {
			newData[k] = v
		}

		// Записываем новые данные в map
		newData[key] = value

		// Сохраняем обновленный map в atomic.Value
		data.Store(newData)
		fmt.Printf("Data written: %s -> %d\n", key, value)
	}

	// Запускаем горутины для конкурентной записи данных
	for i := 0; i < 5; i++ {
		key := fmt.Sprintf("Key%d", i)
		value := i * 10

		wg.Add(1)
		go writeData(key, value)
	}

	// Ожидаем завершения всех горутин
	wg.Wait()

	// Получаем окончательное состояние map из atomic.Value
	finalData := data.Load().(map[string]int)

	// Выводим окончательное состояние map
	fmt.Println("Final data in map:", finalData)
}
