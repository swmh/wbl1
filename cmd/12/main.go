package main

import "fmt"

func MakeSet(arr []string) map[string]struct{} {
	result := make(map[string]struct{})

	for _, v := range arr {
		result[v] = struct{}{}
	}

	return result
}

// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Printf("%v\n", MakeSet(arr))
}
