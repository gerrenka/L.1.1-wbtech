package main

import (
	"fmt"
	"time"
)

func worker(stopCh chan bool) {
	for {
		select {
		case <-stopCh:
			fmt.Println("Горутина остановлена")
			return
		default:
			fmt.Println("Работаю...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	stopCh := make(chan bool)
	go worker(stopCh)

	time.Sleep(2 * time.Second)
	stopCh <- true
	time.Sleep(1 * time.Second)
}