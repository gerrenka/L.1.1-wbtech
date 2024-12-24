package main

import (
    "fmt"
)

func setBit(x int64, i uint) int64 {
    return x | (1 << i)
}


func clearBit(x int64, i uint) int64 {
    return x & ^(1 << i)
}

func main() {
    var x int64 = 0

    x = setBit(x, 4)
    fmt.Printf("После установки 4-го бита: %064b\n", x)

    x = clearBit(x, 4)
    fmt.Printf("После сброса 4-го бита:   %064b\n", x)
}
