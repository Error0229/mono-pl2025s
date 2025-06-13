package test

import (
	"shapesapp/shapes"
	"testing"
)

func TestComposite(t *testing.T) {
	var cs shapes.Shape = &shapes.Composite{}
	if ok := cs.Add(shapes.NewCircle(1.0, 2.0, 10.0)); !ok {
		t.Errorf("Expected Add to return true")
	}
	cs.Add(shapes.NewPoint(1.0, 2.0))

	area := cs.Area()

	if area != 314.159 {
		t.Errorf("Expected area to be 314.159, got %f", area)
	}
}

func TestPointAdd(t *testing.T) {
	var p shapes.Point
	if ok := p.Add(shapes.NewPoint(1.0, 2.0)); ok {
		t.Errorf("Expected Add to return false")
	}
}
