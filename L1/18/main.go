package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	value int
	mutex sync.Mutex
}

func (c *Counter) Increment() {
	c.value++
}

func main() {
	counter := Counter{}
	gr := 10              // кол-во потоков
	var wg sync.WaitGroup // синхронизатор горутин!
	wg.Add(gr)            // закидываем в счетчик
	for i := 0; i < gr; i++ {
		go func() {
			defer wg.Done()      // декримент
			counter.mutex.Lock() // чтоб не сломать
			counter.Increment()
			counter.mutex.Unlock() // чтоб тоже не сломать тем, что не хотел сломать
		}()
	}
	wg.Wait() // ждем всех!
	fmt.Println(counter.value)
}
