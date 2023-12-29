package main

import (
	"fmt"
	"slices"
)

func reverse(s string) string {
	runes := []rune(s)
	slices.Reverse(runes)

	return string(runes)
}

// Разработать программу, которая переворачивает подаваемую
// на ход строку (например: «главрыба — абырвалг»).
// Символы могут быть unicode.
func main() {
	input := "главрыба"

	fmt.Println(reverse(input))
}
