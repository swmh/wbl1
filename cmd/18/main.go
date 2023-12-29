package main

import (
	"fmt"
	"sync"
)


type Counter struct {
	value int
	mu   sync.Mutex
}

func (c *Counter) Inc() {
	c.mu.Lock()
	c.value++
	c.mu.Unlock()
}


// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
// По завершению программа должна выводить итоговое значение счетчика.
func main() {
	c := Counter{}

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				c.Inc()
			}
		}()
	}

	wg.Wait()
	fmt.Println(c.value)
}
