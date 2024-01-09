package main

import (
	"fmt"
	"sync"
	"time"
)

// Структура счетчика
type Counter struct {
	mu    sync.Mutex
	value int
}

// Метод для инкрементации счетчика
func (c *Counter) Increment(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
	ch <- c.value
}

func main() {
	// Создаем экземпляр счетчика
	counter := Counter{}

	// Количество горутин для инкрементации
	numGoroutines := 10

	// Создаем канал для передачи результатов
	resultCh := make(chan int, numGoroutines)

	// Создаем WaitGroup для дожидания завершения всех горутин
	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	// Запускаем горутины
	for i := 0; i < numGoroutines; i++ {
		go func() {
			// Имитируем какую-то работу
			time.Sleep(time.Millisecond * 100)
			// Инкрементируем счетчик и передаем результат в канал
			counter.Increment(resultCh, &wg)
		}()
	}

	// Закрываем канал после завершения всех горутин
	go func() {
		wg.Wait()
		close(resultCh)
	}()

	// Собираем результаты из канала
	var finalValue int
	for value := range resultCh {
		finalValue = value
	}

	// Выводим итоговое значение счетчика
	fmt.Printf("Итоговое значение счетчика: %d\n", finalValue)
}
