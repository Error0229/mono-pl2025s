package shapes

import "math"

// Point represents a point in 2D space.
type Point struct {
	x float64 // field X float64
	Y float64 // field Y float64
	DefaultAdd
}

func NewPoint(x, y float64) *Point {
	return &Point{x: x, Y: y} // structure literal
}

func (p Point) Distance(q Point) float64 { // p is a value receiver
	dx := p.x - q.x
	dy := p.Y - q.Y
	p.x += 1
	return math.Sqrt(dx*dx + dy*dy) // return expression
}

func (p *Point) Move(dx, dy float64) { // p is a pointer receiver
	p.x += dx
	p.Y += dy
}

func (p Point) Area() float64 {
	return 0.0
}
