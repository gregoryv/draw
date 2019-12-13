package shape

import (
	"testing"
)

func TestComponent(t *testing.T) {
	r := NewComponent("a")
	testShape(t, r)
}
