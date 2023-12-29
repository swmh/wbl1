package main

import (
	"fmt"
	"sync"
)

func squares(arr []int) int {
	wg := sync.WaitGroup{}
	wg.Add(len(arr))

	ch := make(chan int, len(arr))

	go func () {
		wg.Wait()
		close(ch)
	}()

	for _, v := range arr {
		v := v
		go func() {
			defer wg.Done()
			ch <- v * v
		}()
	}

	var sum int
	for v := range ch {
		sum += v
	}

	return sum
}

// Дана последовательность чисел: 2,4,6,8,10.
// Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.
func main() {
	arr := []int{2, 4, 6, 8, 10}
	fmt.Println(squares(arr))
}
