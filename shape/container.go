package shape

import (
	"fmt"
	"io"

	"github.com/gregoryv/nexus"
)

// NewContainer returns a dashed box surrounding the given shapes,
// placing the label shape in the bottom left corner.
func NewContainer(label Shape, shapes ...Shape) *Container {
	return &Container{
		Label: label,
		Group: NewGroup(shapes...),
		class: "container",
	}
}

type Container struct {
	Label Shape
	*Group
	class string
}

func (c *Container) String() string {
	return fmt.Sprintf("Container")
}

func (c *Container) WriteSVG(out io.Writer) error {
	w, err := nexus.NewPrinter(out)
	x, y := c.TopLeftPos()
	w.Printf(
		`<rect class="%s" x="%v" y="%v" width="%v" height="%v"/>`,
		c.class, x, y, c.Width(), c.Height())
	w.Printf("\n")

	// move label to bottom left
	y += c.Height()
	y -= c.Label.Height()
	y -= c.Group.Pad.Bottom

	x += c.Group.Pad.Left

	c.Label.SetX(x)
	c.Label.SetY(y)
	c.Label.WriteSVG(w)
	return *err
}

func (c *Container) Height() int {
	h := c.Group.Height() + c.Label.Height()
	h += c.Group.Pad.Top // above the label
	return h
}

func (c *Container) Width() int {
	return c.Group.Width() + c.Label.Width()
}
