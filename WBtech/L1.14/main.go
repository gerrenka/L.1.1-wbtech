package main

import "fmt"

func determine_type(tip interface{}){
	switch tip.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case chan int:
		fmt.Println("channel chan int")
	default:
		fmt.Println("unknown")
	}
}

func main() {
	determine_type(1)
	determine_type("string")
	determine_type(true)
	determine_type(make(chan int))
	determine_type(3.14)
}