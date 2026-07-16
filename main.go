package main

import (
	"fmt"
)

type Point struct {
	X, Y int
}

func (p Point) EuclideanDistance(p1 Point) int {
	dx := p1.X - p.X
	dy := p1.Y - p.Y
	return (dx * dx) + (dy * dy)
}

func main() {
	var x1, y1, x2, y2 int
	fmt.Scan(&x1)
	fmt.Scan(&y1)
	fmt.Scan(&x2)
	fmt.Scan(&y2)
	// Build two Points and print squared distance.
	p1 := Point{X: x1, Y: y1}
	p2 := Point{X: x2, Y: y2}

	d := p1.EuclideanDistance(p2)
	fmt.Println(d)
}
