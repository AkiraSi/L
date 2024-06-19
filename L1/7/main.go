package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	data := make(map[string]int)

	var mutex sync.Mutex

	for i := 0; i < 5; i++ {
		go func(id int) {
			for j := 0; j < 5; j++ {
				mutex.Lock()                             // локаем на запись
				data[fmt.Sprintf("id_%d_%d", id, j)] = j // присвоение ключу значения
				mutex.Unlock()                           // разблокируем после записи
			}
		}(i)
	}
	time.Sleep(time.Millisecond * 100) // задержка, чтобы все горутины успели отработать

	for k, v := range data { // вывод
		fmt.Printf("ID: %s, %d\n", k, v)
	}
}
