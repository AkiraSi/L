package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func mainGoroutineExit() {
	go func() { // Главная горутина завершится, не дожидаясь завершения фоновых горутин
		for i := 0; i < 5; i++ {
			fmt.Println(i)
			time.Sleep(500 * time.Millisecond)
		} // тут уже все, gg
	}()
}

func channelSynchronization() {
	quit := make(chan bool)
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println(i)
			time.Sleep(500 * time.Millisecond)
		}
		quit <- true // завершил работу
	}()
	<-quit // Ожидание сигнала от горутины 2 об успешном завершении
}

func contextCancellation() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		for i := 0; i < 5; i++ {
			select {
			case <-ctx.Done(): // контекст отменили - горутина завершилась
				return
			default:
				fmt.Println(i)
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()
	time.Sleep(2 * time.Second) // просто для жизни горутины, чтобы что-то вывело
	cancel()                    // Отмена контекста
	time.Sleep(1 * time.Second) // задержка для отмены
}

func waitGroupSynchronization() {
	var wg sync.WaitGroup // счетчик горутин
	wg.Add(1)             // добавили 1 горутину
	go func() {
		defer wg.Done() // счетчику сообщается, что горутина завершила свою работу
		for i := 0; i < 5; i++ {
			fmt.Println(i)
			time.Sleep(500 * time.Millisecond)
		}
	}()
	wg.Wait() // Ожидание завершения горутины
}

func main() {
	mainGoroutineExit()        // Завершение главной горутины
	channelSynchronization()   // Использование каналов для синхронизации
	contextCancellation()      // Использование контекста
	waitGroupSynchronization() // Использование sync.WaitGroup
}
