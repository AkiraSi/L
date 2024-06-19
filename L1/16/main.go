package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{7, 2, 9, 1, 5, 3, 8, 4, 6}
	sort.Ints(a) // встроенная
	fmt.Println(a)
	a = []int{7, 2, 9, 1, 5, 3, 8, 4, 6}
	quicksort(a) // своя
	fmt.Println(a)
}

func quicksort(data []int) {
	if len(data) <= 1 {
		return // один элемент или пустота
	}
	op := data[len(data)/2] // опорник, база
	left := []int{}
	right := []int{}
	for _, v := range data {
		if v < op {
			left = append(left, v)
		} else if v > op {
			right = append(right, v)
		}
	}
	quicksort(left)
	quicksort(right) // рекурсия левой и правой
	for i := 0; i < len(left); i++ {
		data[i] = left[i]
	}
	data[len(left)] = op
	for i := 0; i < len(right); i++ {
		data[len(left)+1+i] = right[i]
	} // соединяем
}
