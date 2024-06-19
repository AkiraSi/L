package main

import (
	"fmt"
)

func reverseString(s string) string {
	runes := []rune(s) // преобразование строки в руны [вжух]
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i] // обмен значениями рун
	}
	return string(runes)
}

func main() {
	inputString := "главрыба"
	fmt.Println(reverseString(inputString))
}
