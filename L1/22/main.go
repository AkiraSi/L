package main

import (
	"fmt"
	"math/big"
)

func main() {
	distance := new(big.Int)
	distance.SetString("24000000000000000000", 10)
	fmt.Println(distance.Add(distance, distance))      // сложение
	fmt.Println(distance.Mul(distance, distance))      // умножение
	fmt.Println(distance.Sub(distance, big.NewInt(5))) // деление
	fmt.Println(distance.Div(distance, distance))      // разность
}
