package main

import (
	"fmt"
	"math"
)

// Point represents a point in 2D space
type Point struct {
	x, y float64
}

// NewPoint is a constructor for creating a new point
func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

// Distance calculates the distance between the current point and another point
func (p *Point) Distance(other *Point) float64 {
	return math.Sqrt(math.Pow(other.x-p.x, 2) + math.Pow(other.y-p.y, 2))
}

func main() {
	// Create two points using the constructor
	p1 := NewPoint(3, 4)
	p2 := NewPoint(7, 1)

	// Calculate the distance between the points
	distance := p1.Distance(p2)

	fmt.Printf("Distance between points: %.2f\n", distance)
}

/*
 - Output: -
Distance between points: 5.00
*/
