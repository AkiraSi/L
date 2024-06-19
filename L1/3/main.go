package main

import (
	"sync"
)

func sum_sq(numbers []int) int {
	var wg sync.WaitGroup // Переменная с кол-вом горутин
	wg.Add(len(numbers))  // кол-во горутин в счетчик

	var sum int = 0
	sumChan := make(chan int)

	for _, number := range numbers {
		go func(n int) {
			defer wg.Done()  // декремент счетчика
			sumChan <- n * n // отправляем в канал sumchan результат
		}(number)
	}

	go func() {
		defer close(sumChan)
		for i := 0; i < len(numbers); i++ { // до len numbers, ибо такое кол-во горутин, иначе тотальный lock
			sum += <-sumChan // получаем результат и суммируем
		}
	}()
	wg.Wait() // ждем конца всех горутин
	return sum
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	sum_sq(numbers)
}
