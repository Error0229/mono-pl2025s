package shapes

func SumOfArea(shapes []Shape) float64 {
	var sum float64
	for _, s := range shapes {
		sum += s.Area() // polymorphism declared in interface Shape
	}
	return sum
}
