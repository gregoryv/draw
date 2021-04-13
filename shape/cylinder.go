package shape

import (
	"fmt"
	"io"

	"github.com/gregoryv/draw/xy"
	"github.com/gregoryv/nexus"
)

func NewCylinder(radius, height int) *Cylinder {
	return &Cylinder{
		Radius: radius,
		Font:   DefaultFont,
		Pad:    DefaultTextPad,
		height: height,
		class:  "cylinder",
	}
}

type Cylinder struct {
	Radius int

	Font Font
	Pad  Padding

	x, y   int // top left
	height int
	class  string
}

func (c *Cylinder) String() string {
	return fmt.Sprintf("Cylinder")
}

func (c *Cylinder) Position() (int, int) { return c.x, c.y }

func (c *Cylinder) SetX(x int) { c.x = x }
func (c *Cylinder) SetY(y int) { c.y = y }
func (c *Cylinder) Width() int {
	stroke := 2
	return (c.Radius+stroke)*2 - 2
}
func (c *Cylinder) Height() int           { return c.height }
func (c *Cylinder) Direction() Direction  { return DirectionRight }
func (c *Cylinder) SetClass(class string) { c.class = class }

func (c *Cylinder) ry() float64 {
	return float64(c.Radius) * 0.2
}

func (c *Cylinder) WriteSVG(out io.Writer) error {
	w, err := nexus.NewPrinter(out)
	rx := c.Radius
	ry := int(c.ry())
	x, y := c.Position()
	cx := x + rx
	cy := y + ry
	h := c.height - 2*ry
	w.Printf("<path class=%q d=\"", c.class)
	w.Printf("M %v %v L %v %v ", x, y+ry, x, y+h)
	w.Printf(
		"C %v %v, %v %v, %v %v ",
		x, y+h+ry*2, x+rx*2, y+h+ry*2, x+rx*2, y+h,
	)
	w.Printf("L %v %v", x+rx*2, y+ry)
	w.Print("\" />\n")
	w.Printf(
		"<ellipse class=\"%s\" cx=\"%v\" cy=\"%v\" rx=\"%v\" ry=\"%v\" />\n",
		c.class, cx, cy, rx, ry,
	)
	return *err
}

func (c *Cylinder) Edge(start xy.Position) xy.Position {
	return boxEdge(start, c)
}
