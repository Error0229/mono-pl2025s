package pipeline

import (
	"testing"
)

func TestCurrying(t *testing.T) {
	x := CurriedAdd(1)(2)(3)

	if x != 5 {
		t.Errorf("Expected %q but got %q", 5, x)
	}
}
