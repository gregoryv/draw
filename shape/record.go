package shape

import (
	"fmt"
	"io"

	"github.com/gregoryv/draw"
	"github.com/gregoryv/draw/xy"
)

func NewRecord(title string) *Record {
	return &Record{
		Title: title,
		Font:  DefaultFont,
		Pad:   DefaultTextPad,
		class: "record",
	}
}

type Record struct {
	X, Y    int
	Title   string
	Fields  []string
	Methods []string

	Font  Font
	Pad   Padding
	class string
}

func (r *Record) String() string {
	return fmt.Sprintf("R %q", r.Title)
}

func (r *Record) Position() (int, int) { return r.X, r.Y }
func (r *Record) SetX(x int)           { r.X = x }
func (r *Record) SetY(y int)           { r.Y = y }
func (r *Record) Direction() Direction { return RightDir }
func (r *Record) SetClass(c string)    { r.class = c }

func (r *Record) WriteSvg(out io.Writer) error {
	w, err := draw.NewTagWriter(out)
	w.Printf(
		`<rect class="%s" x="%v" y="%v" width="%v" height="%v"/>`,
		r.class, r.X, r.Y, r.Width(), r.Height())
	w.Printf("\n")
	var y = boxHeight(r.Font, r.Pad, 1) + r.Pad.Top
	hasFields := len(r.Fields) != 0
	if hasFields {
		r.writeSeparator(w, r.Y+y)
		for _, txt := range r.Fields {
			label := &Label{

				x:     r.X + r.Pad.Left,
				y:     r.Y + y,
				Font:  r.Font,
				Text:  txt,
				class: "field",
			}
			label.WriteSvg(w)
			y += r.Font.LineHeight
			w.Printf("\n")
		}
	}
	if len(r.Methods) != 0 {
		if hasFields {
			y += r.Pad.Bottom
		}
		r.writeSeparator(w, r.Y+y)
		for _, txt := range r.Methods {
			label := &Label{
				x:     r.X + r.Pad.Left,
				y:     r.Y + y,
				Font:  r.Font,
				Text:  txt,
				class: "method",
			}
			label.WriteSvg(w)
			y += r.Font.LineHeight
			w.Printf("\n")
		}
	}
	r.title().WriteSvg(w)
	return *err
}

func (r *Record) writeSeparator(w io.Writer, y1 int) error {
	line := NewLine(
		r.X, y1,
		r.X+r.Width(), y1,
	)
	line.SetClass("record-line")
	return line.WriteSvg(w)
}

func (r *Record) title() *Label {
	return &Label{
		x:     r.X + r.Pad.Left,
		y:     r.Y + r.Pad.Top,
		Font:  r.Font,
		Text:  r.Title,
		class: "record-title",
	}
}

func (r *Record) HideFields()  { r.Fields = []string{} }
func (r *Record) HideMethods() { r.Methods = []string{} }

func (r *Record) SetFont(f Font)         { r.Font = f }
func (r *Record) SetTextPad(pad Padding) { r.Pad = pad }

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
func (r *Record) Edge(start xy.Position) xy.Position {
	return boxEdge(start, r)
}
