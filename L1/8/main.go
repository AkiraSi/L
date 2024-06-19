package main

import (
	"fmt"
)

func main() {
	var udar int64 = 0
	var i int
	var tip bool
	fmt.Scan(&i, &tip)
	if i > 64 || i < 1 { // проверка на верность бита
		return
	}
	if tip { // если 1, то 1, если 0, то не делаем ничего
		udar = 1 << (i - 1)
	}
	fmt.Printf("%064b", uint64(udar))
}
