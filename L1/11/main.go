package main

import "fmt"

func perekrestok(set1, set2 []int) []int {
	setMap := make(map[int]bool) // число и простой bool для if
	intersection := []int{}      // пересечение

	for _, num := range set1 {
		setMap[num] = true
	} // элементы 1 сета в мапу, что по индексу сравнивать с элементами множества сет 2

	for _, num := range set2 {
		if setMap[num] { // если число в мапе и сет2
			intersection = append(intersection, num) // добавляем в пересечение, тк число Num гарантированно есть и в 1 и во 2 сете
		}
	}

	return intersection
}

func main() {
	set1 := []int{1, 2, 3, 4, 5}
	set2 := []int{3, 4, 5, 6, 7}
	result := perekrestok(set1, set2)
	fmt.Println(result)
}
