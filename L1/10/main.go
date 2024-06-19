package main

import (
	"fmt"
)

func main() {
	temps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	groupedTemps := make(map[int][]float64) // мапа с группами с шагом в 10
	for _, temp := range temps {
		group := (int)(temp/10) * 10                            // группы температур с шагом 10
		groupedTemps[group] = append(groupedTemps[group], temp) // добавление в группу
	}

	for k, v := range groupedTemps {
		fmt.Printf("%d: %v", k, v) // вывод группы темперутар по ключу
	}
}
