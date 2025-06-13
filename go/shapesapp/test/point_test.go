package test

import (
	"shapesapp/shapes"
	"testing"
)

// func TestPoint(t *testing.T) {
// 	p := shapes.Point{X: 1.0, Y: 2.0} // structure literal
// 	if p.X != 1.0 {
// 		t.Errorf("Expected X to be 1.0, got %f", p.X)
// 	}
// 	if p.Y != 2.0 {
// 		t.Errorf("Expected Y to be 2.0, got %f", p.Y)
// 	}
// }

func TestDistance(t *testing.T) {
	p1 := shapes.NewPoint(1.0, 1.0)
	p2 := shapes.NewPoint(4.0, 5.0)
	d := p1.Distance(*p2)
	if d != 5.0 {
		t.Errorf("Expected distance to be 5.0, got %f", d)
	}
	// if p1.X != 1.0 {
	// 	t.Errorf("Expected X to be 1.0, got %f", p1.X)
	// }
	// if p1.Y != 1.0 {
	// 	t.Errorf("Expected Y to be 1.0, got %f", p1.Y)
	// }
}

func TestMove(t *testing.T) {
	p := shapes.NewPoint(1.0, 2.0)
	p.Move(2.0, 3.0) // (&p).Move(2.0, 3.0)
	// if p.X != 3.0 {
	// 	t.Errorf("Expected X to be 3.0, got %f", p.X)
	// }
	if p.Y != 5.0 {
		t.Errorf("Expected Y to be 5.0, got %f", p.Y)
	}
	// pp := &p
	// pp.Distance(shapes.Point{X: 0.0, Y: 1.0}) // (*pp).Distance(shapes.Point{X: 0.0, Y: 1.0})
}

func TestArea(t *testing.T) {
	var p shapes.Shape
	p = shapes.NewPoint(1.0, 2.0)
	a := p.Area()
	if a != 0.0 {
		t.Errorf("Expected area to be 0.0, got %f", a)
	}
}
