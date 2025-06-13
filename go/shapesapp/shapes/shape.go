package shapes

type Shape interface {
	Area() float64
	Add(Shape) bool
}

type DefaultAdd struct{}

func (DefaultAdd) Add(Shape) bool {
	return false
}
