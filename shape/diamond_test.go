package shape

import "testing"

func TestDiamond(t *testing.T) {
	d := NewDiamond()
	d.SetX(3)
	d.SetY(10)
	saveAsSvg(t, d, "testdata/diamond.svg")
	testShape(t, d)
}
