package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

func (p *Point) Point() (float64, float64) {
	return p.x, p.y
}

func (p *Point) Distance(other *Point) float64 {
	x1, y1 := p.Point()
	x2, y2 := other.Point()
	dx := x2 - x1
	dy := y2 - y1
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	p1 := NewPoint(1, 2)
	p2 := NewPoint(-3, -4)
	distance := p1.Distance(p2)
	fmt.Printf("Distance between p1 and p2: %.2f\n", distance)
}
