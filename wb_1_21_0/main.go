package main

import (
	"fmt"
)

// OldGraphicsLibrary представляет старую графическую библиотеку с устаревшим интерфейсом.
type OldGraphicsLibrary struct {
}

// DrawLine рисует линию в старой графической библиотеке.
func (o *OldGraphicsLibrary) DrawLine(x1, y1, x2, y2 int) {
	fmt.Printf("OldLibrary: DrawLine from (%d, %d) to (%d, %d)\n", x1, y1, x2, y2)
}

// NewGraphicsLibrary представляет новую графическую библиотеку с современным интерфейсом.
type NewGraphicsLibrary interface {
	Draw(x1, y1, x2, y2 int)
}

// OldLibraryAdapter - адаптер, который позволяет использовать OldGraphicsLibrary как NewGraphicsLibrary.
type OldLibraryAdapter struct {
	OldLibrary *OldGraphicsLibrary
}

// Draw рисует линию в новой графической библиотеке, используя адаптер.
func (a *OldLibraryAdapter) Draw(x1, y1, x2, y2 int) {
	// Просто делегируем вызов старой библиотеке через адаптер
	a.OldLibrary.DrawLine(x1, y1, x2, y2)
}

func main() {
	// Создаем экземпляр старой графической библиотеки
	oldLibrary := &OldGraphicsLibrary{}

	// Создаем адаптер для использования OldGraphicsLibrary как NewGraphicsLibrary
	adapter := &OldLibraryAdapter{OldLibrary: oldLibrary}

	// Используем адаптер в контексте, предназначенном для NewGraphicsLibrary
	useNewGraphicsLibrary(adapter)
}

// useNewGraphicsLibrary использует NewGraphicsLibrary для рисования линии.
func useNewGraphicsLibrary(library NewGraphicsLibrary) {
	library.Draw(0, 0, 100, 100)
}
