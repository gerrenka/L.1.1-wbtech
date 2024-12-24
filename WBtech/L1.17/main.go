package main

import (
    "fmt"
    "sort"
)

func main() {
    arr := []int{1, 3, 4, 7, 9, 11}
    target := 7

    index := sort.Search(len(arr), func(i int) bool {
        return arr[i] >= target
    })

    if index < len(arr) && arr[index] == target {
        fmt.Printf("Число %d найдено по индексу %d\n", target, index)
    } else {
        fmt.Printf("Число %d не найдено\n", target)
    }
}
