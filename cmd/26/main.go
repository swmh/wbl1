package main

import (
	"fmt"
	"strings"
)


func IsUniq(s string) bool {
	s = strings.ToLower(s)

	resMap := make(map[rune]bool)

	for _, v := range s {
		if resMap[v] {
			return false
		}

		resMap[v] = true
	}

	return true
} 


// Разработать программу, которая проверяет, что все символы в строке уникальные
// (true — если уникальные, false etc).
// Функция проверки должна быть регистронезависимой.
//
// Например: 
// abcd — true
// abCdefAaf — false
// 	aabcd — false
func main() {
	s1 := "abcd"
	s2 := "abCdefAaf"
	s3 := "aabcd"

	fmt.Println(IsUniq(s1))
	fmt.Println(IsUniq(s2))
	fmt.Println(IsUniq(s3))

}
