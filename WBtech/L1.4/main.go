package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(ctx context.Context, id int, dataChannel <-chan string, wg *sync.WaitGroup) {
	defer wg.Done() 
	for {
		select {
		case <-ctx.Done(): 
			fmt.Printf("Worker %d stopping...\n", id)
			return
		case data, ok := <-dataChannel: 
			if !ok {
				fmt.Printf("Worker %d: Channel closed. Stopping...\n", id)
				return
			}
			fmt.Printf("Worker %d received: %s\n", id, data)
		}
	}
}

func main() {

	var numWorkers int
	fmt.Print("Enter number of workers: ")
	_, err := fmt.Scan(&numWorkers)
	if err != nil || numWorkers <= 0 {
		fmt.Println("Invalid number of workers. Exiting...")
		return
	}

	dataChannel := make(chan string)

	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(ctx, i, dataChannel, &wg)
	}

	sigChannel := make(chan os.Signal, 1)
	signal.Notify(sigChannel, os.Interrupt, syscall.SIGTERM)

	go func() {
		counter := 0
		for {
			select {
			case <-ctx.Done():
				return
			default:
				dataChannel <- fmt.Sprintf("Message %d", counter)
				counter++
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	<-sigChannel
	fmt.Println("\nReceived interrupt signal. Shutting down...")

	cancel() 
	close(dataChannel) 
	wg.Wait()
	fmt.Println("All workers stopped. Program exited.")
}
