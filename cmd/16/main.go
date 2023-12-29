package main

import (
	"cmp"
	"fmt"
	"slices"
)

func partition[T cmp.Ordered](x []T, l int, r int) int {
	idx := r
	pivot := x[idx]
	i := l - 1

	for j := l; j < idx; j++ {
		if x[j] <= pivot {
			i++
			x[i], x[j] = x[j], x[i]
		}
	}

	x[i+1], x[r] = x[r], x[i+1]
	return i + 1
}

func quickSort[T cmp.Ordered](x []T, l int, r int) {
	if l < r {
		q := partition(x, l, r)
		// _ = q
		quickSort(x, l, q-1)
		quickSort(x, q+1, r)
	}
}

func QuickSort[T cmp.Ordered](x []T) {
	quickSort(x, 0, len(x)-1)
}

// Реализовать быструю сортировку массива (quicksort) встроенными методами языка
func main() {
	x := []int{1, 5, 2, 6, 3, 7, 4, 8}
	fmt.Printf("%v\n", x)

	quickSort(x, 0, len(x)-1)
	fmt.Printf("%v\n", x)

	//

	x = []int{1, 5, 2, 6, 3, 7, 4, 8}
	fmt.Printf("%v\n", x)

	slices.Sort(x)
	fmt.Printf("%v\n", x)
}
