package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4}
	var i int
	fmt.Scan(&i)
	for j := i - 1; j < len(a)-1; j++ {
		a[j] = a[j+1] // перезаписываем каждый элемент со сдвигом влево от i-1 до конца
	}
	a = a[:len(a)-1] // уменьшаем длину на 1
}
