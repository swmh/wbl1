package main

import "fmt"

// Поменять местами два числа без создания временной переменной.
func main() {
	a := 6
	b := 3
	fmt.Printf("a: %d, b: %d\n", a, b)

	a, b = b, a
	fmt.Printf("a: %d, b: %d\n", a, b)



}
