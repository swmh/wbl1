package main

import (
	"fmt"
	"slices"
	"strings"
)

func reverse(s string) string {
	words := strings.Split(s, " ")
	slices.Reverse(words)

	return strings.Join(words, " ")
}

// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».
func main() {
	input := "snow dog sun"
	result := reverse(input)

	fmt.Printf("%s - %s", input, result)
}
