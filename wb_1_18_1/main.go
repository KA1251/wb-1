package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// Структура счетчика
type Counter struct {
	value int64
}

// Метод для инкрементации счетчика
func (c *Counter) Increment() {
	atomic.AddInt64(&c.value, 1)
}

// Метод для получения текущего значения счетчика
func (c *Counter) GetValue() int64 {
	return atomic.LoadInt64(&c.value)
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
