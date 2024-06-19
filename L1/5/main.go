package main

import (
	"fmt"
	"time"
)

func sendValues(ch chan<- int) {
	for i := 0; ; i++ {
		ch <- i                            // число в канал
		time.Sleep(200 * time.Millisecond) // задержатель
	}
}

func readValues(ch <-chan int) {
	for val := range ch { // читаем канал
		fmt.Println(val)
	}
}

func main() {
	var N int = 3 // Число секунд
	_, err := fmt.Scan(&N)
	if err != nil {
		fmt.Println(err)
		return
	}

	ch := make(chan int) // канал
	go sendValues(ch)    // отправлятель
	go readValues(ch)    // читатель

	timeout := time.After(time.Duration(N) * time.Second)
	<-timeout
	fmt.Println("Программа завершена")
}
