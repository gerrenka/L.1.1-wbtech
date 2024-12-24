package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	ch1 := make(chan int)


	matrix := []int{2, 4, 6, 8, 10}
	go func() {
		defer close(ch)
	for _, num := range matrix {
		ch <- num
		}
		}()

	go func() {
		for num := range ch {
			ch1 <- num * num
		}
		defer close(ch1)
	}()

	for num := range ch1 {
		fmt.Println(num)
		time.Sleep(500 * time.Millisecond)		
	}
	}