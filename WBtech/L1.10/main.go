package main

import (
	"fmt"
	"math"
)

func groupTemperatures(temps []float64) map[int][]float64 {

	groups := make(map[int][]float64)

	for _, temp := range temps {
		var group int
		if temp < 0 {
			group = int(math.Ceil(temp/10) * 10)
		} else {
			group = int(math.Floor(temp/10) * 10)
		}

		groups[group] = append(groups[group], temp)
	}

	return groups
}

func main() {
	temps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	groupedTemps := groupTemperatures(temps)

	for group, values := range groupedTemps {
		fmt.Printf("%d: %v\n", group, values)
	}
}
