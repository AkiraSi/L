package main

import (
	"fmt"
	"time"
)

func Sleep(d time.Duration) {
	if d < 0 { // если закинули отрицательное время
		return
	}
	ticker := time.Tick(d) // тикер
	for _ = range ticker { // до конца for и выход
		return
	}
}

func main() {
	fmt.Println("hello")
	Sleep(-5 * time.Second)
	fmt.Println("world")
}
