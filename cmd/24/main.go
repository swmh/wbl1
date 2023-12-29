package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func New(x float64, y float64) Point {
	return Point{x, y}
}

func (p Point) Distance(t Point) float64 {
	return math.Sqrt(math.Pow(t.x-p.x, 2) + math.Pow(t.y-p.y, 2))
}

// Разработать программу нахождения расстояния между двумя точками, 
// которые представлены в виде
// структуры Point с инкапсулированными параметрами x,y и конструктором.
func main() {
	p1 := New(0, 0)
	p2 := New(3, 4)

	fmt.Println(p1.Distance(p2))

}
