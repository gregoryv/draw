package shape

import (
	"fmt"
	"io"
	"strings"

	"github.com/gregoryv/draw"
	"github.com/gregoryv/draw/xy"
)

func NewNote(text string) *Note {
	return &Note{
		Text:  text,
		Font:  DefaultFont,
		Pad:   DefaultPad,
		class: "note",
	}
}

type Note struct {
	Pos  xy.Position
	Text string

	Font
	Pad   Padding
	class string
}

func (n *Note) String() string {
	return fmt.Sprintf("Note at %v", n.Pos)
}
func (note *Note) Position() (int, int) { return note.Pos.XY() }
func (note *Note) Direction() Direction { return LR }
func (note *Note) SetX(x int)           { note.Pos.X = x }
func (note *Note) SetY(y int)           { note.Pos.Y = y }

func (n *Note) Width() int {
	var width int
	var widestLine string
	for _, line := range strings.Split(n.Text, "\n") {
		w := n.TextWidth(line)
		if w > width {
			width = w
			widestLine = line
		}
	}
	return boxWidth(n.Font, n.Pad, widestLine)
}

func (n *Note) Height() int {
	lines := strings.Count(n.Text, "\n") + 1
	return boxHeight(n.Font, n.Pad, lines)
}
func (n *Note) SetClass(c string) { n.class = c }

func (n *Note) WriteSvg(out io.Writer) error {
	x, y := n.Pos.XY()
	w := n.Width()
	h := n.Height()
	flap := 10
	t, err := draw.NewTagPrinter(out)
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
	for i, line := range strings.Split(n.Text, "\n") {
		t.Printf(`<text class="note" font-size="%vpx" x="%v" y="%v">%s</text>`,
			n.Font.Height, x, y+(n.Font.LineHeight*(i+1)), line)
		t.Print("\n")
	}
	return *err
}

func (n *Note) Edge(start xy.Position) xy.Position {
	return boxEdge(start, n)
}
