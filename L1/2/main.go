package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup // счетчик горутин
	wg.Add(len(numbers))  // присвоение кол-ва, равное кол-ву значений numbers

	for _, number := range numbers {
		go func(num int) {
			defer wg.Done() // декремент счетчика
			fmt.Println(num * num)
		}(number)
	}

	wg.Wait() // ждет конца горутин
}
