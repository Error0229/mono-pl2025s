package shapes

type Circle struct {
	Point  // structure embedding
	Radius float64
}

func NewCircle(x, y, r float64) *Circle {
	return &Circle{Point: Point{x: x, Y: y}, Radius: r}
	// return &Circle{Point{x, y}, r}
}

func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}
