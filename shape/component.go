package shape

import (
	"fmt"
	"io"

	"github.com/gregoryv/draw"
	"github.com/gregoryv/draw/xy"
	"github.com/gregoryv/nexus"
)

func NewComponent(title string) *Component {
	return &Component{
		Title:    title,
		Font:     draw.DefaultFont,
		Pad:      draw.DefaultTextPad,
		class:    "component",
		sbWidth:  10,
		sbHeight: 5,
	}
}

type Component struct {
	x, y  int
	Title string
	href  string // optional link

	Font  draw.Font
	Pad   draw.Padding
	class string

	//smallBoxWidth
	sbWidth  int
	sbHeight int
}

func (c *Component) String() string {
	return fmt.Sprintf("R %q", c.Title)
}

func (c *Component) Position() (x int, y int) { return c.x, c.y }
func (c *Component) SetX(x int)               { c.x = x }
func (c *Component) SetY(y int)               { c.y = y }
func (c *Component) Direction() Direction     { return DirectionRight }
func (c *Component) SetClass(v string)        { c.class = v }

// SetHref links the title of the component. As of v0.29.0 you can use
// [Anchor] to link the entire shape.
func (c *Component) SetHref(v string) { c.href = v }

func (c *Component) WriteSVG(out io.Writer) error {
	w, err := nexus.NewPrinter(out)
	w.Printf(
		`<rect class="%s" x="%v" y="%v" width="%v" height="%v"/>`,
		c.class, c.x, c.y, c.Width(), c.Height())
	w.Printf("\n")
	// small boxes
	w.Printf(
		`<rect class="%s" x="%v" y="%v" width="%v" height="%v"/>`,
		c.class, c.x-c.sbWidth/2, c.y+c.sbHeight, c.sbWidth, c.sbHeight)
	w.Printf(
		`<rect class="%s" x="%v" y="%v" width="%v" height="%v"/>`,
		c.class, c.x-c.sbWidth/2, c.y+c.Height()-c.sbHeight*2, c.sbWidth, c.sbHeight)

	c.title().WriteSVG(w)
	return *err
}

func (c *Component) title() *Label {
	return &Label{
		x:     c.x + c.Pad.Left + c.sbWidth/2,
		y:     c.y + c.Pad.Top/2,
		Font:  c.Font,
		Text:  c.Title,
		href:  c.href,
		class: c.class + "-title",
	}
}

func (c *Component) SetFont(f draw.Font)         { c.Font = f }
func (c *Component) SetTextPad(pad draw.Padding) { c.Pad = pad }

func (c *Component) Height() int {
	return boxHeight(c.Font, c.Pad, 1)
}

func (c *Component) Width() int {
	return boxWidth(c.Font, c.Pad, c.Title) + c.sbWidth/2
}

// Edge returns intersecting position of a line starting at start and
// pointing to the components center.
func (c *Component) Edge(start xy.Point) xy.Point {
	return boxEdge(start, c)
}
