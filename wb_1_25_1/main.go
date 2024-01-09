package main

import (
	"fmt"
	"time"
)

// Sleep2 реализует функцию sleep с использованием канала и горутины.
func Sleep2(seconds int) {
	done := make(chan bool)
	go func() {
		time.Sleep(time.Duration(seconds) * time.Second)
		done <- true
	}()
	<-done
}

func main() {
	fmt.Println("Before sleep")
	Sleep2(3) // Пауза в три секунды
	fmt.Println("After sleep")
}
