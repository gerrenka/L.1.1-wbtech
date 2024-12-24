package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(wg *sync.WaitGroup, stopCh chan bool) {
	defer wg.Done()
	for {
		select {
		case <-stopCh:
			fmt.Println("Горутина остановлена через WaitGroup")
			return
		default:
			fmt.Println("Работаю...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	stopCh := make(chan bool)

	wg.Add(1)
	go worker(&wg, stopCh)

	time.Sleep(2 * time.Second)
	stopCh <- true
	wg.Wait()
}
