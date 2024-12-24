package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Горутина остановлена по таймауту")
			return
		default:
			fmt.Println("Работаю...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go worker(ctx)

	time.Sleep(3 * time.Second)
}
