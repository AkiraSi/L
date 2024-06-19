package main

import (
	"fmt"
	"strings"
)

func task26(str string) bool {
	str = strings.ToLower(str)       // все в нижний, тк регистронезависимая
	charCounts := make(map[rune]int) // сет символов
	for _, char := range str {
		charCounts[char]++        // встретили букву - увеличили на 1 кол-во
		if charCounts[char] > 1 { // если их 2 и более, значит не удовл
			return false
		}
	}
	return true // если не встретили двух букв одинаковых, то тру
}

func main() {
	inp := []string{"abcd", "abCdefAaf", "aabcd", "", "  "}

	for _, str := range inp {
		fmt.Println(str, " — ", task26(str))
	}
}
