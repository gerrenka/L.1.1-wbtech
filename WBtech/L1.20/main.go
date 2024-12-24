package main

import (
	"fmt"
	"strings"
)

func reverseWords(input string) string {
	words := strings.Fields(input)

	i := 0
	j := len(words) - 1
	
	for i < j {
		words[i], words[j] = words[j], words[i]
		i++
		j--
	}

	return strings.Join(words, " ")
}

func main() {
	input := "snow dog sun"

	result := reverseWords(input)
	fmt.Println("Input:", input)
	fmt.Println("Reversed:", result)
}
