package shape

import (
	"fmt"
	"io"

	"github.com/gregoryv/draw/xy"
	"github.com/gregoryv/nexus"
)

// NewHexagon with a title. Radius must be > 0 and is the distance from
// left/right corners to imaginary vertical line.
func NewHexagon(title string, width, height, radius int) *Hexagon {
	if radius < 1 {
		radius = 1
	}
	return &Hexagon{
		Title:  title,
		Font:   DefaultFont,
		Pad:    DefaultTextPad,
		class:  "hexagon",
		width:  width,
		height: height,
		radius: radius,
	}
}

type Hexagon struct {
	x, y  int
	Title string

	Font  Font
	Pad   Padding
	class string

	width, height, radius int // radius is the left/right corner distance

	textAlign string
}

func (r *Hexagon) String() string {
	return fmt.Sprintf("R %q at %v, %v", r.Title, r.x, r.y)
}

func (r *Hexagon) Position() (x int, y int) { return r.x, r.y }
func (r *Hexagon) SetX(x int)               { r.x = x }
func (r *Hexagon) SetY(y int)               { r.y = y }

func (r *Hexagon) Direction() Direction { return DirectionRight }
func (r *Hexagon) SetClass(c string)    { r.class = c }

func (r *Hexagon) WriteSVG(out io.Writer) error {
	w, err := nexus.NewPrinter(out)
	x, y := r.Position()
	rad := r.radius
	hlen := r.width - 2*rad
	h2 := r.height / 2

	w.Printf(`<path class="%s" d="M%v,%v l %v,%v %v,%v %v,%v %v,%v %v,%v %v,%v" />`,
		r.class,
		x+rad, y,
		hlen, 0,
		rad, h2, // far right corner
		-rad, h2,
		-hlen, 0,
		-rad, -h2, // far left corner
		rad, -h2,
	)

	w.Printf("\n")
	r.title().WriteSVG(w)
	return *err
}

func (r *Hexagon) title() *Label {
	label := &Label{
		Font:  r.Font,
		Text:  r.Title,
		class: r.class + "-title",
	}
	label.Font.LineHeight = DefaultFont.Height - 4
	align := Aligner{}
	align.VAlignCenter(r, label)
	align.HAlignCenter(r, label)
	return label
}

func (r *Hexagon) SetFont(f Font)         { r.Font = f }
func (r *Hexagon) SetTextPad(pad Padding) { r.Pad = pad }
func (r *Hexagon) Width() int             { return r.width }
func (r *Hexagon) Height() int            { return r.height }
func (r *Hexagon) SetWidth(w int)         { r.width = w }
func (r *Hexagon) SetHeight(h int)        { r.height = h }

// Edge returns intersecting position of a line starting at start and
// pointing to the rect center.
func (r *Hexagon) Edge(start xy.Point) xy.Point {
	return boxEdge(start, r)
}
