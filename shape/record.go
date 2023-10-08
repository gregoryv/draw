package shape

import (
	"fmt"
	"io"

	"github.com/gregoryv/draw"
	"github.com/gregoryv/draw/xy"
	"github.com/gregoryv/nexus"
)

func NewRecord(title string) *Record {
	return &Record{
		Title: title,
		Font:  draw.DefaultFont,
		Pad:   draw.DefaultTextPad,
		class: "record",
	}
}

type Record struct {
	x, y    int
	Title   string
	Fields  []string
	Methods []string

	Font  draw.Font
	Pad   draw.Padding
	class string
}

func (r *Record) String() string {
	return fmt.Sprintf("R %q", r.Title)
}

func (r *Record) Position() (x int, y int) { return r.x, r.y }
func (r *Record) SetX(x int)               { r.x = x }
func (r *Record) SetY(y int)               { r.y = y }
func (r *Record) Direction() Direction     { return DirectionRight }
func (r *Record) SetClass(c string)        { r.class = c }

func (r *Record) WriteSVG(out io.Writer) error {
	w, err := nexus.NewPrinter(out)
	w.Printf(
		`<rect class="%s" x="%v" y="%v" width="%v" height="%v"/>`,
		r.class, r.x, r.y, r.Width(), r.Height())
	w.Printf("\n")
	var y = boxHeight(r.Font, r.Pad, 1)
	hasFields := len(r.Fields) != 0
	if hasFields {
		r.writeSeparator(w, r.y+y)
		for _, txt := range r.Fields {
			label := &Label{

				x:     r.x + r.Pad.Left,
				y:     r.y + y,
				Font:  r.Font,
				text:  txt,
				class: "field",
			}
			label.WriteSVG(w)
			y += r.Font.LineHeight
			w.Printf("\n")
		}
	}
	if len(r.Methods) != 0 {
		if hasFields {
			y += r.Pad.Bottom
		}
		r.writeSeparator(w, r.y+y)
		for _, txt := range r.Methods {
			label := &Label{
				x:     r.x + r.Pad.Left,
				y:     r.y + y,
				Font:  r.Font,
				text:  txt,
				class: "method",
			}
			label.WriteSVG(w)
			y += r.Font.LineHeight
			w.Printf("\n")
		}
	}
	r.title().WriteSVG(w)
	return *err
}

func (r *Record) writeSeparator(w io.Writer, y1 int) error {
	line := NewLine(
		r.x, y1,
		r.x+r.Width(), y1,
	)
	line.SetClass("record-line")
	return line.WriteSVG(w)
}

func (r *Record) title() *Label {
	return &Label{
		x:     r.x + r.Pad.Left,
		y:     r.y,
		Font:  r.Font,
		text:  r.Title,
		class: "record-title",
	}
}

func (r *Record) HideFields()  { r.Fields = []string{} }
func (r *Record) HideMethods() { r.Methods = []string{} }

func (r *Record) SetFont(f draw.Font)         { r.Font = f }
func (r *Record) SetTextPad(pad draw.Padding) { r.Pad = pad }

func (r *Record) hasFields() bool  { return len(r.Fields) != 0 }
func (r *Record) hasMethods() bool { return len(r.Methods) != 0 }
func (r *Record) isEmpty() bool    { return !r.hasFields() && !r.hasMethods() }

func (r *Record) HideMethod(m string) (found bool) {
	rest := make([]string, 0)
	for _, n := range r.Methods {
		if n == m {
			found = true
			continue
		}
		rest = append(rest, n)
	}
	r.Methods = rest
	return
}

func (r *Record) Height() int {
	first := boxHeight(r.Font, r.Pad, 1)
	if r.isEmpty() {
		return first
	}
	l := len(r.Fields) + len(r.Methods)
	rest := boxHeight(r.Font, r.Pad, l)
	if r.hasFields() && r.hasMethods() {
		rest += r.Pad.Bottom
	}
	return first + rest
}

func (r *Record) Width() int {
	width := boxWidth(r.Font, r.Pad, r.Title)
	for _, txt := range r.Fields {
		w := boxWidth(r.Font, r.Pad, txt)
		if w > width {
			width = w
		}
	}
	for _, txt := range r.Methods {
		w := boxWidth(r.Font, r.Pad, txt)
		if w > width {
			width = w
		}
	}
	return width
}

// Edge returns intersecting position of a line starting at start and
// pointing to the records center.
func (r *Record) Edge(start xy.Point) xy.Point {
	return boxEdge(start, r)
}
