package main

import (
	"fmt"
	"slices"
)

// Удалить i-ый элемент из слайса.
func main() {
	a := []int{4, 45, 0, 499, -12}

	i := 2
	result := slices.Delete(a, i, i+1)

	fmt.Printf("%v\n", result)

}
