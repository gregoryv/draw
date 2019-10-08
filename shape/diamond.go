package shape

import (
	"fmt"
	"io"

	"github.com/gregoryv/go-design/xy"
)

func NewDiamond(x, y int) *Diamond {
	return &Diamond{
		Pos:    xy.Position{X: x, Y: y},
		width:  12,
		height: 8,
		class:  "diamond",
	}
}

type Diamond struct {
	Pos    xy.Position
	width  int
	height int
	class  string
}

func (d *Diamond) String() string {
	return fmt.Sprintf("Diamond at %v", d.Pos)
}

func (d *Diamond) WriteSvg(out io.Writer) error {
	w, err := newTagPrinter(out)
	x, y := d.Pos.XY()
	w2 := d.width / 2
	h2 := d.height / 2
	// the path is drawn from left to right
	w.printf(`<path class="%s" d="M%v,%v l %v,%v %v,%v %v,%v %v,%v" />`,
		d.class, x, y, w2, -h2, w2, h2, -w2, h2, -w2, -h2)
	return *err
}
