package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(0)
	b := big.NewInt(0)

	a.SetString("1048577", 10) 
	b.SetString("2097153", 10)

	sum := big.NewInt(0).Add(a, b) 
	diff := big.NewInt(0).Sub(a, b) 
	product := big.NewInt(0).Mul(a, b)
	quotient := big.NewInt(0).Div(a, b)

	fmt.Println("Число a:", a)
	fmt.Println("Число b:", b)
	fmt.Println("Сумма a + b:", sum)
	fmt.Println("Разность a - b:", diff)
	fmt.Println("Произведение a * b:", product)
	fmt.Println("Частное a / b:", quotient)
}
