package main

import "fmt"

func Intersection(set1, set2 map[int]struct{}) map[int]struct{} {
	result := make(map[int]struct{})

	maxSet := set1
	minSet := set2

	if len(set1) < len(set2) {
		maxSet = set2
		minSet = set1
	}

	for k := range minSet {
		if _, ok := maxSet[k]; ok {
			result[k] = struct{}{}
		}
	}

	return result
}

// Реализовать пересечение двух неупорядоченных множеств.
func main() {
	set1 := map[int]struct{}{
		3:   {},
		-4:  {},
		2:   {},
		123: {},
		14:  {},
		5:   {},
		-99: {},
		0:   {},
		15:  {},
	}

	set2 := map[int]struct{}{
		14: {},
		8:  {},
		15: {},
		1:  {},
		5:  {},
		23: {},
		2:  {},
		74: {},
	}

	fmt.Printf("%v\n", Intersection(set1, set2))
}
