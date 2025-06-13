package test

import (
	"shapesapp/shapes"
	"testing"
)

func TestCircle(t *testing.T) {
	c := shapes.NewCircle(1.0, 2.0, 3.0)
	// if c.Center.x != 1.0 {
	// 	t.Errorf("Expected X to be 1.0, got %f", c.Center.X)
	// }
	if c.Point.Y != 2.0 {
		t.Errorf("Expected Y to be 2.0, got %f", c.Point.Y)
	}
	if c.Y != 2.0 { // field promotion
		t.Errorf("Expected Y to be 2.0, got %f", c.Y)
	}
	if c.Radius != 3.0 {
		t.Errorf("Expected Radius to be 3.0, got %f", c.Radius)
	}
}

func TestCircleArea(t *testing.T) {
	var c shapes.Shape = shapes.NewCircle(1.0, 2.0, 10.0)
	a := c.Area()
	if a != 314.159 {
		t.Errorf("Expected area to be 314.159, got %f", a)
	}
}

func TestSumOfArea(t *testing.T) {
	var ss = []shapes.Shape{
		shapes.NewPoint(1.0, 2.0),
		shapes.NewCircle(1.0, 2.0, 10.0),
	}
	sum := shapes.SumOfArea(ss)
	if sum != 314.159 {
		t.Errorf("Expected sum to be 314.159, got %f", sum)
	}
}
