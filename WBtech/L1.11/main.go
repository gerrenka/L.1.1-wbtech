package main

import "fmt"

func intersect(set1, set2 []int) []int {

	elementMap := make(map[int]bool)
	for _, v := range set1 {
		elementMap[v] = true
	}

	var result []int
	for _, v := range set2 {

		if elementMap[v] {
			result = append(result, v)

			delete(elementMap, v)
		}
	}
	return result
}

func main() {
	set1 := []int{1, 2, 3, 4, 5}
	set2 := []int{3, 4, 5, 6, 7}

	intersection := intersect(set1, set2)
	fmt.Println("Пересечение множеств:", intersection)
}
