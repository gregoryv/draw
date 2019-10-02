package shape

import "testing"

func TestTriangle(t *testing.T) {
	testShape(t, NewTriangle(0, 0, ""))
}
