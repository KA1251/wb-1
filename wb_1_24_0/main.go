package main

import (
	"fmt"
	"math"
)

// Point представляет координаты точки в двумерном пространстве.
type Point struct {
	x, y float64
}

// NewPoint создает новый экземпляр точки с заданными координатами.
func NewPoint(x, y float64) Point {
	return Point{x, y}
}

// Distance вычисляет расстояние между двумя точками.
func Distance(p1, p2 Point) float64 {
	dx := p2.x - p1.x
	dy := p2.y - p1.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	// Создаем две точки
	point1 := NewPoint(1, 2)
	point2 := NewPoint(4, 6)

	// Вычисляем расстояние между точками
	distance := Distance(point1, point2)

	// Выводим результат
	fmt.Printf("Точка 1: (%.2f, %.2f)\n", point1.x, point1.y)
	fmt.Printf("Точка 2: (%.2f, %.2f)\n", point2.x, point2.y)
	fmt.Printf("Расстояние между точками: %.2f\n", distance)
}
