package main

import "fmt"

type Human struct {
	name string
	age uint
}

func (h Human) Walk() {
	fmt.Println(h.name, "is walking")
}

type Action struct {
	Human
}

// Дана структура Human (с произвольным набором полей и методов). 
// Реализовать встраивание методов в структуре Action от родительской структуры 
// Human (аналог наследования).
func main() {
	a := Action{Human{"John", 30}}
	a.Walk()
}
