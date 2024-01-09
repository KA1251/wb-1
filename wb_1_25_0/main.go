package main

import (
	"fmt"
	"time"
)

// Sleep1 реализует функцию sleep с использованием time.Sleep.
func Sleep1(seconds int) {
	time.Sleep(time.Duration(seconds) * time.Second)
}

func main() {
	fmt.Println("Before sleep")
	Sleep1(3) // Пауза в три секунды
	fmt.Println("After sleep")
}
