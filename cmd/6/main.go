package main

import (
	"context"
	"fmt"
	"time"
)

func first(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println("first")
			time.Sleep(time.Second)
		}
	}
}

func second(stopCh <-chan struct{}) {
	for {
		select {
		case <-stopCh:
			return
		default:
			fmt.Println("second")
			time.Sleep(time.Second)
		}
	}
}

// Реализовать все возможные способы остановки выполнения горутины.
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	go first(ctx)

	stopCh := make(chan struct{})

	go second(stopCh)

	time.Sleep(8 * time.Second)
	stopCh <- struct{}{}
}
