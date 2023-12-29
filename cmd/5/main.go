package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

// Разработать программу, которая будет последовательно отправлять значения в канал,
// а с другой стороны канала — читать.
// По истечению N секунд программа должна завершаться.
func main() {
	var t int
	flag.IntVar(&t, "t", 1, "seconds")
	flag.Parse()

	if t <= 0 {
		panic("t must be > 0")
	}

	fmt.Println(t, "seconds")

	ch := make(chan any)

	go func() {
		ticker := time.NewTicker(50 * time.Millisecond)
		defer ticker.Stop()

		timer := time.After(time.Duration(t) * time.Second)

		for {
			select {
			case <-timer:
				close(ch)
				return
			case <-ticker.C:
				ch <- rand.Intn(1000)
			}
		}
	}()

	for v := range ch {
		fmt.Println(v)
	}
}
