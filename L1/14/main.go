package main

import (
	"fmt"
	"reflect"
)

func main() {
	var i int = 10
	var s string = "Hello"
	var b bool = true
	var ch chan int = make(chan int)

	var v1 interface{} = i
	var v2 interface{} = s
	var v3 interface{} = b
	var v4 interface{} = ch                                                                     // преобразование в интерфейсы
	fmt.Println(reflect.TypeOf(v1), reflect.TypeOf(v2), reflect.TypeOf(v3), reflect.TypeOf(v4)) // reflect.TypeOf({interface}) - выводит тип данных
}
