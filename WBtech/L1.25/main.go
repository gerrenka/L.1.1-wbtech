package main

import (
	"fmt"
	"time"
)

func CustomSleep(milliseconds int) {
	start := time.Now()
	for {
		if time.Since(start) >= time.Duration(milliseconds)*time.Millisecond {
			break
		}
	}
}

func main() {
	fmt.Println("Начало работы")

	fmt.Println("Засыпаем на 3 секунды...")
	CustomSleep(3000)

	fmt.Println("Проснулись!")
}
