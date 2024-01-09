package main

import "fmt"

// Функция бинарного поиска в отсортированном массиве.
// Возвращает индекс элемента, если он найден, и -1 в противном случае.
func binarySearch(arr []int, target int) int {
	low, high := 0, len(arr)-1

	for low <= high {
		mid := (low + high) / 2

		if arr[mid] == target {
			return mid // Элемент найден, возвращаем индекс
		} else if arr[mid] < target {
			low = mid + 1 // Искомый элемент во второй половине
		} else {
			high = mid - 1 // Искомый элемент в первой половине
		}
	}

	return -1 // Элемент не найден
}

func main() {
	// Пример отсортированного массива
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Искомое значение
	target := 5

	// Вызываем функцию бинарного поиска
	index := binarySearch(arr, target)

	// Проверяем результат
	if index != -1 {
		fmt.Printf("Элемент %d найден по индексу %d\n", target, index)
	} else {
		fmt.Printf("Элемент %d не найден\n", target)
	}
}
