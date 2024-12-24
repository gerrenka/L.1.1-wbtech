package main

import (
	"fmt"
	"strings"
)

func areAllCharactersUnique(input string) bool {
	// Приводим строку к нижнему регистру для регистронезависимой проверки
	input = strings.ToLower(input)

	// Создаем карту для отслеживания символов
	charMap := make(map[rune]bool)

	// Проходим по каждому символу строки
	for _, char := range input {
		// Проверяем, встречался ли символ ранее
		if charMap[char] {
			return false
		}
		// Если символ не встречался, добавляем его в карту
		charMap[char] = true
	}
	return true
}

func main() {
	// Примеры использования
	testCases := []string{"Hello", "World", "GoLang", "Unique"}

	for _, testCase := range testCases {
		result := areAllCharactersUnique(testCase)
		fmt.Printf("String: %s, Unique: %t\n", testCase, result)
	}
}
