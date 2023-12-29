package main

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

// К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.
//
// Так как gc у нас довольно ленивый, он не будет освобождать неиспользуемую часть базового массива
// Поэтому надо самим выделить область памяти для строки и туда скопировать срез строки

func createHugeString(size int) string {
	var b strings.Builder
	b.WriteString(strings.Repeat("a", size))

	return b.String()
}

var justString string

func someFunc() string {
	v := createHugeString(1 << 29)

	return string([]byte(v[:100]))
	// return strings.Clone(v[:100])
}

func main() {
	justString = someFunc()
	fmt.Println(justString)
	for i := 0; i < 60; i++ {
		runtime.GC()
		time.Sleep(time.Second)
	}
}
