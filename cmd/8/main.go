package main

import (
	"flag"
	"fmt"
)


func setBit(v int64, i int, n bool) int64 {
	var result int64

	// 1 << i: создает число со здвигом битов на i влево с заполнением нулями
	if n {
		// 1<<i: маска при | (OR) заменяет 0 на 1 в v по 1 битам в маске
		result = v | 1<<i
	} else {
		// ^(1<<i): инвертированная маска при & (AND) заменяет 1 на 0 в v по 0 битам в маске
		result = v & ^(1 << i)
	}

	return result
}

// Дана переменная int64.
// Разработать программу которая устанавливает i-й бит в 1 или 0.
func main() {
	var v int64
	var i uint
	var n int

	flag.Int64Var(&v, "v", 0, "value")
	flag.UintVar(&i, "i", 0, "bit index")
	flag.IntVar(&n, "n", 0, "bit")

	flag.Parse()

	var bit bool
	switch n {
	case 1:
		bit = true
	case 0:
		bit = false
	default:
		panic("n must be 0 or 1")
	}

	fmt.Printf("%d: %b\n", v, v)

	result := setBit(v, int(i), bit)
	fmt.Printf("%d: %b\n", result, result)
}
