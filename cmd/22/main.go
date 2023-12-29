package main

import (
	"fmt"
	"math"
	"math/big"
)

// Разработать программу, которая
// перемножает, делит, складывает, вычитает
// две числовых переменных a,b, значение которых > 2^20.
func main() {
	i1 := big.NewInt(math.MaxInt)
	i2 := big.NewInt(math.MaxInt)

	result := i1.Mul(i1, i2)

	fmt.Println(result)
}
