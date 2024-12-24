package main
import (
	"fmt"
	"context"
	"time"
)

func worker(ctx context.Context) {
	for{
		select{
			case <-ctx.Done():
				fmt.Println("Горутина остановлена")
				return
			default:
				fmt.Println("Работаю...")
				time.Sleep(500 * time.Millisecond)
		}
	}
}

func main () {
	ctx, cancel := context.WithCancel(context.Background())
	go worker(ctx)
	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)

}