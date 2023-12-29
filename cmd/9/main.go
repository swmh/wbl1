package main

import "fmt"

func conveyor(ch <-chan int, do func(v int)) {
	for v := range ch {
		do(v)
	}
}

// Разработать конвейер чисел.
// Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2,
// после чего данные из второго канала должны выводиться в stdout.
func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	numbersCh := make(chan int)

	go func() {
		for _, v := range numbers {
			numbersCh <- v
		}

		close(numbersCh)
	}()

	squareNumbersCh := make(chan int)

	go func() {
		defer close(squareNumbersCh)

		conveyor(numbersCh, func(v int) {
			squareNumbersCh <- v * 2
		})
	}()

	conveyor(squareNumbersCh, func(v int) {
		fmt.Println(v)
	})
}
