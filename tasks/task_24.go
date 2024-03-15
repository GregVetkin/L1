package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func (p1 *Point) Distance(p2 Point) float64 {
	return math.Sqrt((p2.x-p1.x)*(p2.x-p1.x) + (p2.y-p1.y)*(p2.y-p1.y))
}

func NewPoint(x, y float64) Point {
	return Point{x: x, y: y}
}

func Distance(p1, p2 Point) float64 {
	return math.Sqrt((p2.x-p1.x)*(p2.x-p1.x) + (p2.y-p1.y)*(p2.y-p1.y))
}



func main() {
	p1 := NewPoint(0.0, 1.0)
	p2 := NewPoint(1.0, 0.0)

	fmt.Println(p1.Distance(p2)) // 1.414
	fmt.Println(p2.Distance(p1)) // 1.414
	fmt.Println(Distance(p1, p2)) // 1.414

	fmt.Println(p1.Distance(p1)) // 0
}