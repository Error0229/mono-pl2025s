package shapes

type Composite struct {
	shapes []Shape
}

func (c *Composite) Add(s Shape) bool {
	c.shapes = append(c.shapes, s)
	return true
}

func (c Composite) Area() float64 {
	return SumOfArea(c.shapes)
}
