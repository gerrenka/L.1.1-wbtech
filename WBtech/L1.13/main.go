package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	a, b = b, a
	fmt.Println(a, b)
}