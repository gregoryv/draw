package shape

import "testing"

func TestNewHexagon_default_radius(t *testing.T) {
	radius := 0
	NewHexagon("", 0, 0, radius)
}
