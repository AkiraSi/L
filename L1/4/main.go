package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
	main := make(chan int)                                  // главный поток
	ctx, cancel := context.WithCancel(context.Background()) // Контекст для отслеживания сигнала о завершении
	defer cancel()

	var numWorkers int
	fmt.Scanf("%d", &numWorkers) // кол-во воркеров

	var wg sync.WaitGroup             // счетчик горутин
	wg.Add(numWorkers)                // добавление в счетчик горутин
	for i := 0; i < numWorkers; i++ { // создает воркеры с привязкой к контексту и счетчику, читает из main (Главного потока)
		go worker(ctx, &wg, main, i)
	}
	go func() { // Главный поток, записывающий данные в канал
		for {
			select {
			case <-ctx.Done(): // если завершили
				return
			default: // обычная работа
				main <- rand.Intn(1000)            // случайное число в канал
				time.Sleep(100 * time.Millisecond) // задержка, для отладки
			}
		}
	}()
	sigChan := make(chan os.Signal, 1) // ждет сигнала прерывания
	signal.Notify(sigChan, os.Interrupt)
	<-sigChan
	cancel()  // отмена контекста
	wg.Wait() // он завершает всех воркеров, когда контекст отменяется
}

func worker(ctx context.Context, wg *sync.WaitGroup, main <-chan int, id int) { // канал, читающий данные из канала main и выводящий в fmt.println (stdout)
	defer wg.Done()
	for {
		select {
		case <-ctx.Done(): // контекст выбран тк дает возможность при отмене главного потока отмену всех других горутин, что позволяет условно в лог записать последние данные или инфу какую-то
			return
		case data := <-main:
			fmt.Println(data, id) // добавлен еще id, чтобы понимать какой воркер
		}
	}
}
