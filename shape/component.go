package shape

import (
	"fmt"
	"io"

	"github.com/gregoryv/draw"
	"github.com/gregoryv/draw/xy"
)

func NewComponent(title string) *Component {
	return &Component{
		Title:    title,
		Font:     DefaultFont,
		Pad:      DefaultTextPad,
		class:    "component",
		sbWidth:  10,
		sbHeight: 5,
	}
}

type Component struct {
	X, Y  int
	Title string

	Font  Font
	Pad   Padding
	class string

	//smallBoxWidth
	sbWidth  int
	sbHeight int
}

func (c *Component) String() string {
	return fmt.Sprintf("R %q", c.Title)
}

func (c *Component) Position() (int, int) { return c.X, c.Y }
func (c *Component) SetX(x int)           { c.X = x }
func (c *Component) SetY(y int)           { c.Y = y }
func (c *Component) Direction() Direction { return RightDir }
func (c *Component) SetClass(v string)    { c.class = v }

func (c *Component) WriteSVG(out io.Writer) error {
	w, err := draw.NewTagWriter(out)
	w.Printf(
		`<rect class="%s" x="%v" y="%v" width="%v" height="%v"/>`,
		c.class, c.X, c.Y, c.Width(), c.Height())
	w.Printf("\n")
	// small boxes
	w.Printf(
		`<rect class="%s" x="%v" y="%v" width="%v" height="%v"/>`,
		c.class, c.X-c.sbWidth/2, c.Y+c.sbHeight, c.sbWidth, c.sbHeight)
	w.Printf(
		`<rect class="%s" x="%v" y="%v" width="%v" height="%v"/>`,
		c.class, c.X-c.sbWidth/2, c.Y+c.Height()-c.sbHeight*2, c.sbWidth, c.sbHeight)

	c.title().WriteSVG(w)
	return *err
}

func (c *Component) title() *Label {
	return &Label{
		x:     c.X + c.Pad.Left + c.sbWidth/2,
		y:     c.Y + c.Pad.Top/2,
		Font:  c.Font,
		Text:  c.Title,
		class: c.class + "-title",
	}
}

func (c *Component) SetFont(f Font)         { c.Font = f }
func (c *Component) SetTextPad(pad Padding) { c.Pad = pad }

func (c *Component) Height() int {
	return boxHeight(c.Font, c.Pad, 1)
}

func (c *Component) Width() int {
	return boxWidth(c.Font, c.Pad, c.Title) + c.sbWidth/2
}

// Edge returns intersecting position of a line starting at start and
// pointing to the components center.
func (c *Component) Edge(start xy.Position) xy.Position {
	return boxEdge(start, c)
}
