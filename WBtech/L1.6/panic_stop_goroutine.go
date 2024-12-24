package main

import (
	"fmt"
	"time"
)

func worker() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Горутина остановлена через panic:", r)
		}
	}()
	for {
		fmt.Println("Работаю...")
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	go worker()

	time.Sleep(2 * time.Second)
	panic("Принудительная остановка горутины")
}
