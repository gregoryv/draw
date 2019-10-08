package shape

import "testing"

func TestCircle(t *testing.T) {
	c := NewCircle(24)
	testShape(t, c)
}
