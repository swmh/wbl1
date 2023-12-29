package main

import "fmt"

// TODO:
func Group(t []float64, step int) map[int][]float64 {
	result := make(map[int][]float64)
	for _, v := range t {
		a := int(v / float64(step))
		fmt.Println("a", a)
		result[step*a] = append(result[step*a], v)
	}

	return result
}

// Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
// Объединить данные значения в группы с шагом в 10 градусов.
// Последовательность в подмножноствах не важна.
//
// Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.
func main() {
	arr := []float64{-29.9, -30, -20, -19.9, -25.4, 0, -9.9, -10, 1.2, -1.5, 9.9, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	result := Group(arr, 10)
	fmt.Printf("%v\n", result)
}
