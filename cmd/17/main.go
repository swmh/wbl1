package main

import (
	"cmp"
	"fmt"
	"slices"
)

func BinarySearch[S ~[]E, E cmp.Ordered](x S, target E) (int, bool) {
	l := 0
	r := len(x) - 1

	for l <= r {
		m := (l + r) / 2

		if x[m] == target {
			return m, true
		}

		if x[m] < target {
			l = m + 1
			continue
		}

		r = m - 1

	}
	return 0, false
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	target := 99

	idx, ok := slices.BinarySearch(arr, target)
	fmt.Printf("idx: %d, ok: %t\n", idx, ok)

	idx, ok = BinarySearch(arr, target)
	fmt.Printf("idx: %d, ok: %t\n", idx, ok)
}
