package main

import (
	"fmt"
	"time"
)

func Sleep(d time.Duration) {
	<-time.After(d)
}

// Реализовать собственную функцию sleep
func main() {
	fmt.Println("Start")
	t := time.Now()
	Sleep(3 * time.Second)
	fmt.Printf("%v\n", time.Since(t))
}
