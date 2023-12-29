package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Реализовать постоянную запись данных в канал (главный поток).
// Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
// Необходима возможность выбора количества воркеров при старте.
//
// Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.

func Worker(ch <-chan any, n int) {
	for v := range ch {
		fmt.Printf("worker %d: %v\n", n, v)
	}
}

func main() {
	var n int
	flag.IntVar(&n, "n", 1, "number of workers")
	flag.Parse()

	if n <= 0 {
		panic("n must be > 0")
	}

	fmt.Println(n, "workers")

	ch := make(chan any)

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	go func(ctx context.Context) {
		ticker := time.NewTicker(50 * time.Millisecond)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				close(ch)
				return
			case <- ticker.C:
				ch <- rand.Intn(1000)
			}
		}
	}(ctx)

	wg := sync.WaitGroup{}
	wg.Add(n)

	for i := 0; i < n; i++ {
		i := i+1
		go func() {
			defer wg.Done()
			Worker(ch, i)
		}()
	}

	wg.Wait()
}
