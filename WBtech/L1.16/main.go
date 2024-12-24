package main

import (
	"fmt"
)

func quicksort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}


	pivot := arr[len(arr)/2]
	var left []int
	var right []int


	for _, v := range arr {
		if v < pivot {
			left = append(left, v)
		} else if v > pivot {
			right = append(right, v)
		}
	}


	return append(append(quicksort(left), pivot), quicksort(right)...)
}

func main() {

	arr := []int{10, 7, 8, 9, 1, 5}
	fmt.Println("Исходный массив:", arr)

	sorted := quicksort(arr)
	fmt.Println("Отсортированный массив:", sorted)
}
