package main

import (
	"fmt"
	"os"
	"time"
)

func worker() {
	for {
		fmt.Println("Работаю...")
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	go worker()

	time.Sleep(2 * time.Second)
	fmt.Println("Завершение программы")
	os.Exit(0) // Принудительное завершение всех горутин
}
