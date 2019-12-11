package shape

import "testing"

func TestDot(t *testing.T) {
	c := NewDot(24)
	testShape(t, c)
}
