package main

import (
	"fmt"
	"sync"
	"time"
)

// Структура счетчика
type Counter struct {
	value int
	mu    sync.Mutex
}

// Метод для инкрементации счетчика
func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// Метод для получения текущего значения счетчика
func (c *Counter) GetValue() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main() {
	// Создаем экземпляр счетчика
	counter := Counter{}

	// Создаем WaitGroup для ожидания завершения горутин
	var wg sync.WaitGroup

	// Количество горутин для инкрементации
	numGoroutines := 10

	// Добавляем горутины в WaitGroup
	wg.Add(numGoroutines)

	// Запускаем горутины
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			// Имитируем какую-то работу
			time.Sleep(time.Millisecond * 100)
			// Инкрементируем счетчик
			counter.Increment()
		}()
	}

	// Ожидаем завершения всех горутин
	wg.Wait()

	// Получаем и выводим итоговое значение счетчика
	finalValue := counter.GetValue()
	fmt.Printf("Итоговое значение счетчика: %d\n", finalValue)
}
