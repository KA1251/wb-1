package main

import (
	"fmt"
)

// Функция для обмена элементов в массиве по индексам
func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

// Функция быстрой сортировки
func quickSort(arr []int, low, high int) {
	if low < high {
		// Индекс опорного элемента после разделения
		pivotIndex := partition(arr, low, high)

		// Рекурсивно сортируем подмассивы до и после опорного элемента
		quickSort(arr, low, pivotIndex-1)
		quickSort(arr, pivotIndex+1, high)
	}
}

// Функция для разделения массива на две части относительно опорного элемента
func partition(arr []int, low, high int) int {
	// Используем последний элемент в качестве опорного
	pivot := arr[high]

	// Индекс для элемента, который будет помещен на свое место
	i := low - 1

	// Проходим по массиву и перемещаем элементы меньше опорного перед ним, а большие - после
	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			swap(arr, i, j)
		}
	}

	// Помещаем опорный элемент на свое место
	swap(arr, i+1, high)

	// Возвращаем индекс опорного элемента
	return i + 1
}

func main() {
	// Пример неотсортированного массива
	arr := []int{5, 2, 9, 1, 5, 6}

	// Применяем быструю сортировку
	quickSort(arr, 0, len(arr)-1)

	// Выводим отсортированный массив
	fmt.Println(arr)
}
