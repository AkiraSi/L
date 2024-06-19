package main

import (
	"fmt"
	"strings"
)

func perevorot(inp string) string {
	finp := strings.Split(inp, " ") // массив из строки по словам
	if len(finp) <= 1 {             // если 1 слово или менее - то переворачивать смысла нет
		return inp
	}
	inp = ""                            // обнуляем
	inp += finp[len(finp)-1] + " "      // добавляем последний элемент первым
	for i := 1; i <= len(finp)-2; i++ { // итерируемся от 1 до предпоследнего
		inp += finp[i] + " "
	}
	inp += finp[0] // добавляем первый элемент последним
	return inp
}

func main() {
	inp := "snow dog sun"
	fmt.Println(inp, " — ", perevorot(inp))
}
