package shape

import (
	"io"
	"strings"

	"github.com/gregoryv/draw"
	"github.com/gregoryv/draw/xy"
	"github.com/gregoryv/nexus"
)

func NewNote(text string) *Note {
	return &Note{
		text:  text,
		Font:  draw.DefaultFont,
		Pad:   draw.DefaultPad,
		class: "note",
	}
}

type Note struct {
	x, y int
	text string

	Font  draw.Font
	Pad   draw.Padding
	class string
}

func (n *Note) SetText(v string) { n.text = v }
func (n *Note) Text() string     { return n.text }

func (n *Note) Position() (x int, y int) { return n.x, n.y }
func (n *Note) SetX(x int)               { n.x = x }
func (n *Note) SetY(y int)               { n.y = y }

func (n *Note) Direction() Direction { return DirectionRight }

func (n *Note) Width() int {
	var width int
	var widestLine string
	for _, line := range strings.Split(n.text, "\n") {
		w := n.Font.TextWidth(line)
		if w > width {
			width = w
			widestLine = line
		}
	}
	return boxWidth(n.Font, n.Pad, widestLine)
}

func (n *Note) Height() int {
	lines := strings.Count(n.text, "\n") + 1
	return boxHeight(n.Font, n.Pad, lines)
}
func (n *Note) SetClass(c string) { n.class = c }

func (n *Note) WriteSVG(out io.Writer) error {
	x, y := n.Position()
	w := n.Width()
	h := n.Height()
	flap := 10
	t, err := nexus.NewPrinter(out)
	/*
	   x,y
	    +---------------+        -
	    |               |\       |  flap
	    |               +-+      -
	    |                 |
	    +-----------------+
	   y+h               x+w
	*/
	t.Printf(`<path class="%s-box" d="M%v,%v `, n.class, x, y)
	t.Printf(`v %v h %v v %v l %v,%v L %v,%v M%v,%v h %v v %v"/>`,
		h, w, -(h - flap), -flap, -flap, x, y, x+w, y+flap, -flap, -flap)
	t.Print("\n")
	x += n.Pad.Left
	for i, line := range strings.Split(n.text, "\n") {
		t.Printf(`<text class="note" font-size="%vpx" x="%v" y="%v">%s</text>`,
			n.Font.Height, x, y+(n.Font.LineHeight*(i+1)), line)
		t.Print("\n")
	}
	return *err
}

func (n *Note) Edge(start xy.Point) xy.Point {
	return boxEdge(start, n)
}
