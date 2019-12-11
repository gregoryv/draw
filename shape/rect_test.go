package shape

import (
	"testing"
)

func TestRect(t *testing.T) {
	r := NewRect("a")
	testShape(t, r)
}
