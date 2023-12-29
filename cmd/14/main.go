package main

import "fmt"

// Разработать программу, которая в рантайме способна
// определить тип переменной: int, string, bool, channel из переменной типа interface{}
func main() {
	a := interface{}(true)

	switch a.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case chan any:
		fmt.Println("chan")
	}
}
