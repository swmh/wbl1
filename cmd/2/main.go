package main

import (
	"fmt"
	"sync"
)

func squares(arr []int) {
	wg := sync.WaitGroup{}
	wg.Add(len(arr))

	for _, v := range arr {
		v := v
		go func() {
			defer wg.Done()
			fmt.Println(v*v)
			
		}()
	}

	wg.Wait()
	
}

// Написать программу, которая конкурентно рассчитает значение квадратов чисел
// взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
func main() {
	arr := []int{2, 4, 6, 8, 10}
	squares(arr)

}
