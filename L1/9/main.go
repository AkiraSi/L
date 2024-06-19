package main

import (
	"fmt"
	"time"
)

func sendValues(ch chan<- int, data []int) {
	for _, v := range data {
		ch <- v // отправляем число в канал
	}
	close(ch) // отправили все - закрыли канал
}

func processValues(in <-chan int, out chan<- int) {
	for val := range in {
		out <- val * 2 // отправляем результат в канал out
	}
	close(out) // отправили все - закрыли канал
}

func readValues(ch <-chan int) {
	for val := range ch { // читаем канал
		fmt.Println(val)
	}
}

func main() {
	a := []int{1, 2, 3, 4, 5}

	ch1 := make(chan int) // канал для отправки чисел
	ch2 := make(chan int) // канал для отправки результата

	go sendValues(ch1, a)      // отправляем числа в канал ch1
	go processValues(ch1, ch2) // обрабатываем значения из ch1 и отправляем в ch2
	go readValues(ch2)         // читаем результат из ch2

	for i := 0; i < len(a); i++ { // ждем обработки всех данных, как закончится поток - значит все закончилось!
		time.Sleep(1 * time.Millisecond) // задержка
		<-ch2
	}
}
