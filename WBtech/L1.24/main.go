package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) Point {
	return Point{x: x, y: y}
}

func (p Point) Distance(other Point) float64 {
	return math.Sqrt(math.Pow(other.x-p.x, 2) + math.Pow(other.y-p.y, 2))
}

func main() {
	
	point1 := NewPoint(3, 4)
	point2 := NewPoint(7, 1)

	distance := point1.Distance(point2)

	fmt.Printf("Расстояние между точками: %.2f\n", distance)
}
