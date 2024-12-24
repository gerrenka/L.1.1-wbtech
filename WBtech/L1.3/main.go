package main

import (
	"fmt"
	"sync"
)

func square(num int, wg *sync.WaitGroup, result_matrix chan int) {
	defer wg.Done()
	result_matrix <- num * num
}

func main (){
	matrix:=[]int {2,4,6,8,10}
	result_matrix:= make(chan int)
	var wg sync.WaitGroup
	for _, num := range matrix {
		wg.Add(1)
		go square(num, &wg, result_matrix)
	}
	
	go func() {
		wg.Wait()
		close(result_matrix)
	}()

	sum := 0
	for num := range result_matrix {
		sum += num
	}

	fmt.Println(sum)
}