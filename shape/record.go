package shape

import (
	"bytes"
	"fmt"
	"io"
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

func (r *Record) HideFields()  { r.Fields = []string{} }
func (r *Record) HideMethods() { r.Methods = []string{} }

func (r *Record) SetFont(f Font)         { r.Font = f }
func (r *Record) SetTextPad(pad Padding) { r.Pad = pad }

func (r *Record) hasFields() bool  { return len(r.Fields) != 0 }
func (r *Record) hasMethods() bool { return len(r.Methods) != 0 }
func (r *Record) isEmpty() bool    { return !r.hasFields() && !r.hasMethods() }

func (rec *Record) addFields(t reflect.Type) {
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if isPublic(field.Name) {
			rec.Fields = append(rec.Fields, field.Name)
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

func (rec *Record) HideMethod(m string) (found bool) {
	rest := make([]string, 0)
	for _, n := range rec.Methods {
		if n == m {
			found = true
			continue
		}
		rest = append(rest, n)
	}
	rec.Methods = rest
	return
}

func isPublic(name string) bool {
	up := bytes.ToUpper([]byte(name))
	return []byte(name)[0] == up[0]
}

func (record *Record) WriteSvg(out io.Writer) error {
	w, err := newTagPrinter(out)
	w.printf(
		`<rect class="%s" x="%v" y="%v" width="%v" height="%v"/>`,
		record.class, record.X, record.Y, record.Width(), record.Height())
	w.printf("\n")
	var y = boxHeight(record.Font, record.Pad, 1) + record.Pad.Top
	hasFields := len(record.Fields) != 0
	if hasFields {
		record.writeSeparator(w, record.Y+y)
		for _, txt := range record.Fields {
			y += record.Font.LineHeight
			label := &Label{
				Pos: xy.Position{
					record.X + record.Pad.Left,
					record.Y + y,
				},
				Text:  txt,
				class: "field",
			}
			label.WriteSvg(w)
			w.printf("\n")
		}
	}
	if len(record.Methods) != 0 {
		if hasFields {
			y += record.Pad.Bottom
		}
		record.writeSeparator(w, record.Y+y)
		for _, txt := range record.Methods {
			y += record.Font.LineHeight
			label := &Label{
				Pos: xy.Position{
					record.X + record.Pad.Left,
					record.Y + y,
				},
				Text:  txt,
				class: "method",
			}
			label.WriteSvg(w)
			w.printf("\n")
		}
	}
	record.title().WriteSvg(w)
	return *err
}

func (record *Record) writeSeparator(w io.Writer, y1 int) error {
	//	y1 := record.Y + boxHeight(record.Font, record.Pad, 1)
	line := NewLine(
		record.X, y1,
		record.X+record.Width(), y1,
	)
	return line.WriteSvg(w)
}

func (r *Record) title() *Label {
	return &Label{
		Pos: xy.Position{
			r.X + r.Pad.Left,
			r.Y + r.Font.LineHeight + r.Pad.Top,
		},
		Text:  r.Title,
		class: "record-title",
	}
}

// NewStructRecord returns a record shape based on a Go struct type.
// Reflection is used
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

func (record *Record) Height() int {
	first := boxHeight(record.Font, record.Pad, 1)
	if record.isEmpty() {
		return first
	}
	l := len(record.Fields) + len(record.Methods)
	rest := boxHeight(record.Font, record.Pad, l)
	if record.hasFields() && record.hasMethods() {
		rest += record.Pad.Bottom
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

func (record *Record) String() string {
	return fmt.Sprintf("Record %q", record.Title)
}

type Edge interface {
	Edge(start xy.Position) xy.Position
}

// Edge returns xy position of a line starting at start and
// pointing to the records center.
func (record *Record) Edge(start xy.Position) xy.Position {
	center := xy.Position{
		record.X + record.Width()/2,
		record.Y + record.Height()/2,
	}
	l1 := xy.Line{start, center}

	lowY := record.Y + record.Height()
	rightX := record.X + record.Width()

	var (
		p     xy.Position
		err   error
		first bool = true
	)
	// nexus returning once intersection found
	intersect := func(side *xy.Line) {
		if !first && err == nil {
			return
		}
		p, err = l1.IntersectSegment(side)
		first = false
	}
	left := xy.NewLine(
		record.X, record.Y,
		record.X, lowY,
	)
	intersect(left)
	bottom := xy.NewLine(
		record.X, lowY,
		rightX, lowY,
	)
	intersect(bottom)
	right := xy.NewLine(
		rightX, record.Y,
		rightX, lowY,
	)
	intersect(right)
	top := xy.NewLine(
		record.X, record.Y,
		rightX, record.Y,
	)
	intersect(top)
	return p
}
