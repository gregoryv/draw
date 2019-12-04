package shape

import (
	"fmt"
	"io"
	"math"

	"github.com/gregoryv/go-design/xy"
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

func (r *Component) String() string {
	return fmt.Sprintf("R %q", r.Title)
}

func (r *Component) Position() (int, int) { return r.X, r.Y }
func (r *Component) SetX(x int)           { r.X = x }
func (r *Component) SetY(y int)           { r.Y = y }
func (r *Component) Direction() Direction { return LR }
func (r *Component) SetClass(c string)    { r.class = c }

func (r *Component) WriteSvg(out io.Writer) error {
	w, err := newTagPrinter(out)
	w.printf(
		`<rect class="%s" x="%v" y="%v" width="%v" height="%v"/>`,
		r.class, r.X, r.Y, r.Width(), r.Height())
	w.printf("\n")
	// small boxes
	w.printf(
		`<rect class="%s" x="%v" y="%v" width="%v" height="%v"/>`,
		r.class, r.X-r.sbWidth/2, r.Y+r.sbHeight, r.sbWidth, r.sbHeight)
	w.printf(
		`<rect class="%s" x="%v" y="%v" width="%v" height="%v"/>`,
		r.class, r.X-r.sbWidth/2, r.Y+r.Height()-r.sbHeight*2, r.sbWidth, r.sbHeight)

	r.title().WriteSvg(w)
	return *err
}

func (r *Component) title() *Label {
	return &Label{
		Pos: xy.Position{
			r.X + r.Pad.Left + r.sbWidth/2,
			r.Y + r.Pad.Top/2,
		},
		Font:  r.Font,
		Text:  r.Title,
		class: "record-title",
	}
}

func (r *Component) SetFont(f Font)         { r.Font = f }
func (r *Component) SetTextPad(pad Padding) { r.Pad = pad }

func (r *Component) Height() int {
	return boxHeight(r.Font, r.Pad, 1)
}

func (r *Component) Width() int {
	return boxWidth(r.Font, r.Pad, r.Title) + r.sbWidth/2
}

// Edge returns xy position of a line starting at start and
// pointing to the records center. The returned position would
// be at the crosspoint.
func (r *Component) Edge(start xy.Position) xy.Position {
	center := xy.Position{
		r.X + r.Width()/2,
		r.Y + r.Height()/2,
	}
	l1 := xy.Line{start, center}

	var (
		d      float64 = math.MaxFloat64
		pos    xy.Position
		lowY   = r.Y + r.Height()
		rightX = r.X + r.Width()
		top    = xy.NewLine(r.X, r.Y, rightX, r.Y)
		left   = xy.NewLine(r.X, r.Y, r.X, lowY)
		right  = xy.NewLine(rightX, r.Y, rightX, lowY)
		bottom = xy.NewLine(r.X, lowY, rightX, lowY)
	)

	for _, side := range []*xy.Line{top, left, right, bottom} {
		p, err := l1.IntersectSegment(side)
		if err != nil {
			continue
		}
		dist := start.Distance(p)
		if dist < d {
			pos = p
			d = dist
		}
	}
	return pos
}
