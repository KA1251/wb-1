package main

import (
	"fmt"
	"sort"
)

func groupTemperatures(temperatures []float64, step float64) map[int][]float64 {
	groups := make(map[int][]float64)

	// Сортируем температуры
	sort.Float64s(temperatures)

	// Группируем температуры по шагу
	for _, temp := range temperatures {
		group := int(temp/step) * int(step)
		groups[group] = append(groups[group], temp)
	}

	return groups
}

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	step := 10.0

	groups := groupTemperatures(temperatures, step)

	// Выводим результат
	for key, values := range groups {
		fmt.Printf("%d: %v\n", key, values)
	}
}
