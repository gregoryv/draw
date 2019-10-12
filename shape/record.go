package shape

import (
	"bytes"
	"fmt"
	"io"
	"math"
	"reflect"

	"github.com/gregoryv/go-design/xy"
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
func (r *Record) Direction() Direction { return LR }
func (r *Record) SetClass(c string)    { r.class = c }

func (r *Record) WriteSvg(out io.Writer) error {
	w, err := newTagPrinter(out)
	w.printf(
		`<rect class="%s" x="%v" y="%v" width="%v" height="%v"/>`,
		r.class, r.X, r.Y, r.Width(), r.Height())
	w.printf("\n")
	var y = boxHeight(r.Font, r.Pad, 1) + r.Pad.Top
	hasFields := len(r.Fields) != 0
	if hasFields {
		r.writeSeparator(w, r.Y+y)
		for _, txt := range r.Fields {
			label := &Label{
				Pos: xy.Position{
					r.X + r.Pad.Left,
					r.Y + y,
				},
				Font:  r.Font,
				Text:  txt,
				class: "field",
			}
			label.WriteSvg(w)
			y += r.Font.LineHeight
			w.printf("\n")
		}
	}
	if len(r.Methods) != 0 {
		if hasFields {
			y += r.Pad.Bottom
		}
		r.writeSeparator(w, r.Y+y)
		for _, txt := range r.Methods {
			label := &Label{
				Pos: xy.Position{
					r.X + r.Pad.Left,
					r.Y + y,
				},
				Font:  r.Font,
				Text:  txt,
				class: "method",
			}
			label.WriteSvg(w)
			y += r.Font.LineHeight
			w.printf("\n")
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
	return line.WriteSvg(w)
}

func (r *Record) title() *Label {
	return &Label{
		Pos: xy.Position{
			r.X + r.Pad.Left,
			r.Y + r.Pad.Top,
		},
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

func (r *Record) addFields(t reflect.Type) {
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if isPublic(field.Name) {
			r.Fields = append(r.Fields, field.Name)
		}
	}
}

func (rec *Record) addMethods(t reflect.Type) {
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		if isPublic(m.Name) {
			rec.Methods = append(rec.Methods, m.Name+"()")
		}
	}
}

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

func isPublic(name string) bool {
	up := bytes.ToUpper([]byte(name))
	return []byte(name)[0] == up[0]
}

// NewStructRecord returns a record shape based on a Go struct type.
// Reflection is used.
func NewStructRecord(obj interface{}) *Record {
	t := reflect.TypeOf(obj)
	rec := NewRecord(t.String() + " struct")
	rec.addFields(t)
	rec.addMethods(reflect.PtrTo(t))
	return rec
}

func NewInterfaceRecord(obj interface{}) *Record {
	t := reflect.TypeOf(obj).Elem()
	rec := NewRecord(t.String() + " interface")
	rec.addMethods(t)
	return rec
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

type Edge interface {
	Edge(start xy.Position) xy.Position
}

// Edge returns xy position of a line starting at start and
// pointing to the rs center.
func (r *Record) Edge(start xy.Position) xy.Position {
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
