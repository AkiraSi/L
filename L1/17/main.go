package main

import (
	"fmt"
	"sort"
)

func binarySearch(data []int, target int) int {
	low := 0
	high := len(data) - 1

	for low <= high {
		mid := (low + high) / 2 // выбор опорного числа, стандартное
		if data[mid] == target {
			return mid // хеппи энд
		} else if data[mid] < target {
			low = mid + 1 // число справа
		} else {
			high = mid - 1 // число слева
		}
	}
	return -1 // если не найдено число
}

func main() {
	a := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	t := 11
	fmt.Println(binarySearch(a, t))    // совими силами
	fmt.Println(sort.SearchInts(a, t)) // тут тоже бинарный поиск, но встроенный
}
