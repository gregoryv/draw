package shape

import "testing"

func TestDiamond(t *testing.T) {
	d := NewDiamond(3, 10)
	saveAsSvg(t, d, "testdata/diamond.svg")
	testShape(t, d)
}
